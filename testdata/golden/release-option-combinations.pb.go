// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: release-option-combinations.proto

package testdata

import (
	_ "github.com/devnev/proto-releases"
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
	EnumNotAnnotated_not_annotated           EnumNotAnnotated = 1
	EnumNotAnnotated_released                EnumNotAnnotated = 2
	EnumNotAnnotated_previewed               EnumNotAnnotated = 3
	EnumNotAnnotated_previewed_then_released EnumNotAnnotated = 4
	EnumNotAnnotated_previewed_then_removed  EnumNotAnnotated = 6
	EnumNotAnnotated_released_then_removed   EnumNotAnnotated = 7
)

// Enum value maps for EnumNotAnnotated.
var (
	EnumNotAnnotated_name = map[int32]string{
		0: "zero_value",
		1: "not_annotated",
		2: "released",
		3: "previewed",
		4: "previewed_then_released",
		6: "previewed_then_removed",
		7: "released_then_removed",
	}
	EnumNotAnnotated_value = map[string]int32{
		"zero_value":              0,
		"not_annotated":           1,
		"released":                2,
		"previewed":               3,
		"previewed_then_released": 4,
		"previewed_then_removed":  6,
		"released_then_removed":   7,
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
	return file_release_option_combinations_proto_enumTypes[0].Descriptor()
}

func (EnumNotAnnotated) Type() protoreflect.EnumType {
	return &file_release_option_combinations_proto_enumTypes[0]
}

func (x EnumNotAnnotated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumNotAnnotated.Descriptor instead.
func (EnumNotAnnotated) EnumDescriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{0}
}

type EmptyMessageNotAnnotated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessageNotAnnotated) Reset() {
	*x = EmptyMessageNotAnnotated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessageNotAnnotated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessageNotAnnotated) ProtoMessage() {}

func (x *EmptyMessageNotAnnotated) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessageNotAnnotated.ProtoReflect.Descriptor instead.
func (*EmptyMessageNotAnnotated) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{0}
}

type EmptyMessageReleased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessageReleased) Reset() {
	*x = EmptyMessageReleased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessageReleased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessageReleased) ProtoMessage() {}

func (x *EmptyMessageReleased) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[1]
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
	return file_release_option_combinations_proto_rawDescGZIP(), []int{1}
}

type EmptyMessagePreviewed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessagePreviewed) Reset() {
	*x = EmptyMessagePreviewed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessagePreviewed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessagePreviewed) ProtoMessage() {}

func (x *EmptyMessagePreviewed) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessagePreviewed.ProtoReflect.Descriptor instead.
func (*EmptyMessagePreviewed) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{2}
}

type EmptyMessagePreviewedThenReleased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessagePreviewedThenReleased) Reset() {
	*x = EmptyMessagePreviewedThenReleased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessagePreviewedThenReleased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessagePreviewedThenReleased) ProtoMessage() {}

func (x *EmptyMessagePreviewedThenReleased) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[3]
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
	return file_release_option_combinations_proto_rawDescGZIP(), []int{3}
}

type EmptyMessagePreviewedThenRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessagePreviewedThenRemoved) Reset() {
	*x = EmptyMessagePreviewedThenRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessagePreviewedThenRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessagePreviewedThenRemoved) ProtoMessage() {}

func (x *EmptyMessagePreviewedThenRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessagePreviewedThenRemoved.ProtoReflect.Descriptor instead.
func (*EmptyMessagePreviewedThenRemoved) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{4}
}

type EmptyMessageReleasedThenRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessageReleasedThenRemoved) Reset() {
	*x = EmptyMessageReleasedThenRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessageReleasedThenRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessageReleasedThenRemoved) ProtoMessage() {}

func (x *EmptyMessageReleasedThenRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessageReleasedThenRemoved.ProtoReflect.Descriptor instead.
func (*EmptyMessageReleasedThenRemoved) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{5}
}

type MessageWithNoAnnotations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotAnnotated int32 `protobuf:"varint,1,opt,name=not_annotated,json=notAnnotated,proto3" json:"not_annotated,omitempty"`
}

func (x *MessageWithNoAnnotations) Reset() {
	*x = MessageWithNoAnnotations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithNoAnnotations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithNoAnnotations) ProtoMessage() {}

func (x *MessageWithNoAnnotations) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithNoAnnotations.ProtoReflect.Descriptor instead.
func (*MessageWithNoAnnotations) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{6}
}

func (x *MessageWithNoAnnotations) GetNotAnnotated() int32 {
	if x != nil {
		return x.NotAnnotated
	}
	return 0
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
		mi := &file_release_option_combinations_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithReleasedField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithReleasedField) ProtoMessage() {}

func (x *MessageWithReleasedField) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[7]
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
	return file_release_option_combinations_proto_rawDescGZIP(), []int{7}
}

func (x *MessageWithReleasedField) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

type MessageWithPreviewField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Released int32 `protobuf:"varint,1,opt,name=released,proto3" json:"released,omitempty"`
}

func (x *MessageWithPreviewField) Reset() {
	*x = MessageWithPreviewField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithPreviewField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithPreviewField) ProtoMessage() {}

func (x *MessageWithPreviewField) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithPreviewField.ProtoReflect.Descriptor instead.
func (*MessageWithPreviewField) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{8}
}

func (x *MessageWithPreviewField) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

type MessageWithUnannotatedOneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NotAnnotatedOneof:
	//	*MessageWithUnannotatedOneof_OneofItemNotAnnotated
	NotAnnotatedOneof isMessageWithUnannotatedOneof_NotAnnotatedOneof `protobuf_oneof:"not_annotated_oneof"`
}

func (x *MessageWithUnannotatedOneof) Reset() {
	*x = MessageWithUnannotatedOneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithUnannotatedOneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithUnannotatedOneof) ProtoMessage() {}

func (x *MessageWithUnannotatedOneof) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithUnannotatedOneof.ProtoReflect.Descriptor instead.
func (*MessageWithUnannotatedOneof) Descriptor() ([]byte, []int) {
	return file_release_option_combinations_proto_rawDescGZIP(), []int{9}
}

func (m *MessageWithUnannotatedOneof) GetNotAnnotatedOneof() isMessageWithUnannotatedOneof_NotAnnotatedOneof {
	if m != nil {
		return m.NotAnnotatedOneof
	}
	return nil
}

func (x *MessageWithUnannotatedOneof) GetOneofItemNotAnnotated() int32 {
	if x, ok := x.GetNotAnnotatedOneof().(*MessageWithUnannotatedOneof_OneofItemNotAnnotated); ok {
		return x.OneofItemNotAnnotated
	}
	return 0
}

type isMessageWithUnannotatedOneof_NotAnnotatedOneof interface {
	isMessageWithUnannotatedOneof_NotAnnotatedOneof()
}

type MessageWithUnannotatedOneof_OneofItemNotAnnotated struct {
	OneofItemNotAnnotated int32 `protobuf:"varint,2,opt,name=oneof_item_not_annotated,json=oneofItemNotAnnotated,proto3,oneof"`
}

func (*MessageWithUnannotatedOneof_OneofItemNotAnnotated) isMessageWithUnannotatedOneof_NotAnnotatedOneof() {
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
		mi := &file_release_option_combinations_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithReleasedOneofItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithReleasedOneofItem) ProtoMessage() {}

func (x *MessageWithReleasedOneofItem) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[10]
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
	return file_release_option_combinations_proto_rawDescGZIP(), []int{10}
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

type MessageNotAnnotated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotAnnotated          int32 `protobuf:"varint,1,opt,name=not_annotated,json=notAnnotated,proto3" json:"not_annotated,omitempty"`
	Released              int32 `protobuf:"varint,2,opt,name=released,proto3" json:"released,omitempty"`
	Previewed             int32 `protobuf:"varint,3,opt,name=previewed,proto3" json:"previewed,omitempty"`
	PreviewedThenReleased int32 `protobuf:"varint,4,opt,name=previewed_then_released,json=previewedThenReleased,proto3" json:"previewed_then_released,omitempty"`
	PreviewedThenRemoved  int32 `protobuf:"varint,6,opt,name=previewed_then_removed,json=previewedThenRemoved,proto3" json:"previewed_then_removed,omitempty"`
	ReleasedThenRemoved   int32 `protobuf:"varint,7,opt,name=released_then_removed,json=releasedThenRemoved,proto3" json:"released_then_removed,omitempty"`
	// Types that are assignable to NotAnnotatedOneof:
	//	*MessageNotAnnotated_OneofItemNotAnnotated
	NotAnnotatedOneof isMessageNotAnnotated_NotAnnotatedOneof `protobuf_oneof:"not_annotated_oneof"`
}

func (x *MessageNotAnnotated) Reset() {
	*x = MessageNotAnnotated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_release_option_combinations_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageNotAnnotated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageNotAnnotated) ProtoMessage() {}

func (x *MessageNotAnnotated) ProtoReflect() protoreflect.Message {
	mi := &file_release_option_combinations_proto_msgTypes[11]
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
	return file_release_option_combinations_proto_rawDescGZIP(), []int{11}
}

func (x *MessageNotAnnotated) GetNotAnnotated() int32 {
	if x != nil {
		return x.NotAnnotated
	}
	return 0
}

func (x *MessageNotAnnotated) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

func (x *MessageNotAnnotated) GetPreviewed() int32 {
	if x != nil {
		return x.Previewed
	}
	return 0
}

func (x *MessageNotAnnotated) GetPreviewedThenReleased() int32 {
	if x != nil {
		return x.PreviewedThenReleased
	}
	return 0
}

func (x *MessageNotAnnotated) GetPreviewedThenRemoved() int32 {
	if x != nil {
		return x.PreviewedThenRemoved
	}
	return 0
}

func (x *MessageNotAnnotated) GetReleasedThenRemoved() int32 {
	if x != nil {
		return x.ReleasedThenRemoved
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

var File_release_option_combinations_proto protoreflect.FileDescriptor

var file_release_option_combinations_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x18, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x22, 0x1e, 0x0a, 0x14,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x3a, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x22, 0x1f, 0x0a, 0x15,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x3a, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0x22, 0x31, 0x0a,
	0x21, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x3a, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x03,
	0x22, 0x30, 0x0a, 0x20, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x3a, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0xfa, 0xff, 0x1f, 0x02,
	0x18, 0x03, 0x22, 0x2f, 0x0a, 0x1f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x3a, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0xfa, 0xff, 0x1f,
	0x02, 0x18, 0x03, 0x22, 0x3f, 0x0a, 0x18, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x4e, 0x6f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x18, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x22, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x22, 0x3d, 0x0a, 0x17, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x22, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x22, 0x6f, 0x0a, 0x1b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x55, 0x6e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x65,
	0x6f, 0x66, 0x12, 0x39, 0x0a, 0x18, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x15, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x49, 0x74, 0x65,
	0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x15, 0x0a,
	0x13, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x22, 0x6b, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x38, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x48, 0x00, 0x52, 0x11, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x11,
	0x0a, 0x0f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0xaa, 0x03, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x6f, 0x74,
	0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x17, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x10,
	0x02, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x03, 0x52, 0x15, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x42,
	0x0a, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c,
	0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0xfa, 0xff, 0x1f, 0x02, 0x18, 0x03, 0x52, 0x14, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x12, 0x40, 0x0a, 0x15, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x74,
	0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0xfa, 0xff, 0x1f, 0x02, 0x18, 0x03, 0x52,
	0x13, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x18, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x48, 0x00,
	0x52, 0x15, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x5f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2a, 0xe0,
	0x01, 0x0a, 0x10, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x10, 0x02, 0x1a, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x12, 0x15, 0x0a, 0x09,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x10, 0x03, 0x1a, 0x06, 0xfa, 0xff, 0x1f,
	0x02, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x17, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64,
	0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x04,
	0x1a, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x10, 0x02, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x03, 0x12, 0x28,
	0x0a, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x06, 0x1a, 0x0c, 0xfa, 0xff, 0x1f, 0x02,
	0x10, 0x02, 0xfa, 0xff, 0x1f, 0x02, 0x18, 0x03, 0x12, 0x27, 0x0a, 0x15, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x10, 0x07, 0x1a, 0x0c, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0xfa, 0xff, 0x1f, 0x02, 0x18,
	0x03, 0x32, 0xaf, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x17, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x13,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x06, 0xfa, 0xff, 0x1f, 0x02, 0x08, 0x02, 0x12, 0x40, 0x0a, 0x12, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x14, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x14, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x44, 0x0a,
	0x0e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12,
	0x14, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x14, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x22, 0x06, 0xfa, 0xff, 0x1f,
	0x02, 0x08, 0x02, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_release_option_combinations_proto_rawDescOnce sync.Once
	file_release_option_combinations_proto_rawDescData = file_release_option_combinations_proto_rawDesc
)

func file_release_option_combinations_proto_rawDescGZIP() []byte {
	file_release_option_combinations_proto_rawDescOnce.Do(func() {
		file_release_option_combinations_proto_rawDescData = protoimpl.X.CompressGZIP(file_release_option_combinations_proto_rawDescData)
	})
	return file_release_option_combinations_proto_rawDescData
}

var file_release_option_combinations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_release_option_combinations_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_release_option_combinations_proto_goTypes = []interface{}{
	(EnumNotAnnotated)(0),                     // 0: EnumNotAnnotated
	(*EmptyMessageNotAnnotated)(nil),          // 1: EmptyMessageNotAnnotated
	(*EmptyMessageReleased)(nil),              // 2: EmptyMessageReleased
	(*EmptyMessagePreviewed)(nil),             // 3: EmptyMessagePreviewed
	(*EmptyMessagePreviewedThenReleased)(nil), // 4: EmptyMessagePreviewedThenReleased
	(*EmptyMessagePreviewedThenRemoved)(nil),  // 5: EmptyMessagePreviewedThenRemoved
	(*EmptyMessageReleasedThenRemoved)(nil),   // 6: EmptyMessageReleasedThenRemoved
	(*MessageWithNoAnnotations)(nil),          // 7: MessageWithNoAnnotations
	(*MessageWithReleasedField)(nil),          // 8: MessageWithReleasedField
	(*MessageWithPreviewField)(nil),           // 9: MessageWithPreviewField
	(*MessageWithUnannotatedOneof)(nil),       // 10: MessageWithUnannotatedOneof
	(*MessageWithReleasedOneofItem)(nil),      // 11: MessageWithReleasedOneofItem
	(*MessageNotAnnotated)(nil),               // 12: MessageNotAnnotated
	(*emptypb.Empty)(nil),                     // 13: google.protobuf.Empty
}
var file_release_option_combinations_proto_depIdxs = []int32{
	13, // 0: TestService.EmptyMethodNotAnnotated:input_type -> google.protobuf.Empty
	13, // 1: TestService.EmptyMethodReleased:input_type -> google.protobuf.Empty
	12, // 2: TestService.MethodNotAnnotated:input_type -> MessageNotAnnotated
	12, // 3: TestService.MethodReleased:input_type -> MessageNotAnnotated
	13, // 4: TestService.EmptyMethodNotAnnotated:output_type -> google.protobuf.Empty
	13, // 5: TestService.EmptyMethodReleased:output_type -> google.protobuf.Empty
	12, // 6: TestService.MethodNotAnnotated:output_type -> MessageNotAnnotated
	12, // 7: TestService.MethodReleased:output_type -> MessageNotAnnotated
	4,  // [4:8] is the sub-list for method output_type
	0,  // [0:4] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_release_option_combinations_proto_init() }
func file_release_option_combinations_proto_init() {
	if File_release_option_combinations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_release_option_combinations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessageNotAnnotated); i {
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
		file_release_option_combinations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_release_option_combinations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessagePreviewed); i {
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
		file_release_option_combinations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_release_option_combinations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessagePreviewedThenRemoved); i {
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
		file_release_option_combinations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessageReleasedThenRemoved); i {
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
		file_release_option_combinations_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithNoAnnotations); i {
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
		file_release_option_combinations_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_release_option_combinations_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithPreviewField); i {
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
		file_release_option_combinations_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithUnannotatedOneof); i {
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
		file_release_option_combinations_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_release_option_combinations_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
	file_release_option_combinations_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*MessageWithUnannotatedOneof_OneofItemNotAnnotated)(nil),
	}
	file_release_option_combinations_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*MessageWithReleasedOneofItem_ReleasedOneofItem)(nil),
	}
	file_release_option_combinations_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*MessageNotAnnotated_OneofItemNotAnnotated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_release_option_combinations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_release_option_combinations_proto_goTypes,
		DependencyIndexes: file_release_option_combinations_proto_depIdxs,
		EnumInfos:         file_release_option_combinations_proto_enumTypes,
		MessageInfos:      file_release_option_combinations_proto_msgTypes,
	}.Build()
	File_release_option_combinations_proto = out.File
	file_release_option_combinations_proto_rawDesc = nil
	file_release_option_combinations_proto_goTypes = nil
	file_release_option_combinations_proto_depIdxs = nil
}
