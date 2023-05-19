// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: ps_args_rabbit_streams.proto

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

type RabbitStreamsConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='DSN used to connect to RabbitMQ',default='rabbitmq-stream://guest:guest@localhost:5552',required"
	Dsn string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty" kong:"help='DSN used to connect to RabbitMQ',default='rabbitmq-stream://guest:guest@localhost:5552',required"`
	// @gotags: kong:"help='Enable TLS usage (regardless of DSN)'"
	UseTls bool `protobuf:"varint,2,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Enable TLS usage (regardless of DSN)'"`
	// @gotags: kong:"help='Whether to verify server certificate'"
	TlsSkipVerify bool `protobuf:"varint,3,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty" kong:"help='Whether to verify server certificate'"`
	// @gotags: kong:"help='Username to authenticate to server with',default=guest"
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty" kong:"help='Username to authenticate to server with',default=guest"`
	// @gotags: kong:"help='Password used to authenticate to server with',default=guest"
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty" kong:"help='Password used to authenticate to server with',default=guest"`
	// @gotags: kong:"help='Consumer or producer name to identify as to RabbitMQ',default=plumber,required"
	ClientName string `protobuf:"bytes,6,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty" kong:"help='Consumer or producer name to identify as to RabbitMQ',default=plumber,required"`
}

func (x *RabbitStreamsConn) Reset() {
	*x = RabbitStreamsConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_rabbit_streams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RabbitStreamsConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RabbitStreamsConn) ProtoMessage() {}

func (x *RabbitStreamsConn) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_rabbit_streams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RabbitStreamsConn.ProtoReflect.Descriptor instead.
func (*RabbitStreamsConn) Descriptor() ([]byte, []int) {
	return file_ps_args_rabbit_streams_proto_rawDescGZIP(), []int{0}
}

func (x *RabbitStreamsConn) GetDsn() string {
	if x != nil {
		return x.Dsn
	}
	return ""
}

func (x *RabbitStreamsConn) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

func (x *RabbitStreamsConn) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *RabbitStreamsConn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RabbitStreamsConn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RabbitStreamsConn) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type RabbitStreamsOffsetOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"group=offset,xor=offset"
	SpecificOffset int64 `protobuf:"varint,1,opt,name=specific_offset,json=specificOffset,proto3" json:"specific_offset,omitempty" kong:"group=offset,xor=offset"`
	// @gotags: kong:"group=offset,xor=offset"
	LastOffset bool `protobuf:"varint,2,opt,name=last_offset,json=lastOffset,proto3" json:"last_offset,omitempty" kong:"group=offset,xor=offset"`
	// @gotags: kong:"group=offset,xor=offset"
	LastConsumed bool `protobuf:"varint,3,opt,name=last_consumed,json=lastConsumed,proto3" json:"last_consumed,omitempty" kong:"group=offset,xor=offset"`
	// @gotags: kong:"group=offset,xor=offset"
	FirstOffset bool `protobuf:"varint,4,opt,name=first_offset,json=firstOffset,proto3" json:"first_offset,omitempty" kong:"group=offset,xor=offset"`
	// @gotags: kong:"group=offset,xor=offset"
	NextOffset bool `protobuf:"varint,5,opt,name=next_offset,json=nextOffset,proto3" json:"next_offset,omitempty" kong:"group=offset,xor=offset"`
}

func (x *RabbitStreamsOffsetOptions) Reset() {
	*x = RabbitStreamsOffsetOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_rabbit_streams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RabbitStreamsOffsetOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RabbitStreamsOffsetOptions) ProtoMessage() {}

func (x *RabbitStreamsOffsetOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_rabbit_streams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RabbitStreamsOffsetOptions.ProtoReflect.Descriptor instead.
func (*RabbitStreamsOffsetOptions) Descriptor() ([]byte, []int) {
	return file_ps_args_rabbit_streams_proto_rawDescGZIP(), []int{1}
}

func (x *RabbitStreamsOffsetOptions) GetSpecificOffset() int64 {
	if x != nil {
		return x.SpecificOffset
	}
	return 0
}

func (x *RabbitStreamsOffsetOptions) GetLastOffset() bool {
	if x != nil {
		return x.LastOffset
	}
	return false
}

func (x *RabbitStreamsOffsetOptions) GetLastConsumed() bool {
	if x != nil {
		return x.LastConsumed
	}
	return false
}

func (x *RabbitStreamsOffsetOptions) GetFirstOffset() bool {
	if x != nil {
		return x.FirstOffset
	}
	return false
}

func (x *RabbitStreamsOffsetOptions) GetNextOffset() bool {
	if x != nil {
		return x.NextOffset
	}
	return false
}

type RabbitStreamsReadArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Stream name',required"
	Stream string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty" kong:"help='Stream name',required"`
	// @gotags: kong:"help='Declare the stream if it does not exist'"
	DeclareStream bool `protobuf:"varint,2,opt,name=declare_stream,json=declareStream,proto3" json:"declare_stream,omitempty" kong:"help='Declare the stream if it does not exist'"`
	// @gotags: kong:"help='Stream capacity to declare (required if declare_stream is true; ex: 1024k; 10mb; 3gb)'"
	DeclareStreamSize string `protobuf:"bytes,3,opt,name=declare_stream_size,json=declareStreamSize,proto3" json:"declare_stream_size,omitempty" kong:"help='Stream capacity to declare (required if declare_stream is true; ex: 1024k; 10mb; 3gb)'"`
	// TODO: Will this break? Might just need to be a string.
	// @gotags: kong:"help='Offset to start reading at',embed"
	OffsetOptions *RabbitStreamsOffsetOptions `protobuf:"bytes,4,opt,name=offset_options,json=offsetOptions,proto3" json:"offset_options,omitempty" kong:"help='Offset to start reading at',embed"`
}

func (x *RabbitStreamsReadArgs) Reset() {
	*x = RabbitStreamsReadArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_rabbit_streams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RabbitStreamsReadArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RabbitStreamsReadArgs) ProtoMessage() {}

func (x *RabbitStreamsReadArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_rabbit_streams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RabbitStreamsReadArgs.ProtoReflect.Descriptor instead.
func (*RabbitStreamsReadArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_rabbit_streams_proto_rawDescGZIP(), []int{2}
}

func (x *RabbitStreamsReadArgs) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *RabbitStreamsReadArgs) GetDeclareStream() bool {
	if x != nil {
		return x.DeclareStream
	}
	return false
}

func (x *RabbitStreamsReadArgs) GetDeclareStreamSize() string {
	if x != nil {
		return x.DeclareStreamSize
	}
	return ""
}

func (x *RabbitStreamsReadArgs) GetOffsetOptions() *RabbitStreamsOffsetOptions {
	if x != nil {
		return x.OffsetOptions
	}
	return nil
}

type RabbitStreamsWriteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Stream name',required"
	Stream string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty" kong:"help='Stream name',required"`
	// @gotags: kong:"help='Declare the stream if it does not exist'"
	DeclareStream bool `protobuf:"varint,2,opt,name=declare_stream,json=declareStream,proto3" json:"declare_stream,omitempty" kong:"help='Declare the stream if it does not exist'"`
	// @gotags: kong:"help='Stream capacity to declare (required if declare_stream is true; ex: 1024k, 10mb, 3gb',default=10mb"
	DeclareStreamSize string `protobuf:"bytes,3,opt,name=declare_stream_size,json=declareStreamSize,proto3" json:"declare_stream_size,omitempty" kong:"help='Stream capacity to declare (required if declare_stream is true; ex: 1024k, 10mb, 3gb',default=10mb"`
}

func (x *RabbitStreamsWriteArgs) Reset() {
	*x = RabbitStreamsWriteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_rabbit_streams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RabbitStreamsWriteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RabbitStreamsWriteArgs) ProtoMessage() {}

func (x *RabbitStreamsWriteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_rabbit_streams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RabbitStreamsWriteArgs.ProtoReflect.Descriptor instead.
func (*RabbitStreamsWriteArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_rabbit_streams_proto_rawDescGZIP(), []int{3}
}

func (x *RabbitStreamsWriteArgs) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *RabbitStreamsWriteArgs) GetDeclareStream() bool {
	if x != nil {
		return x.DeclareStream
	}
	return false
}

func (x *RabbitStreamsWriteArgs) GetDeclareStreamSize() string {
	if x != nil {
		return x.DeclareStreamSize
	}
	return ""
}

var File_ps_args_rabbit_streams_proto protoreflect.FileDescriptor

var file_ps_args_rabbit_streams_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x73, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x5f, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x11,
	0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x43, 0x6f, 0x6e,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x73, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcf, 0x01,
	0x0a, 0x1a, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0xd6, 0x01, 0x0a, 0x15, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x6c,
	0x61, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4e, 0x0a, 0x0e, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67, 0x73, 0x2e, 0x52,
	0x61, 0x62, 0x62, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x52, 0x61, 0x62,
	0x62, 0x69, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x64,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x69,
	0x7a, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x70, 0x6c, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x72, 0x67, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ps_args_rabbit_streams_proto_rawDescOnce sync.Once
	file_ps_args_rabbit_streams_proto_rawDescData = file_ps_args_rabbit_streams_proto_rawDesc
)

func file_ps_args_rabbit_streams_proto_rawDescGZIP() []byte {
	file_ps_args_rabbit_streams_proto_rawDescOnce.Do(func() {
		file_ps_args_rabbit_streams_proto_rawDescData = protoimpl.X.CompressGZIP(file_ps_args_rabbit_streams_proto_rawDescData)
	})
	return file_ps_args_rabbit_streams_proto_rawDescData
}

var file_ps_args_rabbit_streams_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ps_args_rabbit_streams_proto_goTypes = []interface{}{
	(*RabbitStreamsConn)(nil),          // 0: protos.args.RabbitStreamsConn
	(*RabbitStreamsOffsetOptions)(nil), // 1: protos.args.RabbitStreamsOffsetOptions
	(*RabbitStreamsReadArgs)(nil),      // 2: protos.args.RabbitStreamsReadArgs
	(*RabbitStreamsWriteArgs)(nil),     // 3: protos.args.RabbitStreamsWriteArgs
}
var file_ps_args_rabbit_streams_proto_depIdxs = []int32{
	1, // 0: protos.args.RabbitStreamsReadArgs.offset_options:type_name -> protos.args.RabbitStreamsOffsetOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ps_args_rabbit_streams_proto_init() }
func file_ps_args_rabbit_streams_proto_init() {
	if File_ps_args_rabbit_streams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ps_args_rabbit_streams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RabbitStreamsConn); i {
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
		file_ps_args_rabbit_streams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RabbitStreamsOffsetOptions); i {
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
		file_ps_args_rabbit_streams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RabbitStreamsReadArgs); i {
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
		file_ps_args_rabbit_streams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RabbitStreamsWriteArgs); i {
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
			RawDescriptor: file_ps_args_rabbit_streams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ps_args_rabbit_streams_proto_goTypes,
		DependencyIndexes: file_ps_args_rabbit_streams_proto_depIdxs,
		MessageInfos:      file_ps_args_rabbit_streams_proto_msgTypes,
	}.Build()
	File_ps_args_rabbit_streams_proto = out.File
	file_ps_args_rabbit_streams_proto_rawDesc = nil
	file_ps_args_rabbit_streams_proto_goTypes = nil
	file_ps_args_rabbit_streams_proto_depIdxs = nil
}