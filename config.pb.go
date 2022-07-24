// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: config.proto

package releases

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of the release.
	//
	// If non-zero, this determines the release number to use when applying the
	// options that specify which fields to include in a release. See the Range
	// message for further details.
	Release uint64 `protobuf:"varint,1,opt,name=release,proto3" json:"release,omitempty"`
	// Whether the release is a preview release. See the Range message for
	// further details.
	Preview   bool                     `protobuf:"varint,2,opt,name=preview,proto3" json:"preview,omitempty"`
	Package   *Config_PackageMapping   `protobuf:"bytes,4,opt,name=package,proto3" json:"package,omitempty"`
	GoPackage *Config_GoPackageMapping `protobuf:"bytes,3,opt,name=go_package,json=goPackage,proto3" json:"go_package,omitempty"`
	HttpRule  *Config_HttpRuleMapping  `protobuf:"bytes,5,opt,name=http_rule,json=httpRule,proto3" json:"http_rule,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetRelease() uint64 {
	if x != nil {
		return x.Release
	}
	return 0
}

func (x *Config) GetPreview() bool {
	if x != nil {
		return x.Preview
	}
	return false
}

func (x *Config) GetPackage() *Config_PackageMapping {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *Config) GetGoPackage() *Config_GoPackageMapping {
	if x != nil {
		return x.GoPackage
	}
	return nil
}

func (x *Config) GetHttpRule() *Config_HttpRuleMapping {
	if x != nil {
		return x.HttpRule
	}
	return nil
}

// PackageMapping specifies how package paths should be transformed.
//
// The output package is concatenated from three parts:
// <release_root>.<source_suffix>.<release_suffix>
//
// <release_root> and <release_suffix> are specified by the fields of
// PackageMapping. <source_suffix> is taken from the input by stripping the
// prefix specified by the <source_root> field of PackageMapping.
type Config_PackageMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceRoot    string `protobuf:"bytes,1,opt,name=source_root,json=sourceRoot,proto3" json:"source_root,omitempty"`
	ReleaseRoot   string `protobuf:"bytes,2,opt,name=release_root,json=releaseRoot,proto3" json:"release_root,omitempty"`
	ReleaseSuffix string `protobuf:"bytes,3,opt,name=release_suffix,json=releaseSuffix,proto3" json:"release_suffix,omitempty"`
}

func (x *Config_PackageMapping) Reset() {
	*x = Config_PackageMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_PackageMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_PackageMapping) ProtoMessage() {}

func (x *Config_PackageMapping) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_PackageMapping.ProtoReflect.Descriptor instead.
func (*Config_PackageMapping) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Config_PackageMapping) GetSourceRoot() string {
	if x != nil {
		return x.SourceRoot
	}
	return ""
}

func (x *Config_PackageMapping) GetReleaseRoot() string {
	if x != nil {
		return x.ReleaseRoot
	}
	return ""
}

func (x *Config_PackageMapping) GetReleaseSuffix() string {
	if x != nil {
		return x.ReleaseSuffix
	}
	return ""
}

// GoPackageMapping specifies how package paths should be transformed.
//
// The output package is concatenated from three parts:
// <release_root>/<source_suffix>/<release_suffix>
//
// <release_root> and <release_suffix> are specified by the fields of
// GoPackageMapping. <source_suffix> is taken from the input by stripping
// the prefix specified by the <source_root> field of GoPackageMapping.
type Config_GoPackageMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceRoot    string `protobuf:"bytes,1,opt,name=source_root,json=sourceRoot,proto3" json:"source_root,omitempty"`
	ReleaseRoot   string `protobuf:"bytes,2,opt,name=release_root,json=releaseRoot,proto3" json:"release_root,omitempty"`
	ReleaseSuffix string `protobuf:"bytes,3,opt,name=release_suffix,json=releaseSuffix,proto3" json:"release_suffix,omitempty"`
}

func (x *Config_GoPackageMapping) Reset() {
	*x = Config_GoPackageMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_GoPackageMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_GoPackageMapping) ProtoMessage() {}

func (x *Config_GoPackageMapping) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_GoPackageMapping.ProtoReflect.Descriptor instead.
func (*Config_GoPackageMapping) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Config_GoPackageMapping) GetSourceRoot() string {
	if x != nil {
		return x.SourceRoot
	}
	return ""
}

func (x *Config_GoPackageMapping) GetReleaseRoot() string {
	if x != nil {
		return x.ReleaseRoot
	}
	return ""
}

func (x *Config_GoPackageMapping) GetReleaseSuffix() string {
	if x != nil {
		return x.ReleaseSuffix
	}
	return ""
}

type Config_HttpRuleMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceRoot  string `protobuf:"bytes,1,opt,name=source_root,json=sourceRoot,proto3" json:"source_root,omitempty"`
	ReleaseRoot string `protobuf:"bytes,2,opt,name=release_root,json=releaseRoot,proto3" json:"release_root,omitempty"`
}

func (x *Config_HttpRuleMapping) Reset() {
	*x = Config_HttpRuleMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_HttpRuleMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_HttpRuleMapping) ProtoMessage() {}

func (x *Config_HttpRuleMapping) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_HttpRuleMapping.ProtoReflect.Descriptor instead.
func (*Config_HttpRuleMapping) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Config_HttpRuleMapping) GetSourceRoot() string {
	if x != nil {
		return x.SourceRoot
	}
	return ""
}

func (x *Config_HttpRuleMapping) GetReleaseRoot() string {
	if x != nil {
		return x.ReleaseRoot
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x22, 0xcb, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x67, 0x6f, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75,
	0x6c, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x75, 0x6c, 0x65, 0x1a, 0x7b, 0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x1a, 0x7d, 0x0a, 0x10, 0x47, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x1a,
	0x55, 0x0a, 0x0f, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_proto_goTypes = []interface{}{
	(*Config)(nil),                  // 0: releases.Config
	(*Config_PackageMapping)(nil),   // 1: releases.Config.PackageMapping
	(*Config_GoPackageMapping)(nil), // 2: releases.Config.GoPackageMapping
	(*Config_HttpRuleMapping)(nil),  // 3: releases.Config.HttpRuleMapping
}
var file_config_proto_depIdxs = []int32{
	1, // 0: releases.Config.package:type_name -> releases.Config.PackageMapping
	2, // 1: releases.Config.go_package:type_name -> releases.Config.GoPackageMapping
	3, // 2: releases.Config.http_rule:type_name -> releases.Config.HttpRuleMapping
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_PackageMapping); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_GoPackageMapping); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_HttpRuleMapping); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
