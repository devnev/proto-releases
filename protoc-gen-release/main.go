// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"strconv"
	"strings"

	releases "github.com/devnev/proto-releases"
	"github.com/devnev/proto-releases/transform"
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
//nolint:cyclop // Not as complex as the metric indicates
func (m *mod) InitContext(bctx pgs.BuildContext) {
	m.c = &releases.Config{}
	for k, v := range bctx.Parameters() {
		switch k {
		case "preview":
			p, err := strconv.ParseBool(v)
			if err != nil {
				bctx.Fail(err)
			}
			m.c.Preview = p
		case "release":
			r, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				bctx.Fail(err)
			}
			m.c.Release = r
		case "go_package":
			m.c.GoPackage = &releases.Config_GoPackageMapping{}
			if err := releases.ParseGoPackageShorthand(v, m.c.GoPackage); err != nil {
				bctx.Fail(err)
			}
		case "package":
			m.c.Package = &releases.Config_PackageMapping{}
			if err := releases.ParsePackageShorthand(v, m.c.Package); err != nil {
				bctx.Fail(err)
			}
		case "http_rule":
			m.c.HttpRule = &releases.Config_HttpRuleMapping{}
			m.c.HttpRule.ReleaseRoot = v
			if part1, part2, found := strings.Cut(v, ":"); found {
				m.c.HttpRule.SourceRoot = part1
				m.c.HttpRule.ReleaseRoot = part2
			}
		case "":
			if v == "" {
				continue
			}
			fallthrough
		default:
			panic(fmt.Sprintf("unknown option %q=%q", k, v))
		}
	}
}

// Execute is called on the module with the target Files as well as all
// loaded Packages from the gatherer. The module should return a slice of
// Artifacts that it would like to be generated.
func (m *mod) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	var fdps []*descriptorpb.FileDescriptorProto
	for _, p := range packages {
		for _, f := range p.Files() {
			fdps = append(fdps, f.Descriptor())
		}
	}
	for _, fdp := range fdps {
		transform.Packages(fdp, m.c.GetPackage())
	}

	fds, err := desc.CreateFileDescriptors(fdps)
	if err != nil {
		return []pgs.Artifact{pgs.GeneratorError{
			Message: err.Error(),
		}}
	}
	outputs := make([]pgs.Artifact, 0, len(targets))
	for _, fdesc := range targets {
		fb, err := builder.FromFile(fds[string(fdesc.Name())])
		if err != nil {
			outputs = append(outputs, pgs.GeneratorError{
				Message: err.Error(),
			})
			continue
		}
		if err = transform.File(fb, m.c); err != nil {
			outputs = append(outputs, pgs.GeneratorError{
				Message: err.Error(),
			})
			continue
		}
		result, err := fb.Build()
		if err != nil {
			outputs = append(outputs, pgs.GeneratorError{
				Message: err.Error(),
			})
			continue
		}
		var printer protoprint.Printer
		contents, err := printer.PrintProtoToString(result)
		if err != nil {
			outputs = append(outputs, pgs.GeneratorError{
				Message: err.Error(),
			})
			continue
		}
		outputs = append(outputs, pgs.GeneratorFile{
			Name:     result.GetName(),
			Contents: contents,
		})
	}
	return outputs
}
