// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: records/ps_records_activemq.proto

package records

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

type ActiveMQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination    string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	ContentType    string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	SubscriptionId string `protobuf:"bytes,3,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	Timestamp      int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value          []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ActiveMQ) Reset() {
	*x = ActiveMQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_ps_records_activemq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveMQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveMQ) ProtoMessage() {}

func (x *ActiveMQ) ProtoReflect() protoreflect.Message {
	mi := &file_records_ps_records_activemq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveMQ.ProtoReflect.Descriptor instead.
func (*ActiveMQ) Descriptor() ([]byte, []int) {
	return file_records_ps_records_activemq_proto_rawDescGZIP(), []int{0}
}

func (x *ActiveMQ) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *ActiveMQ) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ActiveMQ) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *ActiveMQ) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ActiveMQ) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_records_ps_records_activemq_proto protoreflect.FileDescriptor

var file_records_ps_records_activemq_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x70, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x6d, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x51,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x70, 0x6c, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_records_ps_records_activemq_proto_rawDescOnce sync.Once
	file_records_ps_records_activemq_proto_rawDescData = file_records_ps_records_activemq_proto_rawDesc
)

func file_records_ps_records_activemq_proto_rawDescGZIP() []byte {
	file_records_ps_records_activemq_proto_rawDescOnce.Do(func() {
		file_records_ps_records_activemq_proto_rawDescData = protoimpl.X.CompressGZIP(file_records_ps_records_activemq_proto_rawDescData)
	})
	return file_records_ps_records_activemq_proto_rawDescData
}

var file_records_ps_records_activemq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_records_ps_records_activemq_proto_goTypes = []interface{}{
	(*ActiveMQ)(nil), // 0: protos.records.ActiveMQ
}
var file_records_ps_records_activemq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_records_ps_records_activemq_proto_init() }
func file_records_ps_records_activemq_proto_init() {
	if File_records_ps_records_activemq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_records_ps_records_activemq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveMQ); i {
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
			RawDescriptor: file_records_ps_records_activemq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_records_ps_records_activemq_proto_goTypes,
		DependencyIndexes: file_records_ps_records_activemq_proto_depIdxs,
		MessageInfos:      file_records_ps_records_activemq_proto_msgTypes,
	}.Build()
	File_records_ps_records_activemq_proto = out.File
	file_records_ps_records_activemq_proto_rawDesc = nil
	file_records_ps_records_activemq_proto_goTypes = nil
	file_records_ps_records_activemq_proto_depIdxs = nil
}