// Code generated by generate-merge-relay-opts.go. DO NOT EDIT.

package opts

import (
	"errors"
)

func MergeTunnelOptions(backend string, tunnelOpts *TunnelOptions, createTunnelOpts *CreateTunnelOptions) error {
	switch backend {
	case "activemq":
		tunnelOpts.Activemq = &TunnelGroupActiveMQOptions{Args: createTunnelOpts.Activemq}
	case "awskinesis":
		tunnelOpts.AwsKinesis = &TunnelGroupAWSKinesisOptions{Args: createTunnelOpts.AwsKinesis}
	case "awssns":
		tunnelOpts.AwsSns = &TunnelGroupAWSSNSOptions{Args: createTunnelOpts.AwsSns}
	case "awssqs":
		tunnelOpts.AwsSqs = &TunnelGroupAWSSQSOptions{Args: createTunnelOpts.AwsSqs}
	case "azureeventhub":
		tunnelOpts.AzureEventHub = &TunnelGroupAzureEventHubOptions{Args: createTunnelOpts.AzureEventHub}
	case "azureservicebus":
		tunnelOpts.AzureServiceBus = &TunnelGroupAzureServiceBusOptions{Args: createTunnelOpts.AzureServiceBus}
	case "gcppubsub":
		tunnelOpts.GcpPubsub = &TunnelGroupGCPPubSubOptions{Args: createTunnelOpts.GcpPubsub}
	case "kafka":
		tunnelOpts.Kafka = &TunnelGroupKafkaOptions{Args: createTunnelOpts.Kafka}
	case "kubemqqueue":
		tunnelOpts.KubemqQueue = &TunnelGroupKubeMQQueueOptions{Args: createTunnelOpts.KubemqQueue}
	case "memphis":
		tunnelOpts.Memphis = &TunnelGroupMemphisOptions{Args: createTunnelOpts.Memphis}
	case "mqtt":
		tunnelOpts.Mqtt = &TunnelGroupMQTTOptions{Args: createTunnelOpts.Mqtt}
	case "nats":
		tunnelOpts.Nats = &TunnelGroupNatsOptions{Args: createTunnelOpts.Nats}
	case "natsjetstream":
		tunnelOpts.NatsJetstream = &TunnelGroupNatsJetstreamOptions{Args: createTunnelOpts.NatsJetstream}
	case "natsstreaming":
		tunnelOpts.NatsStreaming = &TunnelGroupNatsStreamingOptions{Args: createTunnelOpts.NatsStreaming}
	case "nsq":
		tunnelOpts.Nsq = &TunnelGroupNSQOptions{Args: createTunnelOpts.Nsq}
	case "pulsar":
		tunnelOpts.Pulsar = &TunnelGroupPulsarOptions{Args: createTunnelOpts.Pulsar}
	case "rabbit":
		tunnelOpts.Rabbit = &TunnelGroupRabbitOptions{Args: createTunnelOpts.Rabbit}
	case "rabbitstreams":
		tunnelOpts.RabbitStreams = &TunnelGroupRabbitStreamsOptions{Args: createTunnelOpts.RabbitStreams}
	case "redispubsub":
		tunnelOpts.RedisPubsub = &TunnelGroupRedisPubSubOptions{Args: createTunnelOpts.RedisPubsub}
	case "redisstreams":
		tunnelOpts.RedisStreams = &TunnelGroupRedisStreamsOptions{Args: createTunnelOpts.RedisStreams}
	default:
		return errors.New("unknown backend")
	}

	return nil
}