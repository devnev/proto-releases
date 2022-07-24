// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestProtoc(t *testing.T) {
	if _, err := exec.LookPath("protoc"); errors.Is(err, exec.ErrNotFound) {
		t.Skip()
	} else if err != nil {
		t.Fatalf("unexpected error checking for `protoc`: %v", err)
	}

	var (
		binDir = t.TempDir()
		outDir = t.TempDir()
	)

	cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, "protoc-gen-go-torelease"), ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("unable to build plugin executable: %v", err)
	}

	cmd = exec.Command(
		"protoc",
		"-I..",
		"--go-torelease_out="+outDir,
		"--go-torelease_opt=paths=source_relative",
		"fixtures/core.proto",
		"fixtures/imported.proto",
		"fixtures/subpackage/subimport.proto",
	)
	cmd.Env = append(os.Environ(), "PATH="+binDir)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to compile protos with %q: %v", cmd.Args, err)
	}
	for _, path := range []string{
		"fixtures/core_torelease.pb.go",
		"fixtures/imported_torelease.pb.go",
		"fixtures/subpackage/subimport_torelease.pb.go",
	} {
		cmd = exec.Command("diff", "--unified", filepath.Join(outDir, path), filepath.Join("..", path))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			t.Fatalf("error from %q: %v", cmd.Args, err)
		}
	}
}
