// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: fixtures/subpackage/subimport.proto

package subpackage

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

type ImportedNotAnnotatedAndEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportedNotAnnotatedAndEmpty) Reset() {
	*x = ImportedNotAnnotatedAndEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_subpackage_subimport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportedNotAnnotatedAndEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportedNotAnnotatedAndEmpty) ProtoMessage() {}

func (x *ImportedNotAnnotatedAndEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_subpackage_subimport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportedNotAnnotatedAndEmpty.ProtoReflect.Descriptor instead.
func (*ImportedNotAnnotatedAndEmpty) Descriptor() ([]byte, []int) {
	return file_fixtures_subpackage_subimport_proto_rawDescGZIP(), []int{0}
}

type ImportedNotAnnotatedWithUnreleasedField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Released int32 `protobuf:"varint,1,opt,name=released,proto3" json:"released,omitempty"`
}

func (x *ImportedNotAnnotatedWithUnreleasedField) Reset() {
	*x = ImportedNotAnnotatedWithUnreleasedField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_subpackage_subimport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportedNotAnnotatedWithUnreleasedField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportedNotAnnotatedWithUnreleasedField) ProtoMessage() {}

func (x *ImportedNotAnnotatedWithUnreleasedField) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_subpackage_subimport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportedNotAnnotatedWithUnreleasedField.ProtoReflect.Descriptor instead.
func (*ImportedNotAnnotatedWithUnreleasedField) Descriptor() ([]byte, []int) {
	return file_fixtures_subpackage_subimport_proto_rawDescGZIP(), []int{1}
}

func (x *ImportedNotAnnotatedWithUnreleasedField) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

type ImportedNotAnnotatedWithReleasedField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Released int32 `protobuf:"varint,1,opt,name=released,proto3" json:"released,omitempty"`
}

func (x *ImportedNotAnnotatedWithReleasedField) Reset() {
	*x = ImportedNotAnnotatedWithReleasedField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fixtures_subpackage_subimport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportedNotAnnotatedWithReleasedField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportedNotAnnotatedWithReleasedField) ProtoMessage() {}

func (x *ImportedNotAnnotatedWithReleasedField) ProtoReflect() protoreflect.Message {
	mi := &file_fixtures_subpackage_subimport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportedNotAnnotatedWithReleasedField.ProtoReflect.Descriptor instead.
func (*ImportedNotAnnotatedWithReleasedField) Descriptor() ([]byte, []int) {
	return file_fixtures_subpackage_subimport_proto_rawDescGZIP(), []int{2}
}

func (x *ImportedNotAnnotatedWithReleasedField) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

var File_fixtures_subpackage_subimport_proto protoreflect.FileDescriptor

var file_fixtures_subpackage_subimport_proto_rawDesc = []byte{
	0x0a, 0x23, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x73, 0x75, 0x62, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x27, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x22, 0x4b, 0x0a, 0x25, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4e, 0x6f,
	0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xfa,
	0xff, 0x1f, 0x02, 0x08, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x76, 0x6e, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fixtures_subpackage_subimport_proto_rawDescOnce sync.Once
	file_fixtures_subpackage_subimport_proto_rawDescData = file_fixtures_subpackage_subimport_proto_rawDesc
)

func file_fixtures_subpackage_subimport_proto_rawDescGZIP() []byte {
	file_fixtures_subpackage_subimport_proto_rawDescOnce.Do(func() {
		file_fixtures_subpackage_subimport_proto_rawDescData = protoimpl.X.CompressGZIP(file_fixtures_subpackage_subimport_proto_rawDescData)
	})
	return file_fixtures_subpackage_subimport_proto_rawDescData
}

var file_fixtures_subpackage_subimport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fixtures_subpackage_subimport_proto_goTypes = []interface{}{
	(*ImportedNotAnnotatedAndEmpty)(nil),            // 0: com.github.devnev.proto_releases.fixtures.subpackage.ImportedNotAnnotatedAndEmpty
	(*ImportedNotAnnotatedWithUnreleasedField)(nil), // 1: com.github.devnev.proto_releases.fixtures.subpackage.ImportedNotAnnotatedWithUnreleasedField
	(*ImportedNotAnnotatedWithReleasedField)(nil),   // 2: com.github.devnev.proto_releases.fixtures.subpackage.ImportedNotAnnotatedWithReleasedField
}
var file_fixtures_subpackage_subimport_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fixtures_subpackage_subimport_proto_init() }
func file_fixtures_subpackage_subimport_proto_init() {
	if File_fixtures_subpackage_subimport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fixtures_subpackage_subimport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportedNotAnnotatedAndEmpty); i {
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
		file_fixtures_subpackage_subimport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportedNotAnnotatedWithUnreleasedField); i {
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
		file_fixtures_subpackage_subimport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportedNotAnnotatedWithReleasedField); i {
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
			RawDescriptor: file_fixtures_subpackage_subimport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fixtures_subpackage_subimport_proto_goTypes,
		DependencyIndexes: file_fixtures_subpackage_subimport_proto_depIdxs,
		MessageInfos:      file_fixtures_subpackage_subimport_proto_msgTypes,
	}.Build()
	File_fixtures_subpackage_subimport_proto = out.File
	file_fixtures_subpackage_subimport_proto_rawDesc = nil
	file_fixtures_subpackage_subimport_proto_goTypes = nil
	file_fixtures_subpackage_subimport_proto_depIdxs = nil
}