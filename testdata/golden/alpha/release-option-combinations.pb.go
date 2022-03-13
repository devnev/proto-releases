// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: alpha/release-option-combinations.proto

package testdata

import (
	_ "github.com/devnev/proto-releases"
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

type RootEnumNotAnnotated int32

const (
	RootEnumNotAnnotated_zero_value              RootEnumNotAnnotated = 0
	RootEnumNotAnnotated_not_annotated           RootEnumNotAnnotated = 1
	RootEnumNotAnnotated_released                RootEnumNotAnnotated = 2
	RootEnumNotAnnotated_released_eventually     RootEnumNotAnnotated = 5
	RootEnumNotAnnotated_previewed               RootEnumNotAnnotated = 3
	RootEnumNotAnnotated_previewed_then_released RootEnumNotAnnotated = 4
	RootEnumNotAnnotated_previewed_then_removed  RootEnumNotAnnotated = 6
	RootEnumNotAnnotated_released_then_removed   RootEnumNotAnnotated = 7
)

// Enum value maps for RootEnumNotAnnotated.
var (
	RootEnumNotAnnotated_name = map[int32]string{
		0: "zero_value",
		1: "not_annotated",
		2: "released",
		5: "released_eventually",
		3: "previewed",
		4: "previewed_then_released",
		6: "previewed_then_removed",
		7: "released_then_removed",
	}
	RootEnumNotAnnotated_value = map[string]int32{
		"zero_value":              0,
		"not_annotated":           1,
		"released":                2,
		"released_eventually":     5,
		"previewed":               3,
		"previewed_then_released": 4,
		"previewed_then_removed":  6,
		"released_then_removed":   7,
	}
)

func (x RootEnumNotAnnotated) Enum() *RootEnumNotAnnotated {
	p := new(RootEnumNotAnnotated)
	*p = x
	return p
}

func (x RootEnumNotAnnotated) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RootEnumNotAnnotated) Descriptor() protoreflect.EnumDescriptor {
	return file_alpha_release_option_combinations_proto_enumTypes[0].Descriptor()
}

func (RootEnumNotAnnotated) Type() protoreflect.EnumType {
	return &file_alpha_release_option_combinations_proto_enumTypes[0]
}

func (x RootEnumNotAnnotated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RootEnumNotAnnotated.Descriptor instead.
func (RootEnumNotAnnotated) EnumDescriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{0}
}

type EmptyRootMessageNotAnnotated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessageNotAnnotated) Reset() {
	*x = EmptyRootMessageNotAnnotated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessageNotAnnotated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessageNotAnnotated) ProtoMessage() {}

func (x *EmptyRootMessageNotAnnotated) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessageNotAnnotated.ProtoReflect.Descriptor instead.
func (*EmptyRootMessageNotAnnotated) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{0}
}

type EmptyRootMessageReleased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessageReleased) Reset() {
	*x = EmptyRootMessageReleased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessageReleased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessageReleased) ProtoMessage() {}

func (x *EmptyRootMessageReleased) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessageReleased.ProtoReflect.Descriptor instead.
func (*EmptyRootMessageReleased) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{1}
}

type EmptyRootMessageReleasedEventually struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessageReleasedEventually) Reset() {
	*x = EmptyRootMessageReleasedEventually{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessageReleasedEventually) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessageReleasedEventually) ProtoMessage() {}

func (x *EmptyRootMessageReleasedEventually) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessageReleasedEventually.ProtoReflect.Descriptor instead.
func (*EmptyRootMessageReleasedEventually) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{2}
}

type EmptyRootMessagePreviewed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessagePreviewed) Reset() {
	*x = EmptyRootMessagePreviewed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessagePreviewed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessagePreviewed) ProtoMessage() {}

func (x *EmptyRootMessagePreviewed) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessagePreviewed.ProtoReflect.Descriptor instead.
func (*EmptyRootMessagePreviewed) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{3}
}

type EmptyRootMessagePreviewedThenReleased struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessagePreviewedThenReleased) Reset() {
	*x = EmptyRootMessagePreviewedThenReleased{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessagePreviewedThenReleased) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessagePreviewedThenReleased) ProtoMessage() {}

func (x *EmptyRootMessagePreviewedThenReleased) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessagePreviewedThenReleased.ProtoReflect.Descriptor instead.
func (*EmptyRootMessagePreviewedThenReleased) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{4}
}

type EmptyRootMessagePreviewedThenRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessagePreviewedThenRemoved) Reset() {
	*x = EmptyRootMessagePreviewedThenRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessagePreviewedThenRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessagePreviewedThenRemoved) ProtoMessage() {}

func (x *EmptyRootMessagePreviewedThenRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessagePreviewedThenRemoved.ProtoReflect.Descriptor instead.
func (*EmptyRootMessagePreviewedThenRemoved) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{5}
}

type EmptyRootMessageReleasedThenRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRootMessageReleasedThenRemoved) Reset() {
	*x = EmptyRootMessageReleasedThenRemoved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRootMessageReleasedThenRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRootMessageReleasedThenRemoved) ProtoMessage() {}

func (x *EmptyRootMessageReleasedThenRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRootMessageReleasedThenRemoved.ProtoReflect.Descriptor instead.
func (*EmptyRootMessageReleasedThenRemoved) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{6}
}

type RootMessageNotAnnotated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotAnnotated          int32 `protobuf:"varint,1,opt,name=not_annotated,json=notAnnotated,proto3" json:"not_annotated,omitempty"`
	Released              int32 `protobuf:"varint,2,opt,name=released,proto3" json:"released,omitempty"`
	ReleasedEventually    int32 `protobuf:"varint,5,opt,name=released_eventually,json=releasedEventually,proto3" json:"released_eventually,omitempty"`
	Previewed             int32 `protobuf:"varint,3,opt,name=previewed,proto3" json:"previewed,omitempty"`
	PreviewedThenReleased int32 `protobuf:"varint,4,opt,name=previewed_then_released,json=previewedThenReleased,proto3" json:"previewed_then_released,omitempty"`
	PreviewedThenRemoved  int32 `protobuf:"varint,6,opt,name=previewed_then_removed,json=previewedThenRemoved,proto3" json:"previewed_then_removed,omitempty"`
	ReleasedThenRemoved   int32 `protobuf:"varint,7,opt,name=released_then_removed,json=releasedThenRemoved,proto3" json:"released_then_removed,omitempty"`
}

func (x *RootMessageNotAnnotated) Reset() {
	*x = RootMessageNotAnnotated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alpha_release_option_combinations_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootMessageNotAnnotated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootMessageNotAnnotated) ProtoMessage() {}

func (x *RootMessageNotAnnotated) ProtoReflect() protoreflect.Message {
	mi := &file_alpha_release_option_combinations_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootMessageNotAnnotated.ProtoReflect.Descriptor instead.
func (*RootMessageNotAnnotated) Descriptor() ([]byte, []int) {
	return file_alpha_release_option_combinations_proto_rawDescGZIP(), []int{7}
}

func (x *RootMessageNotAnnotated) GetNotAnnotated() int32 {
	if x != nil {
		return x.NotAnnotated
	}
	return 0
}

func (x *RootMessageNotAnnotated) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

func (x *RootMessageNotAnnotated) GetReleasedEventually() int32 {
	if x != nil {
		return x.ReleasedEventually
	}
	return 0
}

func (x *RootMessageNotAnnotated) GetPreviewed() int32 {
	if x != nil {
		return x.Previewed
	}
	return 0
}

func (x *RootMessageNotAnnotated) GetPreviewedThenReleased() int32 {
	if x != nil {
		return x.PreviewedThenReleased
	}
	return 0
}

func (x *RootMessageNotAnnotated) GetPreviewedThenRemoved() int32 {
	if x != nil {
		return x.PreviewedThenRemoved
	}
	return 0
}

func (x *RootMessageNotAnnotated) GetReleasedThenRemoved() int32 {
	if x != nil {
		return x.ReleasedThenRemoved
	}
	return 0
}

var File_alpha_release_option_combinations_proto protoreflect.FileDescriptor

var file_alpha_release_option_combinations_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2d,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x22, 0x24, 0x0a, 0x22, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x6f,
	0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x22, 0x1b, 0x0a, 0x19, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x22, 0x27, 0x0a, 0x25, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x22, 0x26, 0x0a, 0x24, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68,
	0x65, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x25, 0x0a, 0x23, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x22, 0xcb, 0x02, 0x0a, 0x17, 0x52, 0x6f, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x2f, 0x0a,
	0x13, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x61, 0x6c, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x17,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65,
	0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54,
	0x68, 0x65, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x2a, 0xc3,
	0x01, 0x0a, 0x14, 0x52, 0x6f, 0x6f, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x6f, 0x74, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x7a, 0x65, 0x72, 0x6f, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x5f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x10,
	0x05, 0x12, 0x0d, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x10, 0x03,
	0x12, 0x1b, 0x0a, 0x17, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68,
	0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x10, 0x04, 0x12, 0x1a, 0x0a,
	0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x10, 0x07, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alpha_release_option_combinations_proto_rawDescOnce sync.Once
	file_alpha_release_option_combinations_proto_rawDescData = file_alpha_release_option_combinations_proto_rawDesc
)

func file_alpha_release_option_combinations_proto_rawDescGZIP() []byte {
	file_alpha_release_option_combinations_proto_rawDescOnce.Do(func() {
		file_alpha_release_option_combinations_proto_rawDescData = protoimpl.X.CompressGZIP(file_alpha_release_option_combinations_proto_rawDescData)
	})
	return file_alpha_release_option_combinations_proto_rawDescData
}

var file_alpha_release_option_combinations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_alpha_release_option_combinations_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_alpha_release_option_combinations_proto_goTypes = []interface{}{
	(RootEnumNotAnnotated)(0),                     // 0: RootEnumNotAnnotated
	(*EmptyRootMessageNotAnnotated)(nil),          // 1: EmptyRootMessageNotAnnotated
	(*EmptyRootMessageReleased)(nil),              // 2: EmptyRootMessageReleased
	(*EmptyRootMessageReleasedEventually)(nil),    // 3: EmptyRootMessageReleasedEventually
	(*EmptyRootMessagePreviewed)(nil),             // 4: EmptyRootMessagePreviewed
	(*EmptyRootMessagePreviewedThenReleased)(nil), // 5: EmptyRootMessagePreviewedThenReleased
	(*EmptyRootMessagePreviewedThenRemoved)(nil),  // 6: EmptyRootMessagePreviewedThenRemoved
	(*EmptyRootMessageReleasedThenRemoved)(nil),   // 7: EmptyRootMessageReleasedThenRemoved
	(*RootMessageNotAnnotated)(nil),               // 8: RootMessageNotAnnotated
}
var file_alpha_release_option_combinations_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_alpha_release_option_combinations_proto_init() }
func file_alpha_release_option_combinations_proto_init() {
	if File_alpha_release_option_combinations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alpha_release_option_combinations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessageNotAnnotated); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessageReleased); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessageReleasedEventually); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessagePreviewed); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessagePreviewedThenReleased); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessagePreviewedThenRemoved); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRootMessageReleasedThenRemoved); i {
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
		file_alpha_release_option_combinations_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootMessageNotAnnotated); i {
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
			RawDescriptor: file_alpha_release_option_combinations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alpha_release_option_combinations_proto_goTypes,
		DependencyIndexes: file_alpha_release_option_combinations_proto_depIdxs,
		EnumInfos:         file_alpha_release_option_combinations_proto_enumTypes,
		MessageInfos:      file_alpha_release_option_combinations_proto_msgTypes,
	}.Build()
	File_alpha_release_option_combinations_proto = out.File
	file_alpha_release_option_combinations_proto_rawDesc = nil
	file_alpha_release_option_combinations_proto_goTypes = nil
	file_alpha_release_option_combinations_proto_depIdxs = nil
}
