// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package transform

import (
	"fmt"
	"path"
	"reflect"
	"strings"

	releases "github.com/devnev/proto-releases"
	"github.com/devnev/proto-releases/internal/protos/google.golang.org/genproto/googleapis/api/annotations"
	"github.com/golang/protobuf/proto" //nolint:staticcheck // Need v1 proto API for use with protoreflect
	"github.com/jhump/protoreflect/desc/builder"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Packages(fdp *descriptorpb.FileDescriptorProto, pm *releases.Config_PackageMapping) {
	if fdp.GetPackage() != "" {
		fdp.Package = proto.String(Package(fdp.GetPackage(), pm))
	}
	for _, mdp := range fdp.GetMessageType() {
		for _, fldp := range mdp.GetField() {
			if fldp.GetTypeName() != "" {
				fldp.TypeName = proto.String(Package(fldp.GetTypeName(), pm))
			}
		}
	}
	for _, sdp := range fdp.GetService() {
		for _, mdp := range sdp.GetMethod() {
			if mdp.GetInputType() != "" {
				mdp.InputType = proto.String(Package(mdp.GetInputType(), pm))
			}
			if mdp.GetOutputType() != "" {
				mdp.OutputType = proto.String(Package(mdp.GetOutputType(), pm))
			}
		}
	}
}

//nolint:cyclop // Not as complex as the metric indicates
func File(b *builder.FileBuilder, c *releases.Config) error {
	if srcgopkg := b.Options.GetGoPackage(); srcgopkg != "" {
		outgopkg := GoPackage(srcgopkg, c.GetGoPackage())
		b.Options.GoPackage = &outgopkg
	}

	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.MessageBuilder:
			if keep, err := Message(cb, c); err != nil {
				return fmt.Errorf("failed to filter message %s of file %s: %w", cb.GetName(), b.GetName(), err)
			} else if !keep {
				b.RemoveMessage(child.GetName())
			}
		case *builder.EnumBuilder:
			if keep, err := Enum(cb, c); err != nil {
				return fmt.Errorf("failed to filter enum %s of file %s: %w", cb.GetName(), b.GetName(), err)
			} else if !keep {
				b.RemoveEnum(child.GetName())
			}
		case *builder.ServiceBuilder:
			if keep, err := Service(cb, c); err != nil {
				return fmt.Errorf("failed to filter service %s of file %s: %w", cb.GetName(), b.GetName(), err)
			} else if !keep {
				b.RemoveService(child.GetName())
			}
		case *builder.FieldBuilder:
			if keep, err := shouldKeep(cb, c, releases.E_Field); err != nil {
				return fmt.Errorf("failed to filter extension %s of file %s: %w", cb.GetName(), b.GetName(), err)
			} else if !keep {
				b.RemoveExtension(child.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected file child %T", cb))
		}
	}

	b.PruneUnusedDependencies()

	return nil
}

//nolint:cyclop // Not as complex as the metric indicates
func Message(b *builder.MessageBuilder, c *releases.Config) (bool, error) {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.FieldBuilder:
			if keep, err := Field(cb, c); err != nil {
				return false, fmt.Errorf("failed to filter field %s of message %s: %w", cb.GetName(), b.GetName(), err)
			} else if keep {
				include = true
			} else {
				b.RemoveField(cb.GetName())
			}
		case *builder.OneOfBuilder:
			if keep, err := OneOf(cb, c); err != nil {
				return false, fmt.Errorf("failed to filter oneof %s of message %s: %w", cb.GetName(), b.GetName(), err)
			} else if keep {
				include = true
			} else {
				b.RemoveOneOf(cb.GetName())
			}
		case *builder.MessageBuilder:
			if keep, err := Message(cb, c); err != nil {
				return false, fmt.Errorf("failed to filter message %s of message %s: %w", cb.GetName(), b.GetName(), err)
			} else if keep {
				include = true
			} else {
				b.RemoveNestedMessage(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
	keep, err := shouldKeep(b, c, releases.E_Message)
	if err != nil {
		return false, err
	}
	return include || keep, nil
}

func Field(b *builder.FieldBuilder, c *releases.Config) (bool, error) {
	return shouldKeep(b, c, releases.E_Field)
}

func OneOf(b *builder.OneOfBuilder, c *releases.Config) (bool, error) {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.FieldBuilder:
			if keep, err := Field(cb, c); err != nil {
				return false, fmt.Errorf("failed to filter field %s of oneof %s: %w", cb.GetName(), b.GetName(), err)
			} else if keep {
				include = true
			} else {
				b.RemoveChoice(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected oneof child %T", cb))
		}
	}
	keep, err := shouldKeep(b, c, releases.E_Oneof)
	if err != nil {
		return false, err
	}
	return include || keep, nil
}

func Enum(b *builder.EnumBuilder, c *releases.Config) (bool, error) {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.EnumValueBuilder:
			if keep, err := EnumValue(cb, c); err != nil {
				return false, fmt.Errorf("failed to filter value %s of enum %s: %w", cb.GetName(), b.GetName(), err)
			} else if keep {
				include = true
			} else if cb.GetNumber() != 0 {
				b.RemoveValue(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected enum child %T", cb))
		}
	}
	return include, nil
}

func EnumValue(b *builder.EnumValueBuilder, c *releases.Config) (bool, error) {
	return shouldKeep(b, c, releases.E_Enum)
}

func Service(b *builder.ServiceBuilder, c *releases.Config) (bool, error) {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.MethodBuilder:
			if keep, err := Method(cb, c); err != nil {
				return false, fmt.Errorf("failed to filter method %s of service %s: %w", cb.GetName(), b.GetName(), err)
			} else if keep {
				include = true
			} else {
				b.RemoveMethod(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected service child %T", cb))
		}
	}
	return include, nil
}

func Method(b *builder.MethodBuilder, c *releases.Config) (bool, error) {
	if keep, err := shouldKeep(b, c, releases.E_Method); err != nil || !keep {
		return false, err
	}

	httpExt, _ := proto.GetExtension(b.Options, annotations.E_Http)
	httpRule, _ := httpExt.(*annotations.HttpRule)
	switch pattern := httpRule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		pattern.Get = HTTPPath(pattern.Get, c.GetHttpRule())
	case *annotations.HttpRule_Post:
		pattern.Post = HTTPPath(pattern.Post, c.GetHttpRule())
	case *annotations.HttpRule_Put:
		pattern.Put = HTTPPath(pattern.Put, c.GetHttpRule())
	case *annotations.HttpRule_Patch:
		pattern.Patch = HTTPPath(pattern.Patch, c.GetHttpRule())
	case *annotations.HttpRule_Delete:
		pattern.Delete = HTTPPath(pattern.Delete, c.GetHttpRule())
	case *annotations.HttpRule_Custom:
		pattern.Custom.Path = HTTPPath(pattern.Custom.Path, c.GetHttpRule())
	}

	return true, nil
}

func HTTPPath(srcPath string, httpRuleConfig *releases.Config_HttpRuleMapping) string {
	if httpRuleConfig.GetSourceRoot() == "" && httpRuleConfig.GetReleaseRoot() == "" {
		return srcPath
	}
	if !strings.HasPrefix(srcPath, httpRuleConfig.SourceRoot) {
		return srcPath
	}
	return path.Join(httpRuleConfig.ReleaseRoot, strings.TrimPrefix(srcPath, httpRuleConfig.SourceRoot))
}

func shouldKeep(b builder.Builder, c *releases.Config, x *protoimpl.ExtensionInfo) (bool, error) {
	desc, err := b.BuildDescriptor()
	if err != nil {
		return false, fmt.Errorf("failed to build descriptor for %s: %w", b.GetName(), err)
	}
	var config *releases.Range
	opts := desc.GetOptions()
	if opts != nil && proto.MessageReflect(opts).IsValid() {
		ext, err := proto.GetExtension(opts, x)
		if err != nil {
			return false, fmt.Errorf("failed to get extension %s of options %s: %w", x.TypeDescriptor().FullName(), opts, err)
		}
		config, _ = ext.(*releases.Range)
	}
	include := releases.Include(config, c)
	msg := reflect.ValueOf(b).Elem().FieldByName("Options").Interface().(protoreflect.ProtoMessage).ProtoReflect()
	if msg.Has(x.TypeDescriptor()) {
		msg.Clear(x.TypeDescriptor())
	}
	return include, nil
}

func GoPackage(srcgopkg string, gopkgconfig *releases.Config_GoPackageMapping) string {
	if gopkgconfig.GetReleaseRoot() == "" && gopkgconfig.GetReleaseSuffix() == "" {
		return srcgopkg
	}

	if srcgopkg == "" {
		return path.Join(gopkgconfig.ReleaseRoot, gopkgconfig.ReleaseSuffix)
	}
	if i := strings.LastIndex(srcgopkg, ";"); i > 0 {
		srcgopkg = srcgopkg[:i]
	}
	return path.Join(gopkgconfig.ReleaseRoot, strings.TrimPrefix(srcgopkg, gopkgconfig.SourceRoot), gopkgconfig.ReleaseSuffix)
}

func Package(srcpkg string, pkgconfig *releases.Config_PackageMapping) string {
	if (pkgconfig.GetReleaseRoot() == "" && pkgconfig.GetReleaseSuffix() == "") || srcpkg == "" {
		return srcpkg
	}

	relpkg := pkgconfig.ReleaseRoot
	if strings.HasPrefix(srcpkg, ".") {
		relpkg = "." + relpkg
	}
	var relsuffix string
	if pkgconfig.ReleaseSuffix != "" {
		relsuffix = "." + pkgconfig.ReleaseSuffix
	}

	srcpkgNodot := strings.TrimPrefix(srcpkg, ".")
	if srcpkgNodot == pkgconfig.SourceRoot {
		return relpkg + relsuffix
	}
	if strings.HasPrefix(strings.TrimPrefix(srcpkgNodot, pkgconfig.SourceRoot), ".") {
		return relpkg + srcpkgNodot[len(pkgconfig.SourceRoot):] + relsuffix
	}
	return srcpkg + relsuffix
}
