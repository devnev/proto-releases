package filter

import (
	"fmt"
	"log"
	"reflect"

	releases "github.com/devnev/proto-releases"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc/builder"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func File(b *builder.FileBuilder, c *releases.Config) {
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.MessageBuilder:
			if !Message(cb, c) {
				b.RemoveMessage(child.GetName())
			}
		case *builder.EnumBuilder:
			if !Enum(cb, c) {
				b.RemoveEnum(child.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
}

func Message(b *builder.MessageBuilder, c *releases.Config) bool {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.FieldBuilder:
			if Field(cb, c) {
				include = true
			} else {
				b.RemoveField(cb.GetName())
			}
		case *builder.OneOfBuilder:
			if OneOf(cb, c) {
				include = true
			} else {
				b.RemoveOneOf(cb.GetName())
			}
		case *builder.MessageBuilder:
			if Message(cb, c) {
				include = true
			} else {
				b.RemoveNestedMessage(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
	extDesc := releases.E_Message.TypeDescriptor()
	reflected := b.Options.ProtoReflect()
	if reflected.Has(extDesc) {
		reflected.Clear(extDesc)
	}
	return include
}

func Field(b *builder.FieldBuilder, c *releases.Config) bool {
	return shouldKeep(b, c, releases.E_Field)
}

func OneOf(b *builder.OneOfBuilder, c *releases.Config) bool {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.FieldBuilder:
			if Field(cb, c) {
				include = true
			} else {
				b.RemoveChoice(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
	return include
}

func Enum(b *builder.EnumBuilder, c *releases.Config) bool {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.EnumValueBuilder:
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
	return include
}

func EnumValue(b *builder.EnumValueBuilder, c *releases.Config) bool {
	include := shouldKeep(b, c, releases.E_Value)
	return include
}

func shouldKeep(b builder.Builder, c *releases.Config, x *protoimpl.ExtensionInfo) bool {
	desc, err := b.BuildDescriptor()
	checkErr(err)
	var config *releases.Range
	opts := desc.GetOptions()
	if opts != nil && proto.MessageReflect(opts).IsValid() {
		ext, err := proto.GetExtension(opts, x)
		checkErr(err)
		config, _ = ext.(*releases.Range)
	}
	include := releases.Include(config, c)
	msg := reflect.ValueOf(b).Elem().FieldByName("Options").Interface().(protoreflect.ProtoMessage).ProtoReflect()
	if msg.Has(x.TypeDescriptor()) {
		msg.Clear(x.TypeDescriptor())
	}
	return include
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
