// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: example.proto

package alpha

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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Not annotated means this is not include in any numbered release.
	//
	// Intended for internal or alpha-only fields.
	NotAnnotated int32 `protobuf:"varint,1,opt,name=not_annotated,json=notAnnotated,proto3" json:"not_annotated,omitempty"`
	// With a designated starting release, this field will be included in any
	// release from this number onwards.
	//
	// Intended for stabilising fields to be included in general availability
	// releases.
	Released int32 `protobuf:"varint,2,opt,name=released,proto3" json:"released,omitempty"`
	// With a designated preview release, this field will be included in any
	// preview releases from this number onwards.
	//
	// Intended to allow fields under development to be included in beta
	// releases.
	Previewed                int32 `protobuf:"varint,3,opt,name=previewed,proto3" json:"previewed,omitempty"`
	PreviewedReleasedRemoved int32 `protobuf:"varint,4,opt,name=previewed_released_removed,json=previewedReleasedRemoved,proto3" json:"previewed_released_removed,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetNotAnnotated() int32 {
	if x != nil {
		return x.NotAnnotated
	}
	return 0
}

func (x *Example) GetReleased() int32 {
	if x != nil {
		return x.Released
	}
	return 0
}

func (x *Example) GetPreviewed() int32 {
	if x != nil {
		return x.Previewed
	}
	return 0
}

func (x *Example) GetPreviewedReleasedRemoved() int32 {
	if x != nil {
		return x.PreviewedReleasedRemoved
	}
	return 0
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa6, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6e,
	0x6f, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x1a, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x64, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x32, 0x6b, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x16, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x55, 0x6e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x08, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x08,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x15, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x08, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x08, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_proto_goTypes = []interface{}{
	(*Example)(nil), // 0: Example
}
var file_example_proto_depIdxs = []int32{
	0, // 0: ExampleService.ExampleUnreleaseMethod:input_type -> Example
	0, // 1: ExampleService.ExampleReleasedMethod:input_type -> Example
	0, // 2: ExampleService.ExampleUnreleaseMethod:output_type -> Example
	0, // 3: ExampleService.ExampleReleasedMethod:output_type -> Example
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
