// Package dataqual is a library that allows running of Plumber data rules against data
// This package is designed to be included in golang message bus libraries. The only public
// method is ApplyRules() which is used to run rules against data.
//
// Use of this package requires a running instance of Plumber.
// Plumber can be downloaded at https://github.com/streamdal/plumber
//
// The following environment variables must be set:
// - PLUMBER_ADDRESS: The address of the Plumber server
// - PLUMBER_TOKEN: The token to use when connecting to the Plumber server
//
// Optional parameters:
// - DATAQUAL_DRY_RUN: If true, rule hits will only be logged, no failure modes will be ran
package dataqual

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/relistan/go-director"

	"github.com/streamdal/dataqual/logger"

	protos "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/pkg/errors"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/streamdal/dataqual/common"
	"github.com/streamdal/dataqual/detective"
	"github.com/streamdal/dataqual/metrics"
	"github.com/streamdal/dataqual/plumber"
)

// Module is a constant that represents which type of WASM module we will run the rules against
type Module string

// Mode is a constant that represents whether we are publishing or consuming,
// it must match the protobuf enum of the rule
type Mode int

const (
	// Publish tells ApplyRules to run the rules against the publish ruleset
	Publish Mode = 1

	// Consume tells ApplyRules to run the rules against the consume ruleset
	Consume Mode = 2

	// RuleUpdateInterval is how often to check for rule updates
	RuleUpdateInterval = time.Second * 30

	// Match is the name of the WASM module that contains the match function
	Match Module = "match"

	// Transform is the name of the WASM module that contains the transform function
	Transform Module = "transform"

	// DefaultMaxDataSize is the maximum size of data that can be sent to the WASM module
	DefaultMaxDataSize = 1024 * 1024 // 1Mi
)

var (
	ErrEmptyConfig        = errors.New("config cannot be empty")
	ErrEmptyDataSource    = errors.New("data source cannot be empty")
	ErrMissingShutdownCtx = errors.New("shutdown context cannot be nil")

	// ErrMessageDropped is returned when a message is dropped by the plumber data rules
	// An end user may check for this error and handle it accordingly in their code
	ErrMessageDropped = errors.New("message dropped by plumber data rules")
)

type IDataQual interface {
	ApplyRules(ctx context.Context, mode Mode, key string, data []byte) ([]byte, error)
}

type DataQual struct {
	*Config
	functions    map[Module]*function
	rules        map[Mode]map[string][]*protos.Rule
	ruleSetMap   map[string]string
	ruleSetMtx   *sync.RWMutex
	functionsMtx *sync.RWMutex
	rulesMtx     *sync.RWMutex
	Plumber      plumber.IPlumberClient
	metrics      *metrics.Metrics
}

type Config struct {
	PlumberURL   string
	PlumberToken string
	WasmTimeout  time.Duration
	DryRun       bool
	DataSource   string
	ShutdownCtx  context.Context
	Logger       logger.Logger
}

func New(cfg *Config) (*DataQual, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, errors.Wrap(err, "unable to validate config")
	}

	// We instantiate this library based on whether or not we have a Plumber URL+token
	// If these are not provided, the wrapper library will not perform rule checks and
	// will act as normal
	if cfg.PlumberURL == "" || cfg.PlumberToken == "" {
		return nil, nil
	}

	plumber, err := plumber.New(cfg.PlumberURL, cfg.PlumberToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to plumber")
	}

	m, err := metrics.New(&metrics.Config{Plumber: plumber})
	if err != nil {
		return nil, errors.Wrap(err, "failed to start metrics service")
	}

	dq := &DataQual{
		functions:    make(map[Module]*function),
		functionsMtx: &sync.RWMutex{},
		Plumber:      plumber,
		rules:        make(map[Mode]map[string][]*protos.Rule),
		rulesMtx:     &sync.RWMutex{},
		ruleSetMap:   make(map[string]string),
		ruleSetMtx:   &sync.RWMutex{},
		Config:       cfg,
		metrics:      m,
	}

	if cfg.DryRun {
		cfg.Logger.Info("Plumber data rules running in dry run mode")
	}

	// Force rule pull on startup
	if err := dq.getRuleUpdates(); err != nil {
		return nil, errors.Wrap(err, "failed to get data quality rules")
	}

	ruleUpdateLooper := director.NewTimedLooper(director.FOREVER, RuleUpdateInterval, make(chan error, 1))
	go dq.watchForRuleUpdates(ruleUpdateLooper)

	return dq, nil
}

func validateConfig(cfg *Config) error {
	if cfg == nil {
		return ErrEmptyConfig
	}

	if cfg.DataSource == "" {
		return ErrEmptyDataSource
	}

	if cfg.ShutdownCtx == nil {
		return ErrMissingShutdownCtx
	}

	// Can be specified in config for lib use, or via envar for shim use
	if cfg.PlumberURL == "" {
		cfg.PlumberURL = os.Getenv("PLUMBER_URL")
	}

	// Can be specified in config for lib use, or via envar for shim use
	if cfg.PlumberToken == "" {
		cfg.PlumberToken = os.Getenv("PLUMBER_TOKEN")
	}

	// Can be specified in config for lib use, or via envar for shim use
	if os.Getenv("DATAQUAL_DRY_RUN") == "true" {
		cfg.DryRun = true
	}

	// Can be specified in config for lib use, or via envar for shim use
	if cfg.WasmTimeout == 0 {
		to := os.Getenv("DATAQUAL_WASM_TIMEOUT")
		if to == "" {
			to = "1s"
		}

		timeout, err := time.ParseDuration(to)
		if err != nil {
			return errors.Wrap(err, "unable to parse DATAQUAL_WASM_TIMEOUT")
		}
		cfg.WasmTimeout = timeout
	}

	// Default to NOOP logger if none is provided
	if cfg.Logger == nil {
		cfg.Logger = &logger.NoOpLogger{}
	}

	return nil
}

func createWASMInstance(wasmBytes []byte) (api.Module, error) {
	if len(wasmBytes) == 0 {
		return nil, errors.New("wasm data is empty")
	}

	ctx := context.Background()
	r := wazero.NewRuntime(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	cfg := wazero.NewModuleConfig().
		WithStderr(io.Discard).
		WithStdout(io.Discard).
		WithStartFunctions("") // We don't need _start() to be called for our purposes

	mod, err := r.InstantiateWithConfig(ctx, wasmBytes, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate wasm module")
	}

	return mod, nil
}

func (d *DataQual) failTransform(ctx context.Context, data []byte, cfg *protos.FailureModeTransform) ([]byte, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, d.WasmTimeout)
	defer cancel()

	// Get WASM module
	f, err := d.getFunction(Transform)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get wasm data")
	}

	var del bool
	switch cfg.Type {
	case protos.FailureModeTransform_TRANSFORM_TYPE_REPLACE:
		del = false
	case protos.FailureModeTransform_TRANSFORM_TYPE_DELETE:
		del = true
	default:
		return nil, errors.Errorf("unknown transform type: %s", cfg.Type)
	}

	request := &common.TransformRequest{
		Path:   cfg.Path,
		Value:  cfg.Value,
		Data:   data,
		Delete: del,
	}

	req, err := request.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate request")
	}

	returnData, err := f.Exec(timeoutCtx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error during wasm exec")
	}

	resp := &common.TransformResponse{}

	if err := resp.UnmarshalJSON(returnData); err != nil {
		return nil, errors.Wrap(err, "error during tinyjson.Unmarshal")
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (d *DataQual) runMatch(ctx context.Context, mt detective.MatchType, path string, data []byte, args []string) (bool, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, d.WasmTimeout)
	defer cancel()

	// Get WASM module
	f, err := d.getFunction(Match)
	if err != nil {
		return false, errors.Wrap(err, "failed to get wasm data")
	}

	request := &common.MatchRequest{
		MatchType: mt,
		Path:      path,
		Data:      data,
		Args:      args,
	}

	req, err := request.MarshalJSON()
	if err != nil {
		return false, fmt.Errorf("unable to generate request: %s", err.Error())
	}

	returnData, err := f.Exec(timeoutCtx, req)
	if err != nil {
		return false, errors.Wrap(err, "error during wasm exec")
	}

	resp := &common.MatchResponse{}

	if err := resp.UnmarshalJSON(returnData); err != nil {
		return false, errors.Wrap(err, "error during tinyjson.Unmarshal")
	}

	if resp.Error != "" {
		return false, errors.New(resp.Error)
	}

	return resp.IsMatch, nil
}

func (d *DataQual) ApplyRules(ctx context.Context, mode Mode, key string, data []byte) ([]byte, error) {
	if len(data) > DefaultMaxDataSize {
		d.Logger.Warnf("data size exceeds maximum, skipping %s rules on key %s", mode.String(), key)
		return data, nil
	}

	d.rulesMtx.RLock()
	modeSets, ok := d.rules[mode]
	defer d.rulesMtx.RUnlock()

	if !ok {
		// No rules for this mode, nothing to do
		return data, nil
	}

	// Check if we have rules for this key
	// Key = kafka topic or rabbit routing/binding key
	rules, ok := modeSets[key]
	if !ok {
		// No rules for this key, nothing to do
		return data, nil
	}

	for _, rule := range rules {
		switch rule.Type {
		case protos.RuleType_RULE_TYPE_MATCH:
			cfg := rule.GetMatchConfig()
			if cfg == nil {
				return nil, errors.New("BUG: match rule is missing match config")
			}

			isMatch, err := d.runMatch(ctx, detective.MatchType(cfg.Type), cfg.Path, data, cfg.Args)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to run match '%s' on field '%s'", cfg.Type, cfg.Path)
			}

			// Didn't hit, nothing further to do
			if !isMatch {
				continue
			}

			// Dry run, do nothing but log and continue on to the next rule
			if d.DryRun {
				d.logDryRun(rule)
				continue
			}

			var shouldDrop bool

			// There can me multiple failure modes per rule
			//
			for _, failCfg := range rule.GetFailureModeConfigs() {
				switch failCfg.Mode {
				case protos.RuleFailureMode_RULE_FAILURE_MODE_REJECT:
					shouldDrop = true
				case protos.RuleFailureMode_RULE_FAILURE_MODE_TRANSFORM:
					transformed, err := d.failTransform(ctx, data, failCfg.GetTransform())
					if err != nil {
						return nil, errors.Wrap(err, "failed to run transform")
					}

					data = transformed
				case protos.RuleFailureMode_RULE_FAILURE_MODE_ALERT_SLACK:
					fallthrough
				case protos.RuleFailureMode_RULE_FAILURE_MODE_DLQ:
					ruleSetID, ok := d.getRuleSetIDForRule(rule.Id)
					if !ok {
						return nil, fmt.Errorf("BUG: failed to get rule set id for rule %s", rule.Id)
					}

					if err := d.Plumber.SendRuleNotification(context.Background(), data, rule, ruleSetID); err != nil {
						// TODO: counters here
						return nil, errors.Wrap(err, "failed to send rule notification")
					}

					shouldDrop = true
				default:
					return nil, errors.Errorf("unknown rule failure mode: %s", failCfg.Mode)
				}
			}

			if shouldDrop {
				return nil, ErrMessageDropped
			}

			return data, nil

		case protos.RuleType_RULE_TYPE_CUSTOM:
			// TODO: implement eventually
			return data, nil
		default:
			return nil, errors.Errorf("unknown rule type: %s", rule.Type)
		}
	}

	return data, nil
}

func (d *DataQual) logDryRun(rule *protos.Rule) {
	cfg := rule.GetMatchConfig()
	for _, failCfg := range rule.FailureModeConfigs {
		switch failCfg.Mode {
		case protos.RuleFailureMode_RULE_FAILURE_MODE_REJECT:
			d.Logger.Infof("DRY RUN: Matched rule '%s' with type '%s' on path '%s' and would have rejected the message", rule.Id, rule.Type, cfg.Path)
		case protos.RuleFailureMode_RULE_FAILURE_MODE_TRANSFORM:
			d.Logger.Infof("DRY RUN: Matched rule '%s' with type '%s' on path '%s' and would have transformed the message", rule.Id, rule.Type, cfg.Path)
		case protos.RuleFailureMode_RULE_FAILURE_MODE_ALERT_SLACK:
			d.Logger.Infof("DRY RUN: Matched rule '%s' with type '%s' on path '%s' and would have sent a slack alert", rule.Id, rule.Type, cfg.Path)
		case protos.RuleFailureMode_RULE_FAILURE_MODE_DLQ:
			d.Logger.Infof("DRY RUN: Matched rule '%s' with type '%s' on path '%s' and would have sent the message to the DLQ", rule.Id, rule.Type, cfg.Path)
		default:
			d.Logger.Infof("DRY RUN: Matched rule '%s' with type '%s' on path '%s' and would have done nothing", rule.Id, rule.Type, cfg.Path)
		}
	}
}

func (d *DataQual) watchForRuleUpdates(looper director.Looper) {
	var quit bool

	looper.Loop(func() error {
		if quit {
			// Give looper time to exit
			time.Sleep(time.Millisecond * 50)
			return nil
		}

		select {
		case <-d.ShutdownCtx.Done():
			quit = true
			looper.Quit()
			return nil
		default:
			// NOOP
		}

		if err := d.getRuleUpdates(); err != nil {
			d.Logger.Error("failed to get rule updates:", err)
		}

		return nil
	})
}

func (d *DataQual) getRuleUpdates() error {
	ruleSets, err := d.Plumber.GetRules(context.Background(), d.DataSource)
	if err != nil {
		return errors.Wrap(err, "failed to get rules")
	}

	// First key is mode: producer, consumer
	// Second map key is rule key: kafka topic, rabbit routing/binding key
	// We need a way to look these up O(1)
	rules := make(map[Mode]map[string][]*protos.Rule)
	ruleSetMap := make(map[string]string)

	for _, set := range ruleSets {
		if _, ok := rules[Mode(set.Mode)]; !ok {
			rules[Mode(set.Mode)] = make(map[string][]*protos.Rule)
		}

		for _, rule := range set.Rules {
			if _, ok := rules[Mode(set.Mode)][set.Key]; !ok {
				rules[Mode(set.Mode)][set.Key] = make([]*protos.Rule, 0)
			}

			rules[Mode(set.Mode)][set.Key] = append(rules[Mode(set.Mode)][set.Key], rule)
			ruleSetMap[rule.Id] = set.Id
		}
	}

	d.rulesMtx.Lock()
	d.rules = rules
	d.rulesMtx.Unlock()

	d.ruleSetMtx.Lock()
	d.ruleSetMap = ruleSetMap
	d.ruleSetMtx.Unlock()

	return nil
}

// getRuleSetIDForRule gets the ruleset ID for a given rule ID.
// This is needed when communicating rule failure alerts to Plumber
func (d *DataQual) getRuleSetIDForRule(ruleID string) (string, bool) {
	d.ruleSetMtx.RLock()
	defer d.ruleSetMtx.RUnlock()

	ruleSetID, ok := d.ruleSetMap[ruleID]
	return ruleSetID, ok
}

func (m Mode) String() string {
	switch m {
	case Publish:
		return "publish"
	case Consume:
		return "consume"
	default:
		return "unknown"
	}
}
