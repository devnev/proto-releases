// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"

	releases "github.com/devnev/proto-releases"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/protoprint"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

var (
	out     = flag.String("out", "out", "")
	release = flag.Int("rel", 0, "")
	preview = flag.Bool("pre", false, "")
)

func main() {
	flag.Parse()
	config := &releases.Config{
		Release: int32(*release),
		Preview: *preview,
	}
	log.Printf("releasing for %q", config)
	parser := protoparse.Parser{
		ImportPaths:           []string{".", ".."},
		InferImportPaths:      true,
		IncludeSourceCodeInfo: true,
		ErrorReporter: func(err protoparse.ErrorWithPos) error {
			log.Println("parse error", err)
			return err
		},
		WarningReporter: func(ewp protoparse.ErrorWithPos) {
			log.Println("parse warning", ewp)
		},
	}
	log.Printf("parsing %q", flag.Args())
	descs, err := parser.ParseFiles(flag.Args()...)
	checkErr(err)
	var kept []*desc.FileDescriptor
	for _, fdesc := range descs {
		log.Println("processing", fdesc.GetName())
		b, err := builder.FromFile(fdesc)
		checkErr(err)
		filterFile(b, config)
		fdesc, err = b.Build()
		checkErr(err)
		kept = append(kept, fdesc)
	}
	var printer protoprint.Printer
	err = printer.PrintProtosToFileSystem(kept, *out)
	checkErr(err)
}

func filterFile(b *builder.FileBuilder, c *releases.Config) {
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.MessageBuilder:
			if !filterMessage(cb, c) {
				b.RemoveMessage(child.GetName())
			}
		case *builder.EnumBuilder:
			if !filterEnum(cb, c) {
				b.RemoveEnum(child.GetName())
			}
		default:
			panic(fmt.Sprintf("unexpected message child %T", cb))
		}
	}
}

func filterMessage(b *builder.MessageBuilder, c *releases.Config) bool {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.FieldBuilder:
			if filterField(cb, c) {
				include = true
			} else {
				b.RemoveField(cb.GetName())
			}
		case *builder.OneOfBuilder:
			if filterOneOf(cb, c) {
				include = true
			} else {
				b.RemoveOneOf(cb.GetName())
			}
		case *builder.MessageBuilder:
			if filterMessage(cb, c) {
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

func filterField(b *builder.FieldBuilder, c *releases.Config) bool {
	return shouldKeep(b, c, releases.E_Field)
}

func filterOneOf(b *builder.OneOfBuilder, c *releases.Config) bool {
	var include bool
	for _, child := range b.GetChildren() {
		switch cb := child.(type) {
		case *builder.FieldBuilder:
			if filterField(cb, c) {
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

func filterEnum(b *builder.EnumBuilder, c *releases.Config) bool {
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

func filterEnumValue(b *builder.EnumValueBuilder, c *releases.Config) bool {
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
