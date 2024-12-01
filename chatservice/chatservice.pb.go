// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: chatservice/chatservice.proto

package chatservice

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

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	mi := &file_chatservice_chatservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chatservice_chatservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_chatservice_chatservice_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_chatservice_chatservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatservice_chatservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_chatservice_chatservice_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RenderOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *RenderOption) Reset() {
	*x = RenderOption{}
	mi := &file_chatservice_chatservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenderOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderOption) ProtoMessage() {}

func (x *RenderOption) ProtoReflect() protoreflect.Message {
	mi := &file_chatservice_chatservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderOption.ProtoReflect.Descriptor instead.
func (*RenderOption) Descriptor() ([]byte, []int) {
	return file_chatservice_chatservice_proto_rawDescGZIP(), []int{2}
}

func (x *RenderOption) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *RenderOption) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Pixel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X     int32  `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y     int32  `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Color *Color `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Pixel) Reset() {
	*x = Pixel{}
	mi := &file_chatservice_chatservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pixel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pixel) ProtoMessage() {}

func (x *Pixel) ProtoReflect() protoreflect.Message {
	mi := &file_chatservice_chatservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pixel.ProtoReflect.Descriptor instead.
func (*Pixel) Descriptor() ([]byte, []int) {
	return file_chatservice_chatservice_proto_rawDescGZIP(), []int{3}
}

func (x *Pixel) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Pixel) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Pixel) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R int32 `protobuf:"varint,1,opt,name=r,proto3" json:"r,omitempty"`
	G int32 `protobuf:"varint,2,opt,name=g,proto3" json:"g,omitempty"`
	B int32 `protobuf:"varint,3,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	mi := &file_chatservice_chatservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_chatservice_chatservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_chatservice_chatservice_proto_rawDescGZIP(), []int{4}
}

func (x *Color) GetR() int32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *Color) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *Color) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_chatservice_chatservice_proto protoreflect.FileDescriptor

var file_chatservice_chatservice_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0d, 0x0a, 0x0b,
	0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x31, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x67, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x62, 0x32, 0x89, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4c, 0x75, 0x6e, 0x61, 0x72, 0x69, 0x73, 0x6e, 0x69, 0x61, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatservice_chatservice_proto_rawDescOnce sync.Once
	file_chatservice_chatservice_proto_rawDescData = file_chatservice_chatservice_proto_rawDesc
)

func file_chatservice_chatservice_proto_rawDescGZIP() []byte {
	file_chatservice_chatservice_proto_rawDescOnce.Do(func() {
		file_chatservice_chatservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatservice_chatservice_proto_rawDescData)
	})
	return file_chatservice_chatservice_proto_rawDescData
}

var file_chatservice_chatservice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chatservice_chatservice_proto_goTypes = []any{
	(*PingMessage)(nil),  // 0: chatservice.PingMessage
	(*PingResponse)(nil), // 1: chatservice.PingResponse
	(*RenderOption)(nil), // 2: chatservice.RenderOption
	(*Pixel)(nil),        // 3: chatservice.Pixel
	(*Color)(nil),        // 4: chatservice.Color
}
var file_chatservice_chatservice_proto_depIdxs = []int32{
	4, // 0: chatservice.Pixel.color:type_name -> chatservice.Color
	0, // 1: chatservice.ChatService.Ping:input_type -> chatservice.PingMessage
	2, // 2: chatservice.ChatService.Render:input_type -> chatservice.RenderOption
	1, // 3: chatservice.ChatService.Ping:output_type -> chatservice.PingResponse
	3, // 4: chatservice.ChatService.Render:output_type -> chatservice.Pixel
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chatservice_chatservice_proto_init() }
func file_chatservice_chatservice_proto_init() {
	if File_chatservice_chatservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chatservice_chatservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatservice_chatservice_proto_goTypes,
		DependencyIndexes: file_chatservice_chatservice_proto_depIdxs,
		MessageInfos:      file_chatservice_chatservice_proto_msgTypes,
	}.Build()
	File_chatservice_chatservice_proto = out.File
	file_chatservice_chatservice_proto_rawDesc = nil
	file_chatservice_chatservice_proto_goTypes = nil
	file_chatservice_chatservice_proto_depIdxs = nil
}
