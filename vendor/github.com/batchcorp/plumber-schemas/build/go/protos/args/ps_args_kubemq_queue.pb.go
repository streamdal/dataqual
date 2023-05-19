// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: ps_args_kubemq_queue.proto

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

type KubeMQQueueConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Dial string for KubeMQ server',env='PLUMBER_RELAY_KUBEMQ_QUEUE_ADDRESS',default='localhost:50000',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Dial string for KubeMQ server',env='PLUMBER_RELAY_KUBEMQ_QUEUE_ADDRESS',default='localhost:50000',required"`
	// @gotags: kong:"help='Client JWT authentication token',env='PLUMBER_RELAY_KUBEMQ_QUEUE_AUTH_TOKEN'"
	AuthToken string `protobuf:"bytes,2,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty" kong:"help='Client JWT authentication token',env='PLUMBER_RELAY_KUBEMQ_QUEUE_AUTH_TOKEN'"`
	// @gotags: kong:"help='KubeMQ client cert file',env='PLUMBER_RELAY_KUBEMQ_QUEUE_TLS_CLIENT_CERT'"
	TlsClientCert string `protobuf:"bytes,3,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty" kong:"help='KubeMQ client cert file',env='PLUMBER_RELAY_KUBEMQ_QUEUE_TLS_CLIENT_CERT'"`
	// @gotags: kong:"help='KubeMQ client ID',env='PLUMBER_RELAY_KUBEMQ_QUEUE_CLIENT_ID',default=plumber"
	ClientId string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" kong:"help='KubeMQ client ID',env='PLUMBER_RELAY_KUBEMQ_QUEUE_CLIENT_ID',default=plumber"`
}

func (x *KubeMQQueueConn) Reset() {
	*x = KubeMQQueueConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_kubemq_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeMQQueueConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeMQQueueConn) ProtoMessage() {}

func (x *KubeMQQueueConn) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_kubemq_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeMQQueueConn.ProtoReflect.Descriptor instead.
func (*KubeMQQueueConn) Descriptor() ([]byte, []int) {
	return file_ps_args_kubemq_queue_proto_rawDescGZIP(), []int{0}
}

func (x *KubeMQQueueConn) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *KubeMQQueueConn) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *KubeMQQueueConn) GetTlsClientCert() string {
	if x != nil {
		return x.TlsClientCert
	}
	return ""
}

func (x *KubeMQQueueConn) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type KubeMQQueueReadArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='KubeMQ queue name',env='PLUMBER_RELAY_KUBEMQ_QUEUE_QUEUE'"
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='KubeMQ queue name',env='PLUMBER_RELAY_KUBEMQ_QUEUE_QUEUE'"`
}

func (x *KubeMQQueueReadArgs) Reset() {
	*x = KubeMQQueueReadArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_kubemq_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeMQQueueReadArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeMQQueueReadArgs) ProtoMessage() {}

func (x *KubeMQQueueReadArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_kubemq_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeMQQueueReadArgs.ProtoReflect.Descriptor instead.
func (*KubeMQQueueReadArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_kubemq_queue_proto_rawDescGZIP(), []int{1}
}

func (x *KubeMQQueueReadArgs) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

type KubeMQQueueWriteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='KubeMQ queue name'"
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='KubeMQ queue name'"`
}

func (x *KubeMQQueueWriteArgs) Reset() {
	*x = KubeMQQueueWriteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_kubemq_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeMQQueueWriteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeMQQueueWriteArgs) ProtoMessage() {}

func (x *KubeMQQueueWriteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_kubemq_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeMQQueueWriteArgs.ProtoReflect.Descriptor instead.
func (*KubeMQQueueWriteArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_kubemq_queue_proto_rawDescGZIP(), []int{2}
}

func (x *KubeMQQueueWriteArgs) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

var File_ps_args_kubemq_queue_proto protoreflect.FileDescriptor

var file_ps_args_kubemq_queue_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x73, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x6d, 0x71,
	0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x4b, 0x75,
	0x62, 0x65, 0x4d, 0x51, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x6c, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x4b,
	0x75, 0x62, 0x65, 0x4d, 0x51, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x35, 0x0a, 0x14, 0x4b, 0x75, 0x62, 0x65, 0x4d, 0x51, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70,
	0x2f, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x61, 0x72, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ps_args_kubemq_queue_proto_rawDescOnce sync.Once
	file_ps_args_kubemq_queue_proto_rawDescData = file_ps_args_kubemq_queue_proto_rawDesc
)

func file_ps_args_kubemq_queue_proto_rawDescGZIP() []byte {
	file_ps_args_kubemq_queue_proto_rawDescOnce.Do(func() {
		file_ps_args_kubemq_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_ps_args_kubemq_queue_proto_rawDescData)
	})
	return file_ps_args_kubemq_queue_proto_rawDescData
}

var file_ps_args_kubemq_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ps_args_kubemq_queue_proto_goTypes = []interface{}{
	(*KubeMQQueueConn)(nil),      // 0: protos.args.KubeMQQueueConn
	(*KubeMQQueueReadArgs)(nil),  // 1: protos.args.KubeMQQueueReadArgs
	(*KubeMQQueueWriteArgs)(nil), // 2: protos.args.KubeMQQueueWriteArgs
}
var file_ps_args_kubemq_queue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ps_args_kubemq_queue_proto_init() }
func file_ps_args_kubemq_queue_proto_init() {
	if File_ps_args_kubemq_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ps_args_kubemq_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeMQQueueConn); i {
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
		file_ps_args_kubemq_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeMQQueueReadArgs); i {
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
		file_ps_args_kubemq_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeMQQueueWriteArgs); i {
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
			RawDescriptor: file_ps_args_kubemq_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ps_args_kubemq_queue_proto_goTypes,
		DependencyIndexes: file_ps_args_kubemq_queue_proto_depIdxs,
		MessageInfos:      file_ps_args_kubemq_queue_proto_msgTypes,
	}.Build()
	File_ps_args_kubemq_queue_proto = out.File
	file_ps_args_kubemq_queue_proto_rawDesc = nil
	file_ps_args_kubemq_queue_proto_goTypes = nil
	file_ps_args_kubemq_queue_proto_depIdxs = nil
}