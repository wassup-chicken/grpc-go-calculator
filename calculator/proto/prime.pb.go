// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.1
// source: prime.proto

package proto

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

type PrimeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Input         int64                  `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	mi := &file_prime_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prime_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_prime_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeRequest) GetInput() int64 {
	if x != nil {
		return x.Input
	}
	return 0
}

type PrimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        int64                  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrimeResponse) Reset() {
	*x = PrimeResponse{}
	mi := &file_prime_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeResponse) ProtoMessage() {}

func (x *PrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prime_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeResponse.ProtoReflect.Descriptor instead.
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return file_prime_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_prime_proto protoreflect.FileDescriptor

var file_prime_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22,
	0x27, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x73, 0x73, 0x75, 0x70, 0x2d, 0x63, 0x68,
	0x69, 0x63, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2d, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prime_proto_rawDescOnce sync.Once
	file_prime_proto_rawDescData = file_prime_proto_rawDesc
)

func file_prime_proto_rawDescGZIP() []byte {
	file_prime_proto_rawDescOnce.Do(func() {
		file_prime_proto_rawDescData = protoimpl.X.CompressGZIP(file_prime_proto_rawDescData)
	})
	return file_prime_proto_rawDescData
}

var file_prime_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prime_proto_goTypes = []any{
	(*PrimeRequest)(nil),  // 0: calculator.PrimeRequest
	(*PrimeResponse)(nil), // 1: calculator.PrimeResponse
}
var file_prime_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prime_proto_init() }
func file_prime_proto_init() {
	if File_prime_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prime_proto_goTypes,
		DependencyIndexes: file_prime_proto_depIdxs,
		MessageInfos:      file_prime_proto_msgTypes,
	}.Build()
	File_prime_proto = out.File
	file_prime_proto_rawDesc = nil
	file_prime_proto_goTypes = nil
	file_prime_proto_depIdxs = nil
}
