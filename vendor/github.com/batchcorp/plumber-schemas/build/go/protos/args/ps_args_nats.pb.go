// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: ps_args_nats.proto

package args

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NatsConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Dial string for NATS server. Ex: nats://localhost:4222',default='nats://localhost:4222',env='PLUMBER_RELAY_NATS_DSN'"
	Dsn string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty" kong:"help='Dial string for NATS server. Ex: nats://localhost:4222',default='nats://localhost:4222',env='PLUMBER_RELAY_NATS_DSN'"`
	// @gotags: kong:"help='NATS .creds file containing authentication credentials',env='PLUMBER_RELAY_NATS_CREDENTIALS'"
	UserCredentials string `protobuf:"bytes,2,opt,name=user_credentials,json=userCredentials,proto3" json:"user_credentials,omitempty" kong:"help='NATS .creds file containing authentication credentials',env='PLUMBER_RELAY_NATS_CREDENTIALS'"`
	// @gotags: kong:"embed"
	TlsOptions *NatsTLSOptions `protobuf:"bytes,3,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty" kong:"embed"`
	// @gotags: kong:"help='File containing NATS NKey',env='PLUMBER_RELAY_NATS_NKEY'"
	Nkey string `protobuf:"bytes,4,opt,name=nkey,proto3" json:"nkey,omitempty" kong:"help='File containing NATS NKey',env='PLUMBER_RELAY_NATS_NKEY'"`
}

func (x *NatsConn) Reset() {
	*x = NatsConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_nats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsConn) ProtoMessage() {}

func (x *NatsConn) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_nats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsConn.ProtoReflect.Descriptor instead.
func (*NatsConn) Descriptor() ([]byte, []int) {
	return file_ps_args_nats_proto_rawDescGZIP(), []int{0}
}

func (x *NatsConn) GetDsn() string {
	if x != nil {
		return x.Dsn
	}
	return ""
}

func (x *NatsConn) GetUserCredentials() string {
	if x != nil {
		return x.UserCredentials
	}
	return ""
}

func (x *NatsConn) GetTlsOptions() *NatsTLSOptions {
	if x != nil {
		return x.TlsOptions
	}
	return nil
}

func (x *NatsConn) GetNkey() string {
	if x != nil {
		return x.Nkey
	}
	return ""
}

type NatsTLSOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Whether to verify server certificate',env='PLUMBER_RELAY_NATS_SKIP_VERIFY_TLS'"
	TlsSkipVerify bool `protobuf:"varint,1,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Whether to verify server certificate',env='PLUMBER_RELAY_NATS_SKIP_VERIFY_TLS'"`
	// @gotags: kong:"help='CA file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CA_CERT'"
	TlsCaCert string `protobuf:"bytes,2,opt,name=tls_ca_cert,json=tlsCaCert,proto3" json:"tls_ca_cert,omitempty" kong:"help='CA file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CA_CERT'"`
	// @gotags: kong:"help='Client cert file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_CERT'"
	TlsClientCert string `protobuf:"bytes,3,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty" kong:"help='Client cert file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_CERT'"`
	// @gotags: kong:"help='client key file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_KEY'"
	TlsClientKey string `protobuf:"bytes,4,opt,name=tls_client_key,json=tlsClientKey,proto3" json:"tls_client_key,omitempty" kong:"help='client key file (only needed if addr is tls://)',env='PLUMBER_RELAY_NATS_TLS_CLIENT_KEY'"`
	// @gotags: kong:"help='Enable TLS',env='PLUMBER_RELAY_NATS_USE_TLS'"
	UseTls bool `protobuf:"varint,5,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Enable TLS',env='PLUMBER_RELAY_NATS_USE_TLS'"`
}

func (x *NatsTLSOptions) Reset() {
	*x = NatsTLSOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_nats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsTLSOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsTLSOptions) ProtoMessage() {}

func (x *NatsTLSOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_nats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsTLSOptions.ProtoReflect.Descriptor instead.
func (*NatsTLSOptions) Descriptor() ([]byte, []int) {
	return file_ps_args_nats_proto_rawDescGZIP(), []int{1}
}

func (x *NatsTLSOptions) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *NatsTLSOptions) GetTlsCaCert() string {
	if x != nil {
		return x.TlsCaCert
	}
	return ""
}

func (x *NatsTLSOptions) GetTlsClientCert() string {
	if x != nil {
		return x.TlsClientCert
	}
	return ""
}

func (x *NatsTLSOptions) GetTlsClientKey() string {
	if x != nil {
		return x.TlsClientKey
	}
	return ""
}

func (x *NatsTLSOptions) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

type NatsReadArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='NATS Subject. Ex: foo.bar.*',env='PLUMBER_RELAY_NATS_SUBJECT'"
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" kong:"help='NATS Subject. Ex: foo.bar.*',env='PLUMBER_RELAY_NATS_SUBJECT'"`
}

func (x *NatsReadArgs) Reset() {
	*x = NatsReadArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_nats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsReadArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsReadArgs) ProtoMessage() {}

func (x *NatsReadArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_nats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsReadArgs.ProtoReflect.Descriptor instead.
func (*NatsReadArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_nats_proto_rawDescGZIP(), []int{2}
}

func (x *NatsReadArgs) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type NatsWriteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='NATS Subject. Ex: foo.bar.*'"
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" kong:"help='NATS Subject. Ex: foo.bar.*'"`
}

func (x *NatsWriteArgs) Reset() {
	*x = NatsWriteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_nats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsWriteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsWriteArgs) ProtoMessage() {}

func (x *NatsWriteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_nats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsWriteArgs.ProtoReflect.Descriptor instead.
func (*NatsWriteArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_nats_proto_rawDescGZIP(), []int{3}
}

func (x *NatsWriteArgs) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

var File_ps_args_nats_proto protoreflect.FileDescriptor

var file_ps_args_nats_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x73, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x5f, 0x6e, 0x61, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67,
	0x73, 0x22, 0x99, 0x01, 0x0a, 0x08, 0x4e, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x73, 0x6e,
	0x12, 0x29, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x74,
	0x6c, 0x73, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67, 0x73, 0x2e, 0x4e,
	0x61, 0x74, 0x73, 0x54, 0x4c, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x74,
	0x6c, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6b, 0x65, 0x79, 0x22, 0xbf, 0x01,
	0x0a, 0x0e, 0x4e, 0x61, 0x74, 0x73, 0x54, 0x4c, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53, 0x6b,
	0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x6c, 0x73, 0x5f,
	0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6c, 0x73, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6c, 0x73, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73, 0x65, 0x54, 0x6c, 0x73, 0x22,
	0x28, 0x0a, 0x0c, 0x4e, 0x61, 0x74, 0x73, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x29, 0x0a, 0x0d, 0x4e, 0x61, 0x74,
	0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x70, 0x6c, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x72, 0x67,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ps_args_nats_proto_rawDescOnce sync.Once
	file_ps_args_nats_proto_rawDescData = file_ps_args_nats_proto_rawDesc
)

func file_ps_args_nats_proto_rawDescGZIP() []byte {
	file_ps_args_nats_proto_rawDescOnce.Do(func() {
		file_ps_args_nats_proto_rawDescData = protoimpl.X.CompressGZIP(file_ps_args_nats_proto_rawDescData)
	})
	return file_ps_args_nats_proto_rawDescData
}

var file_ps_args_nats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ps_args_nats_proto_goTypes = []interface{}{
	(*NatsConn)(nil),       // 0: protos.args.NatsConn
	(*NatsTLSOptions)(nil), // 1: protos.args.NatsTLSOptions
	(*NatsReadArgs)(nil),   // 2: protos.args.NatsReadArgs
	(*NatsWriteArgs)(nil),  // 3: protos.args.NatsWriteArgs
}
var file_ps_args_nats_proto_depIdxs = []int32{
	1, // 0: protos.args.NatsConn.tls_options:type_name -> protos.args.NatsTLSOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ps_args_nats_proto_init() }
func file_ps_args_nats_proto_init() {
	if File_ps_args_nats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ps_args_nats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatsConn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ps_args_nats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatsTLSOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ps_args_nats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatsReadArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ps_args_nats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatsWriteArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ps_args_nats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ps_args_nats_proto_goTypes,
		DependencyIndexes: file_ps_args_nats_proto_depIdxs,
		MessageInfos:      file_ps_args_nats_proto_msgTypes,
	}.Build()
	File_ps_args_nats_proto = out.File
	file_ps_args_nats_proto_rawDesc = nil
	file_ps_args_nats_proto_goTypes = nil
	file_ps_args_nats_proto_depIdxs = nil
}