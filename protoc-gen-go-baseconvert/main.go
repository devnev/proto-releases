// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"

	releases "github.com/devnev/proto-releases"
	"github.com/devnev/proto-releases/transform"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var pkgConfig releases.Config_GoPackageMapping
	flag.Var(&releases.GoPackageFlagValue{Config: &pkgConfig}, "base_go_package", "")
	protogen.Options{
		ParamFunc: flag.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			goPackage := f.Desc.Options().(*descriptorpb.FileOptions).GetGoPackage()
			baseGoPackage := transform.GoPackage(goPackage, &pkgConfig)
			generateFile(gen, f, protogen.GoImportPath(baseGoPackage))
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File, base protogen.GoImportPath) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_baseconvert.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-baseconvert. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, m := range file.Messages {
		generateMessage(g, m, base)
	}

	for _, e := range file.Enums {
		generateEnum(g, e, base)
	}

	for _, s := range file.Services {
		generateService(g, s, base)
	}

	return g
}

func generateMessage(g *protogen.GeneratedFile, m *protogen.Message, base protogen.GoImportPath) {
	generateMessageToBase(g, m, base)
	generateMessageFromBase(g, m, base)

	for _, o := range m.Oneofs {
		for _, f := range o.Fields {
			generateOneofFieldToBase(g, f, base)
			generateOneofFieldFromBase(g, f, base)
		}
	}
}

func generateMessageToBase(g *protogen.GeneratedFile, m *protogen.Message, base protogen.GoImportPath) {
	baseMsg := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       m.GoIdent.GoName,
		GoImportPath: base,
	})

	g.P("func (m *", m.GoIdent.GoName, ") ToBase() *", baseMsg, " {")
	g.P("msg := &", baseMsg, "{")
	for _, f := range m.Fields {
		if f.Oneof != nil {
			continue
		}
		if f.Message != nil || f.Enum != nil {
			g.P(f.GoName, ": m.Get", f.GoName, "().ToBase(),")
		} else {
			g.P(f.GoName, ": m.Get", f.GoName, "(),")
		}
	}
	g.P("}")
	for _, o := range m.Oneofs {
		g.P("switch o := m.Get", o.GoName, "().(type) {")
		for _, f := range o.Fields {
			g.P("case *", f.GoIdent, ":")
			g.P("msg.", o.GoName, " = o.ToBase()")
		}
		g.P("}")
	}
	g.P("return msg")
	g.P("}")
}

//nolint:cyclop // Unclear how to reduce complexity
func generateMessageFromBase(g *protogen.GeneratedFile, m *protogen.Message, base protogen.GoImportPath) {
	baseMsg := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       m.GoIdent.GoName,
		GoImportPath: base,
	})

	g.P("func (m *", m.GoIdent, ") FromBase(b *", baseMsg, ") *", m.GoIdent, " {")
	g.P("if m != nil {")
	g.P("m.Reset()")
	g.P("} else {")
	g.P("m = new(", m.GoIdent, ")")
	g.P("}")
	for _, f := range m.Fields {
		if f.Oneof == nil && f.Message == nil && f.Enum == nil {
			g.P("m.", f.GoName, "= b.Get", f.GoName, "()")
		}
	}
	for _, f := range m.Fields {
		if f.Oneof == nil && (f.Message != nil || f.Enum != nil) {
			g.P("m.", f.GoName, "= m.", f.GoName, ".FromBase(b.Get", f.GoName, "())")
		}
	}
	for _, o := range m.Oneofs {
		g.P("switch o := b.Get", o.GoName, "().(type) {")
		for _, f := range o.Fields {
			g.P("case *", g.QualifiedGoIdent(protogen.GoIdent{
				GoName:       f.GoIdent.GoName,
				GoImportPath: base,
			}), ":")
			g.P("m.", o.GoName, " = (*", f.GoIdent, ")(nil).FromBase(o)")
		}
		g.P("}")
	}
	g.P("return m")
	g.P("}")
}

func generateOneofFieldToBase(g *protogen.GeneratedFile, f *protogen.Field, base protogen.GoImportPath) {
	baseType := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       f.GoIdent.GoName,
		GoImportPath: base,
	})
	g.P("func (m *", f.GoIdent, ") ToBase() *", baseType, "{")
	if f.Message == nil && f.Enum == nil {
		g.P("return (*", baseType, ")(m)")
	} else {
		g.P("return &", baseType, "{")
		g.P(f.GoName, ": m.", f.GoName, ".ToBase(),")
		g.P("}")
	}
	g.P("}")
}

func generateOneofFieldFromBase(g *protogen.GeneratedFile, f *protogen.Field, base protogen.GoImportPath) {
	baseType := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       f.GoIdent.GoName,
		GoImportPath: base,
	})
	g.P("func (m *", f.GoIdent, ") FromBase(b *", baseType, ") *", f.GoIdent, " {")
	if f.Message != nil {
		g.P("return &", f.GoIdent, "{")
		g.P(f.GoName, ": (*", f.Message.GoIdent, ")(nil).FromBase(b.", f.GoName, "),")
		g.P("}")
	} else if f.Enum != nil {
		g.P("return &", f.GoIdent, "{")
		g.P(f.GoName, ": (", f.Enum.GoIdent, ")(0).FromBase(b.", f.GoName, "),")
		g.P("}")
	} else {
		g.P("return (*", f.GoIdent, ")(b)")
	}
	g.P("}")
}

func generateEnum(g *protogen.GeneratedFile, e *protogen.Enum, base protogen.GoImportPath) {
	baseType := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       e.GoIdent.GoName,
		GoImportPath: base,
	})
	g.P("func (e ", e.GoIdent, ") FromBase(b ", baseType, ") ", e.GoIdent, " {")
	g.P("switch b {")
	for _, v := range e.Values {
		g.P("case ", g.QualifiedGoIdent(protogen.GoIdent{
			GoName:       v.GoIdent.GoName,
			GoImportPath: base,
		}), ":")
		g.P("return ", v.GoIdent)
	}
	g.P("default:")
	g.P("return ", e.GoIdent, "(0)")
	g.P("}")
	g.P("}")
	g.P("func (e ", e.GoIdent, ") ToBase() ", baseType, " {")
	g.P("return ", baseType, "(e)")
	g.P("}")
}

func generateService(g *protogen.GeneratedFile, s *protogen.Service, base protogen.GoImportPath) {
	baseSrvIface := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       s.GoName + "Server",
		GoImportPath: base,
	})
	srvName := "base" + s.GoName + "Server"

	g.P("type ", srvName, " struct {")
	g.P("Unsafe", s.GoName, "Server")
	g.P("base ", baseSrvIface)
	g.P("}")
	g.P()

	registrarIface := protogen.GoIdent{
		GoName:       "ServiceRegistrar",
		GoImportPath: "google.golang.org/grpc",
	}

	g.P("func Register", s.GoName, "BaseServer(s ", registrarIface, ", base ", baseSrvIface, ") {")
	g.P("Register", s.GoName, "Server(s, ", srvName, "{base: base})")
	g.P("}")
	g.P()

	for _, m := range s.Methods {
		generateMethod(g, m, srvName)
	}
}

func generateMethod(g *protogen.GeneratedFile, m *protogen.Method, srvName string) {
	emptyDesc := (&emptypb.Empty{}).ProtoReflect().Descriptor()
	ctxType := g.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})

	g.P("func (s ", srvName, ") ", m.GoName, "(ctx ", ctxType, ", in *", m.Input.GoIdent, ") (*", m.Output.GoIdent, ", error) {")
	inVar := "in"
	if m.Input.Desc.FullName() != emptyDesc.FullName() {
		g.P("baseIn := in.ToBase()")
		inVar = "baseIn"
	}
	if m.Output.Desc.FullName() == emptyDesc.FullName() {
		g.P("return s.base.", m.GoName, "(ctx, ", inVar, ")")
		g.P("}")
		return
	}
	g.P("baseOut, err := s.base.", m.GoName, "(ctx, ", inVar, ")")
	g.P("if err != nil { return nil, err }")
	g.P("out := new(", m.Input.GoIdent, ")")
	g.P("out.FromBase(baseOut)")
	g.P("return out, nil")
	g.P("}")
}
