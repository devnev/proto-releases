package filter

import (
	"fmt"
	"reflect"

	releases "github.com/devnev/proto-releases"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc/builder"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func File(b *builder.FileBuilder, c *releases.Config) error {
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
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
	return nil
}

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
			panic(fmt.Sprintf("unexpected message child %T", cb))
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
			} else {
				b.RemoveValue(cb.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
	keep, err := shouldKeep(b, c, releases.E_Enum)
	if err != nil {
		return false, err
	}
	return include || keep, nil
}

func EnumValue(b *builder.EnumValueBuilder, c *releases.Config) (bool, error) {
	return shouldKeep(b, c, releases.E_Value)
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
			return false, fmt.Errorf("failed to get extension %s of options %s: %w", x.Name, opts, err)
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
