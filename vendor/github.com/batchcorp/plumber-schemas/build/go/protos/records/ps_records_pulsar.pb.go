// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: records/ps_records_pulsar.proto

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

type Pulsar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key             string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Topic           string            `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Properties      map[string]string `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RedeliveryCount uint32            `protobuf:"varint,5,opt,name=redelivery_count,json=redeliveryCount,proto3" json:"redelivery_count,omitempty"`
	EventTime       string            `protobuf:"bytes,6,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	IsReplicated    bool              `protobuf:"varint,7,opt,name=is_replicated,json=isReplicated,proto3" json:"is_replicated,omitempty"`
	OrderingKey     string            `protobuf:"bytes,8,opt,name=ordering_key,json=orderingKey,proto3" json:"ordering_key,omitempty"`
	ProducerName    string            `protobuf:"bytes,9,opt,name=producer_name,json=producerName,proto3" json:"producer_name,omitempty"`
	PublishTime     string            `protobuf:"bytes,10,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Timestamp       int64             `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value           []byte            `protobuf:"bytes,12,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Pulsar) Reset() {
	*x = Pulsar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_records_ps_records_pulsar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pulsar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pulsar) ProtoMessage() {}

func (x *Pulsar) ProtoReflect() protoreflect.Message {
	mi := &file_records_ps_records_pulsar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pulsar.ProtoReflect.Descriptor instead.
func (*Pulsar) Descriptor() ([]byte, []int) {
	return file_records_ps_records_pulsar_proto_rawDescGZIP(), []int{0}
}

func (x *Pulsar) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pulsar) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pulsar) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Pulsar) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Pulsar) GetRedeliveryCount() uint32 {
	if x != nil {
		return x.RedeliveryCount
	}
	return 0
}

func (x *Pulsar) GetEventTime() string {
	if x != nil {
		return x.EventTime
	}
	return ""
}

func (x *Pulsar) GetIsReplicated() bool {
	if x != nil {
		return x.IsReplicated
	}
	return false
}

func (x *Pulsar) GetOrderingKey() string {
	if x != nil {
		return x.OrderingKey
	}
	return ""
}

func (x *Pulsar) GetProducerName() string {
	if x != nil {
		return x.ProducerName
	}
	return ""
}

func (x *Pulsar) GetPublishTime() string {
	if x != nil {
		return x.PublishTime
	}
	return ""
}

func (x *Pulsar) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Pulsar) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_records_ps_records_pulsar_proto protoreflect.FileDescriptor

var file_records_ps_records_pulsar_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x70, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x5f, 0x70, 0x75, 0x6c, 0x73, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0xd5, 0x03, 0x0a, 0x06, 0x50, 0x75, 0x6c, 0x73, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x73, 0x61, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72,
	0x70, 0x2f, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_records_ps_records_pulsar_proto_rawDescOnce sync.Once
	file_records_ps_records_pulsar_proto_rawDescData = file_records_ps_records_pulsar_proto_rawDesc
)

func file_records_ps_records_pulsar_proto_rawDescGZIP() []byte {
	file_records_ps_records_pulsar_proto_rawDescOnce.Do(func() {
		file_records_ps_records_pulsar_proto_rawDescData = protoimpl.X.CompressGZIP(file_records_ps_records_pulsar_proto_rawDescData)
	})
	return file_records_ps_records_pulsar_proto_rawDescData
}

var file_records_ps_records_pulsar_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_records_ps_records_pulsar_proto_goTypes = []interface{}{
	(*Pulsar)(nil), // 0: protos.records.Pulsar
	nil,            // 1: protos.records.Pulsar.PropertiesEntry
}
var file_records_ps_records_pulsar_proto_depIdxs = []int32{
	1, // 0: protos.records.Pulsar.properties:type_name -> protos.records.Pulsar.PropertiesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_records_ps_records_pulsar_proto_init() }
func file_records_ps_records_pulsar_proto_init() {
	if File_records_ps_records_pulsar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_records_ps_records_pulsar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pulsar); i {
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
			RawDescriptor: file_records_ps_records_pulsar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_records_ps_records_pulsar_proto_goTypes,
		DependencyIndexes: file_records_ps_records_pulsar_proto_depIdxs,
		MessageInfos:      file_records_ps_records_pulsar_proto_msgTypes,
	}.Build()
	File_records_ps_records_pulsar_proto = out.File
	file_records_ps_records_pulsar_proto_rawDesc = nil
	file_records_ps_records_pulsar_proto_goTypes = nil
	file_records_ps_records_pulsar_proto_depIdxs = nil
}