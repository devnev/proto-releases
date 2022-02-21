package main

import (
	"fmt"
	"log"
	"strconv"

	releases "github.com/devnev/proto-releases"
	"github.com/devnev/proto-releases/filter"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).RegisterModule(&mod{}).Render()
}

type mod struct {
	c *releases.Config
}

func (m *mod) Name() string {
	return "release"
}

// InitContext is called on a Module with a pre-configured BuildContext that
// should be stored and used by the Module.
func (m *mod) InitContext(bctx pgs.BuildContext) {
	m.c = &releases.Config{}
	for k, v := range bctx.Parameters() {
		switch k {
		case "preview":
			p, err := strconv.ParseBool(v)
			checkErr(err)
			m.c.Preview = p
		case "release":
			r, err := strconv.ParseInt(v, 10, 32)
			checkErr(err)
			m.c.Release = int32(r)
		case "":
			if v == "" {
				continue
			}
			fallthrough
		default:
			panic(fmt.Sprintf("unkown option %q=%q", k, v))
		}
	}
}

// Execute is called on the module with the target Files as well as all
// loaded Packages from the gatherer. The module should return a slice of
// Artifacts that it would like to be generated.
func (m *mod) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	var fdps []*descriptorpb.FileDescriptorProto
	for _, f := range targets {
		fdps = append(fdps, f.Descriptor())
	}
	for _, p := range packages {
		for _, f := range p.Files() {
			fdps = append(fdps, f.Descriptor())
		}
	}
	fds, err := desc.CreateFileDescriptors(fdps)
	checkErr(err)
	var outputs []pgs.Artifact
	for _, fdesc := range targets {
		fb, err := builder.FromFile(fds[string(fdesc.Name())])
		checkErr(err)
		filter.File(fb, m.c)
		result, err := fb.Build()
		checkErr(err)
		var printer protoprint.Printer
		contents, err := printer.PrintProtoToString(result)
		checkErr(err)
		outputs = append(outputs, pgs.GeneratorFile{
			Name:     result.GetName(),
			Contents: contents,
		})
	}
	return outputs
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
