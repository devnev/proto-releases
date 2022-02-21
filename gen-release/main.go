// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"log"

	releases "github.com/devnev/proto-releases"
	"github.com/devnev/proto-releases/filter"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/protoprint"
)

var (
	out     = flag.String("out", "out", "")
	release = flag.Int("rel", 0, "")
	preview = flag.Bool("pre", false, "")
	include = flag.String("inc", ".", "")
)

func main() {
	flag.Parse()
	config := &releases.Config{
		Release: int32(*release),
		Preview: *preview,
	}
	log.Printf("releasing for %q", config)
	parser := protoparse.Parser{
		ImportPaths:           []string{*include, "."},
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
	if err != nil {
		log.Fatal(err)
	}
	var kept []*desc.FileDescriptor
	for _, fdesc := range descs {
		log.Println("processing", fdesc.GetName())
		b, err := builder.FromFile(fdesc)
		if err != nil {
			log.Fatal(err)
		}
		filter.File(b, config)
		fdesc, err = b.Build()
		if err != nil {
			log.Fatal(err)
		}
		kept = append(kept, fdesc)
	}
	var printer protoprint.Printer
	err = printer.PrintProtosToFileSystem(kept, *out)
	if err != nil {
		log.Fatal(err)
	}
}
