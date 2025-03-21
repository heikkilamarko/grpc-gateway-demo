// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: calculator/v1/calculator.proto

package calculator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             float32                `protobuf:"fixed32,1,opt,name=a,proto3" json:"a,omitempty"`
	B             float32                `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddRequest) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float32                `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type SubtractRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             float32                `protobuf:"fixed32,1,opt,name=a,proto3" json:"a,omitempty"`
	B             float32                `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubtractRequest) Reset() {
	*x = SubtractRequest{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractRequest) ProtoMessage() {}

func (x *SubtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractRequest.ProtoReflect.Descriptor instead.
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *SubtractRequest) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SubtractRequest) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

type SubtractResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float32                `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubtractResponse) Reset() {
	*x = SubtractResponse{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubtractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractResponse) ProtoMessage() {}

func (x *SubtractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractResponse.ProtoReflect.Descriptor instead.
func (*SubtractResponse) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *SubtractResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_v1_calculator_proto protoreflect.FileDescriptor

var file_calculator_v1_calculator_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22,
	0x28, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x22, 0x25, 0x0a, 0x0b, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x22,
	0x2a, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa2, 0x01, 0x0a, 0x11,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0xac, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0xa2, 0x02, 0x03, 0x43,
	0x58, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x19, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_calculator_v1_calculator_proto_rawDescOnce sync.Once
	file_calculator_v1_calculator_proto_rawDescData []byte
)

func file_calculator_v1_calculator_proto_rawDescGZIP() []byte {
	file_calculator_v1_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_v1_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_v1_calculator_proto_rawDesc), len(file_calculator_v1_calculator_proto_rawDesc)))
	})
	return file_calculator_v1_calculator_proto_rawDescData
}

var file_calculator_v1_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calculator_v1_calculator_proto_goTypes = []any{
	(*AddRequest)(nil),       // 0: calculator.v1.AddRequest
	(*AddResponse)(nil),      // 1: calculator.v1.AddResponse
	(*SubtractRequest)(nil),  // 2: calculator.v1.SubtractRequest
	(*SubtractResponse)(nil), // 3: calculator.v1.SubtractResponse
}
var file_calculator_v1_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.v1.CalculatorService.Add:input_type -> calculator.v1.AddRequest
	2, // 1: calculator.v1.CalculatorService.Subtract:input_type -> calculator.v1.SubtractRequest
	1, // 2: calculator.v1.CalculatorService.Add:output_type -> calculator.v1.AddResponse
	3, // 3: calculator.v1.CalculatorService.Subtract:output_type -> calculator.v1.SubtractResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_v1_calculator_proto_init() }
func file_calculator_v1_calculator_proto_init() {
	if File_calculator_v1_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_v1_calculator_proto_rawDesc), len(file_calculator_v1_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_v1_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_v1_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_v1_calculator_proto_msgTypes,
	}.Build()
	File_calculator_v1_calculator_proto = out.File
	file_calculator_v1_calculator_proto_goTypes = nil
	file_calculator_v1_calculator_proto_depIdxs = nil
}
