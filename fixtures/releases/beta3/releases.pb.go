// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: releases.proto

package beta3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_releases_proto protoreflect.FileDescriptor

var file_releases_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6e, 0x65,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x66, 0x69, 0x78, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_releases_proto_goTypes = []interface{}{}
var file_releases_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_releases_proto_init() }
func file_releases_proto_init() {
	if File_releases_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_releases_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_releases_proto_goTypes,
		DependencyIndexes: file_releases_proto_depIdxs,
	}.Build()
	File_releases_proto = out.File
	file_releases_proto_rawDesc = nil
	file_releases_proto_goTypes = nil
	file_releases_proto_depIdxs = nil
}