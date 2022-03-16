// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: fixtures/core.proto

package fixtures

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnumNotAnnotated int32

const (
	EnumNotAnnotated_zero_value              EnumNotAnnotated = 0
	EnumNotAnnotated_released                EnumNotAnnotated = 2
	EnumNotAnnotated_previewed_then_released EnumNotAnnotated = 4
)

// Enum value maps for EnumNotAnnotated.
var (
	EnumNotAnnotated_name = map[int32]string{
		0: "zero_value",
		2: "released",
		4: "previewed_then_released",
	}
	EnumNotAnnotated_value = map[string]int32{
		"zero_value":              0,
		"released":                2,
		"previewed_then_released": 4,
	}
)

func (x EnumNotAnnotated) Enum() *EnumNotAnnotated {
	p := new(EnumNotAnnotated)
	*p = x
	return p
}

func (x EnumNotAnnotated) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumNotAnnotated) Descriptor() protoreflect.EnumDescriptor {
	return file_fixtures_core_proto_enumTypes[0].Descriptor()
}

func (EnumNotAnnotated) Type() protoreflect.EnumType {
	return &file_fixtures_core_proto_enumTypes[0]
}

func (x EnumNotAnnotated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumNotAnnotated.Descriptor instead.
func (EnumNotAnnotated) EnumDescriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{0}
}

type EmptyMessageReleased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessageReleased) Reset() {
	*x = EmptyMessageReleased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessageReleased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessageReleased) ProtoMessage() {}

func (x *EmptyMessageReleased) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessageReleased.ProtoReflect.Descriptor instead.
func (*EmptyMessageReleased) Descriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{0}
}

type EmptyMessagePreviewedThenReleased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessagePreviewedThenReleased) Reset() {
	*x = EmptyMessagePreviewedThenReleased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessagePreviewedThenReleased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessagePreviewedThenReleased) ProtoMessage() {}

func (x *EmptyMessagePreviewedThenReleased) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessagePreviewedThenReleased.ProtoReflect.Descriptor instead.
func (*EmptyMessagePreviewedThenReleased) Descriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{1}
}

type MessageWithReleasedField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Released int32 `protobuf:"varint,1,opt,name=released,proto3" json:"released,omitempty"`
}

func (x *MessageWithReleasedField) Reset() {
	*x = MessageWithReleasedField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithReleasedField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithReleasedField) ProtoMessage() {}

func (x *MessageWithReleasedField) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithReleasedField.ProtoReflect.Descriptor instead.
func (*MessageWithReleasedField) Descriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{2}
}

func (x *MessageWithReleasedField) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

type MessageWithReleasedOneofItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OneofWithItem:
	//	*MessageWithReleasedOneofItem_ReleasedOneofItem
	OneofWithItem isMessageWithReleasedOneofItem_OneofWithItem `protobuf_oneof:"oneof_with_item"`
}

func (x *MessageWithReleasedOneofItem) Reset() {
	*x = MessageWithReleasedOneofItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithReleasedOneofItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithReleasedOneofItem) ProtoMessage() {}

func (x *MessageWithReleasedOneofItem) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithReleasedOneofItem.ProtoReflect.Descriptor instead.
func (*MessageWithReleasedOneofItem) Descriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{3}
}

func (m *MessageWithReleasedOneofItem) GetOneofWithItem() isMessageWithReleasedOneofItem_OneofWithItem {
	if m != nil {
		return m.OneofWithItem
	}
	return nil
}

func (x *MessageWithReleasedOneofItem) GetReleasedOneofItem() int32 {
	if x, ok := x.GetOneofWithItem().(*MessageWithReleasedOneofItem_ReleasedOneofItem); ok {
		return x.ReleasedOneofItem
	}
	return 0
}

type isMessageWithReleasedOneofItem_OneofWithItem interface {
	isMessageWithReleasedOneofItem_OneofWithItem()
}

type MessageWithReleasedOneofItem_ReleasedOneofItem struct {
	ReleasedOneofItem int32 `protobuf:"varint,2,opt,name=released_oneof_item,json=releasedOneofItem,proto3,oneof"`
}

func (*MessageWithReleasedOneofItem_ReleasedOneofItem) isMessageWithReleasedOneofItem_OneofWithItem() {
}

type MessageWithOneofWithMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OneofWithMessage:
	//	*MessageWithOneofWithMessages_MessageWithReleaseAnnotation
	OneofWithMessage isMessageWithOneofWithMessages_OneofWithMessage `protobuf_oneof:"oneof_with_message"`
}

func (x *MessageWithOneofWithMessages) Reset() {
	*x = MessageWithOneofWithMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOneofWithMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOneofWithMessages) ProtoMessage() {}

func (x *MessageWithOneofWithMessages) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOneofWithMessages.ProtoReflect.Descriptor instead.
func (*MessageWithOneofWithMessages) Descriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{4}
}

func (m *MessageWithOneofWithMessages) GetOneofWithMessage() isMessageWithOneofWithMessages_OneofWithMessage {
	if m != nil {
		return m.OneofWithMessage
	}
	return nil
}

func (x *MessageWithOneofWithMessages) GetMessageWithReleaseAnnotation() *MessageWithReleasedField {
	if x, ok := x.GetOneofWithMessage().(*MessageWithOneofWithMessages_MessageWithReleaseAnnotation); ok {
		return x.MessageWithReleaseAnnotation
	}
	return nil
}

type isMessageWithOneofWithMessages_OneofWithMessage interface {
	isMessageWithOneofWithMessages_OneofWithMessage()
}

type MessageWithOneofWithMessages_MessageWithReleaseAnnotation struct {
	MessageWithReleaseAnnotation *MessageWithReleasedField `protobuf:"bytes,3,opt,name=message_with_release_annotation,json=messageWithReleaseAnnotation,proto3,oneof"`
}

func (*MessageWithOneofWithMessages_MessageWithReleaseAnnotation) isMessageWithOneofWithMessages_OneofWithMessage() {
}

type MessageNotAnnotated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Released              int32 `protobuf:"varint,2,opt,name=released,proto3" json:"released,omitempty"`
	PreviewedThenReleased int32 `protobuf:"varint,4,opt,name=previewed_then_released,json=previewedThenReleased,proto3" json:"previewed_then_released,omitempty"`
	// Types that are assignable to NotAnnotatedOneof:
	//	*MessageNotAnnotated_OneofItemNotAnnotated
	NotAnnotatedOneof isMessageNotAnnotated_NotAnnotatedOneof `protobuf_oneof:"not_annotated_oneof"`
}

func (x *MessageNotAnnotated) Reset() {
	*x = MessageNotAnnotated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageNotAnnotated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageNotAnnotated) ProtoMessage() {}

func (x *MessageNotAnnotated) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageNotAnnotated.ProtoReflect.Descriptor instead.
func (*MessageNotAnnotated) Descriptor() ([]byte, []int) {
	return file_fixtures_core_proto_rawDescGZIP(), []int{5}
}

func (x *MessageNotAnnotated) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

func (x *MessageNotAnnotated) GetPreviewedThenReleased() int32 {
	if x != nil {
		return x.PreviewedThenReleased
	}
	return 0
}

func (m *MessageNotAnnotated) GetNotAnnotatedOneof() isMessageNotAnnotated_NotAnnotatedOneof {
	if m != nil {
		return m.NotAnnotatedOneof
	}
	return nil
}

func (x *MessageNotAnnotated) GetOneofItemNotAnnotated() int32 {
	if x, ok := x.GetNotAnnotatedOneof().(*MessageNotAnnotated_OneofItemNotAnnotated); ok {
		return x.OneofItemNotAnnotated
	}
	return 0
}

type isMessageNotAnnotated_NotAnnotatedOneof interface {
	isMessageNotAnnotated_NotAnnotatedOneof()
}

type MessageNotAnnotated_OneofItemNotAnnotated struct {
	OneofItemNotAnnotated int32 `protobuf:"varint,8,opt,name=oneof_item_not_annotated,json=oneofItemNotAnnotated,proto3,oneof"`
}

func (*MessageNotAnnotated_OneofItemNotAnnotated) isMessageNotAnnotated_NotAnnotatedOneof() {}

var File_fixtures_core_proto protoreflect.FileDescriptor

var file_fixtures_core_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a,
	0x14, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x22, 0x23, 0x0a, 0x21, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68,
	0x65, 0x6e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x18, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x22, 0x63, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x11, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x49, 0x74, 0x65, 0x6d, 0x42, 0x11, 0x0a, 0x0f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x22, 0xc3, 0x01, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x57, 0x69, 0x74, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x1f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x14, 0x0a, 0x12, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbb, 0x01,
	0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74,
	0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68, 0x65,
	0x6e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x18, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x15, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2a, 0x4d, 0x0a, 0x10, 0x45,
	0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x0a, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x02, 0x12, 0x1b, 0x0a,
	0x17, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x04, 0x32, 0xe7, 0x01, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x13, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x90, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x33, 0x2f, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_fixtures_core_proto_rawDescOnce sync.Once
	file_fixtures_core_proto_rawDescData = file_fixtures_core_proto_rawDesc
)

func file_fixtures_core_proto_rawDescGZIP() []byte {
	file_fixtures_core_proto_rawDescOnce.Do(func() {
		file_fixtures_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_fixtures_core_proto_rawDescData)
	})
	return file_fixtures_core_proto_rawDescData
}

var file_fixtures_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fixtures_core_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_fixtures_core_proto_goTypes = []interface{}{
	(EnumNotAnnotated)(0),                     // 0: com.github.devnev.proto_releases.fixtures.EnumNotAnnotated
	(*EmptyMessageReleased)(nil),              // 1: com.github.devnev.proto_releases.fixtures.EmptyMessageReleased
	(*EmptyMessagePreviewedThenReleased)(nil), // 2: com.github.devnev.proto_releases.fixtures.EmptyMessagePreviewedThenReleased
	(*MessageWithReleasedField)(nil),          // 3: com.github.devnev.proto_releases.fixtures.MessageWithReleasedField
	(*MessageWithReleasedOneofItem)(nil),      // 4: com.github.devnev.proto_releases.fixtures.MessageWithReleasedOneofItem
	(*MessageWithOneofWithMessages)(nil),      // 5: com.github.devnev.proto_releases.fixtures.MessageWithOneofWithMessages
	(*MessageNotAnnotated)(nil),               // 6: com.github.devnev.proto_releases.fixtures.MessageNotAnnotated
	(*emptypb.Empty)(nil),                     // 7: google.protobuf.Empty
}
var file_fixtures_core_proto_depIdxs = []int32{
	3, // 0: com.github.devnev.proto_releases.fixtures.MessageWithOneofWithMessages.message_with_release_annotation:type_name -> com.github.devnev.proto_releases.fixtures.MessageWithReleasedField
	7, // 1: com.github.devnev.proto_releases.fixtures.TestService.EmptyMethodReleased:input_type -> google.protobuf.Empty
	6, // 2: com.github.devnev.proto_releases.fixtures.TestService.MethodReleased:input_type -> com.github.devnev.proto_releases.fixtures.MessageNotAnnotated
	7, // 3: com.github.devnev.proto_releases.fixtures.TestService.EmptyMethodReleased:output_type -> google.protobuf.Empty
	6, // 4: com.github.devnev.proto_releases.fixtures.TestService.MethodReleased:output_type -> com.github.devnev.proto_releases.fixtures.MessageNotAnnotated
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fixtures_core_proto_init() }
func file_fixtures_core_proto_init() {
	if File_fixtures_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fixtures_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessageReleased); i {
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
		file_fixtures_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessagePreviewedThenReleased); i {
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
		file_fixtures_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithReleasedField); i {
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
		file_fixtures_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithReleasedOneofItem); i {
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
		file_fixtures_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOneofWithMessages); i {
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
		file_fixtures_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageNotAnnotated); i {
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
	file_fixtures_core_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MessageWithReleasedOneofItem_ReleasedOneofItem)(nil),
	}
	file_fixtures_core_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*MessageWithOneofWithMessages_MessageWithReleaseAnnotation)(nil),
	}
	file_fixtures_core_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*MessageNotAnnotated_OneofItemNotAnnotated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fixtures_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fixtures_core_proto_goTypes,
		DependencyIndexes: file_fixtures_core_proto_depIdxs,
		EnumInfos:         file_fixtures_core_proto_enumTypes,
		MessageInfos:      file_fixtures_core_proto_msgTypes,
	}.Build()
	File_fixtures_core_proto = out.File
	file_fixtures_core_proto_rawDesc = nil
	file_fixtures_core_proto_goTypes = nil
	file_fixtures_core_proto_depIdxs = nil
}
