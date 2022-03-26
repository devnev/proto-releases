package main

import (
	"flag"
	"fmt"

	releases "github.com/devnev/proto-releases"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{
		ParamFunc: flag.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_torelease.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-torelease. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, m := range file.Messages {
		generateMessage(g, m)
	}

	for _, e := range file.Enums {
		generateEnum(g, e)
	}

	return g
}

func generateMessage(g *protogen.GeneratedFile, m *protogen.Message) {
	configIdent := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/devnev/proto-releases",
		GoName:       "Config",
	})

	g.P("func (m *", m.GoIdent, ") ToRelease(c *", configIdent, ") {")
	if len(m.Fields) == 0 {
		g.P("}")
		return
	}
	g.P("if m == nil || c.GetRelease() == 0 { return }")
	if hasReleaseOptions(m.Fields) {
		g.P("r, p := c.GetRelease(), c.GetPreview()")
		g.P("_, _ = r, p // prevent unused variable")
	}
	for _, f := range m.Fields {
		if f.Oneof != nil {
			continue
		}
		cond := buildCondition(f.Desc, releases.E_Field)
		if f.Message != nil {
			if cond != "" {
				g.P("if !(", cond, ") {")
				g.P("m.", f.GoName, " = nil")
				g.P("}")
			}
			g.P("m.", f.GoName, ".ToRelease(c)")
		} else {
			if cond != "" {
				g.P("if !(", cond, ") {")
			}
			g.P("m.", f.GoName, " = ", f.Desc.Default())
			if cond != "" {
				g.P("}")
				if f.Enum != nil {
					g.P("m.", f.GoName, " = m.", f.GoName, ".ToRelease(c)")
				}
			}
		}
	}
	for _, o := range m.Oneofs {
		if hasReleaseOptions(o.Fields) {
			g.P("m.", o.GoName, " = toRelease_", o.GoIdent.GoName, "(m.", o.GoName, ", c)")
		} else {
			g.P("m.", o.GoName, " = nil")
		}
	}
	g.P("}")
	for _, o := range m.Oneofs {
		if hasReleaseOptions(o.Fields) {
			generateOneOfHelper(g, o)
		}
	}
}

func generateOneOfHelper(g *protogen.GeneratedFile, o *protogen.Oneof) {
	configIdent := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/devnev/proto-releases",
		GoName:       "Config",
	})

	g.P("func toRelease_", o.GoIdent.GoName, "(o is", o.GoIdent, ", c *", configIdent, ") is", o.GoIdent, " {")
	if !hasReleaseOptions(o.Fields) {
		g.P("return nil")
		g.P("}")
		return
	}

	g.P("r, p := c.GetRelease(), c.GetPreview()")
	g.P("_, _ = r, p // prevent unused variable")
	g.P("switch t := o.(type) {")
	for _, f := range o.Fields {
		cond := buildCondition(f.Desc, releases.E_Field)
		if f.Message != nil {
			g.P("case *", f.GoIdent, ":")
			if cond != "" {
				g.P("if ", cond, " {")
			}
			g.P("t.", f.GoName, ".ToRelease(c)")
			g.P("return o")
			if cond != "" {
				g.P("}")
			}
		} else if cond != "" {
			g.P("case *", f.GoIdent, ":")
			g.P("if ", cond, " {")
			if f.Enum != nil {
				g.P("t.", f.GoName, " = t.", f.GoName, ".ToRelease(c)")
			}
			g.P("return t")
			g.P("}")
		}
	}
	g.P("}")
	g.P("return nil")
	g.P("}")
}

func generateEnum(g *protogen.GeneratedFile, e *protogen.Enum) {
	configIdent := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/devnev/proto-releases",
		GoName:       "Config",
	})

	g.P("func (e ", e.GoIdent, ") ToRelease(c *", configIdent, ") ", e.GoIdent, " {")
	g.P("if c.GetRelease() == 0 { return e }")
	g.P("r, p := c.GetRelease(), c.GetPreview()")
	g.P("_, _ = r, p // prevent unused variable")
	g.P("switch e {")
	for _, v := range e.Values {
		if v.Desc.Number() == 0 {
			continue
		}
		cond := buildCondition(v.Desc, releases.E_Enum)
		if cond == "" {
			continue
		}
		g.P("case ", v.GoIdent, ":")
		g.P("if ", cond, " { return e }")
	}
	g.P("}")
	g.P("return 0")
	g.P("}")
}

func hasReleaseOptions(fs []*protogen.Field) bool {
	for _, f := range fs {
		range_, _ := proto.GetExtension(f.Desc.Options(), releases.E_Field).(*releases.Range)
		if range_.GetReleaseIn() > 0 || range_.GetPreviewIn() > 0 {
			return true
		}
	}
	return false
}

func buildCondition(d protoreflect.Descriptor, xt protoreflect.ExtensionType) string {
	range_, _ := proto.GetExtension(d.Options(), xt).(*releases.Range)
	var cond string
	if range_.GetReleaseIn() > 0 {
		cond = fmt.Sprintf("r >= %d", range_.GetReleaseIn())
	}
	if range_.GetPreviewIn() > 0 {
		if cond != "" {
			cond += " || "
		}
		cond += fmt.Sprintf("(p && r >= %d)", range_.GetPreviewIn())
	}
	if range_.GetRemoveIn() > 0 && cond != "" {
		cond = fmt.Sprintf("r < %d && (%s)", range_.GetRemoveIn(), cond)
	}
	return cond
}
