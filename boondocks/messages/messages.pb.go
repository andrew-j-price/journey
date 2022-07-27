// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: boondocks/messages/messages.proto

package messages

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

type RpsChoice2_Play int32

const (
	RpsChoice2_ROCK     RpsChoice2_Play = 0
	RpsChoice2_PAPER    RpsChoice2_Play = 1
	RpsChoice2_SCISSORS RpsChoice2_Play = 2
)

// Enum value maps for RpsChoice2_Play.
var (
	RpsChoice2_Play_name = map[int32]string{
		0: "ROCK",
		1: "PAPER",
		2: "SCISSORS",
	}
	RpsChoice2_Play_value = map[string]int32{
		"ROCK":     0,
		"PAPER":    1,
		"SCISSORS": 2,
	}
)

func (x RpsChoice2_Play) Enum() *RpsChoice2_Play {
	p := new(RpsChoice2_Play)
	*p = x
	return p
}

func (x RpsChoice2_Play) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpsChoice2_Play) Descriptor() protoreflect.EnumDescriptor {
	return file_boondocks_messages_messages_proto_enumTypes[0].Descriptor()
}

func (RpsChoice2_Play) Type() protoreflect.EnumType {
	return &file_boondocks_messages_messages_proto_enumTypes[0]
}

func (x RpsChoice2_Play) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RpsChoice2_Play.Descriptor instead.
func (RpsChoice2_Play) EnumDescriptor() ([]byte, []int) {
	return file_boondocks_messages_messages_proto_rawDescGZIP(), []int{3, 0}
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boondocks_messages_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boondocks_messages_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_boondocks_messages_messages_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boondocks_messages_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boondocks_messages_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_boondocks_messages_messages_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RpsChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Throw string `protobuf:"bytes,1,opt,name=throw,proto3" json:"throw,omitempty"`
}

func (x *RpsChoice) Reset() {
	*x = RpsChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boondocks_messages_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpsChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpsChoice) ProtoMessage() {}

func (x *RpsChoice) ProtoReflect() protoreflect.Message {
	mi := &file_boondocks_messages_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpsChoice.ProtoReflect.Descriptor instead.
func (*RpsChoice) Descriptor() ([]byte, []int) {
	return file_boondocks_messages_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RpsChoice) GetThrow() string {
	if x != nil {
		return x.Throw
	}
	return ""
}

type RpsChoice2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RpsChoice2) Reset() {
	*x = RpsChoice2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boondocks_messages_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpsChoice2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpsChoice2) ProtoMessage() {}

func (x *RpsChoice2) ProtoReflect() protoreflect.Message {
	mi := &file_boondocks_messages_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpsChoice2.ProtoReflect.Descriptor instead.
func (*RpsChoice2) Descriptor() ([]byte, []int) {
	return file_boondocks_messages_messages_proto_rawDescGZIP(), []int{3}
}

type RpsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameResult string `protobuf:"bytes,1,opt,name=gameResult,proto3" json:"gameResult,omitempty"`
	UserWins   string `protobuf:"bytes,2,opt,name=userWins,proto3" json:"userWins,omitempty"`
	CompWins   string `protobuf:"bytes,3,opt,name=compWins,proto3" json:"compWins,omitempty"`
	Draws      string `protobuf:"bytes,4,opt,name=draws,proto3" json:"draws,omitempty"`
}

func (x *RpsResponse) Reset() {
	*x = RpsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_boondocks_messages_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpsResponse) ProtoMessage() {}

func (x *RpsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boondocks_messages_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpsResponse.ProtoReflect.Descriptor instead.
func (*RpsResponse) Descriptor() ([]byte, []int) {
	return file_boondocks_messages_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RpsResponse) GetGameResult() string {
	if x != nil {
		return x.GameResult
	}
	return ""
}

func (x *RpsResponse) GetUserWins() string {
	if x != nil {
		return x.UserWins
	}
	return ""
}

func (x *RpsResponse) GetCompWins() string {
	if x != nil {
		return x.CompWins
	}
	return ""
}

func (x *RpsResponse) GetDraws() string {
	if x != nil {
		return x.Draws
	}
	return ""
}

var File_boondocks_messages_messages_proto protoreflect.FileDescriptor

var file_boondocks_messages_messages_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x6f, 0x6f, 0x6e, 0x64, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x22, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x52, 0x70, 0x73, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x72,
	0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x22,
	0x37, 0x0a, 0x0a, 0x52, 0x70, 0x73, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x22, 0x29, 0x0a,
	0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x43, 0x4b, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x41, 0x50, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x43,
	0x49, 0x53, 0x53, 0x4f, 0x52, 0x53, 0x10, 0x02, 0x22, 0x7b, 0x0a, 0x0b, 0x52, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x57,
	0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x57,
	0x69, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x57, 0x69, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x57, 0x69, 0x6e, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x72, 0x61, 0x77, 0x73, 0x32, 0x8e, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x07, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x70, 0x73, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x52, 0x70, 0x73, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x15, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x2d, 0x6a, 0x2d, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6e,
	0x64, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_boondocks_messages_messages_proto_rawDescOnce sync.Once
	file_boondocks_messages_messages_proto_rawDescData = file_boondocks_messages_messages_proto_rawDesc
)

func file_boondocks_messages_messages_proto_rawDescGZIP() []byte {
	file_boondocks_messages_messages_proto_rawDescOnce.Do(func() {
		file_boondocks_messages_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_boondocks_messages_messages_proto_rawDescData)
	})
	return file_boondocks_messages_messages_proto_rawDescData
}

var file_boondocks_messages_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_boondocks_messages_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_boondocks_messages_messages_proto_goTypes = []interface{}{
	(RpsChoice2_Play)(0),  // 0: messages.RpsChoice2.Play
	(*HelloRequest)(nil),  // 1: messages.HelloRequest
	(*HelloResponse)(nil), // 2: messages.HelloResponse
	(*RpsChoice)(nil),     // 3: messages.RpsChoice
	(*RpsChoice2)(nil),    // 4: messages.RpsChoice2
	(*RpsResponse)(nil),   // 5: messages.RpsResponse
}
var file_boondocks_messages_messages_proto_depIdxs = []int32{
	1, // 0: messages.BoonService.PerformHelloWorld:input_type -> messages.HelloRequest
	3, // 1: messages.BoonService.PlayRps:input_type -> messages.RpsChoice
	2, // 2: messages.BoonService.PerformHelloWorld:output_type -> messages.HelloResponse
	5, // 3: messages.BoonService.PlayRps:output_type -> messages.RpsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_boondocks_messages_messages_proto_init() }
func file_boondocks_messages_messages_proto_init() {
	if File_boondocks_messages_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_boondocks_messages_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_boondocks_messages_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_boondocks_messages_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpsChoice); i {
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
		file_boondocks_messages_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpsChoice2); i {
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
		file_boondocks_messages_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpsResponse); i {
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
			RawDescriptor: file_boondocks_messages_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_boondocks_messages_messages_proto_goTypes,
		DependencyIndexes: file_boondocks_messages_messages_proto_depIdxs,
		EnumInfos:         file_boondocks_messages_messages_proto_enumTypes,
		MessageInfos:      file_boondocks_messages_messages_proto_msgTypes,
	}.Build()
	File_boondocks_messages_messages_proto = out.File
	file_boondocks_messages_messages_proto_rawDesc = nil
	file_boondocks_messages_messages_proto_goTypes = nil
	file_boondocks_messages_messages_proto_depIdxs = nil
}
