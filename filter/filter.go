package filter

import (
	"fmt"
	"path"
	"reflect"
	"strings"
	"unicode/utf8"

	releases "github.com/devnev/proto-releases"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc/builder"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func File(b *builder.FileBuilder, c *releases.Config) error {
	if relgopkg, srcgopkg := c.GetGoPackage(), b.Options.GetGoPackage(); relgopkg != "" && srcgopkg != "" {
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
	return shouldKeep(b, c, releases.E_Method)
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

func GoPackage(srcgopkg, relgopkg string) string {
	if relgopkg == "" {
		return srcgopkg
	}
	if srcgopkg == "" {
		return relgopkg
	}
	if i := strings.LastIndex(srcgopkg, ";"); i > 0 {
		srcgopkg = srcgopkg[:i]
	}
	var relprefix, relsuffix string
	if i := strings.Index(relgopkg, ":"); i > 0 {
		relprefix = relgopkg[:i]
		relsuffix = relgopkg[i+1:]
	} else {
		relprefix = commonPrefix(relgopkg, srcgopkg, '/')
		relsuffix = relgopkg[len(relprefix):]
	}
	return path.Join(relprefix, relsuffix, strings.TrimPrefix(srcgopkg, relprefix))
}

func commonPrefix(path1, path2 string, sep rune) string {
	afterPrefix := 0
	byteOff := 0
	for {
		r1, sz1 := utf8.DecodeRuneInString(path1[byteOff:])
		r2, sz2 := utf8.DecodeRuneInString(path2[byteOff:])
		if (r1 == utf8.RuneError && sz1 > 0) || (r2 == utf8.RuneError && sz2 > 0) {
			// on error, return what we have
			return path1[:afterPrefix]
		}
		endOf1, endOf2 := r1 == utf8.RuneError, r2 == utf8.RuneError
		if endOf1 && endOf2 {
			// if both ended, return entire path
			return path1
		}
		if sep != 0 && endOf1 && r2 == sep {
			// path1 ended at separator in path2
			return path1
		}
		if sep != 0 && endOf2 && r1 == sep {
			// path2 ended at separator in path1
			return path2
		}
		if r1 != r2 || sz1 != sz2 {
			// paths diverged, return found prefix
			return path1[:afterPrefix]
		}
		byteOff += sz1
		if sep == 0 || r1 == sep {
			afterPrefix = byteOff
		}
	}
}
