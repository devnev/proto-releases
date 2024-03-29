// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"fmt"
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
		binDir      = t.TempDir()
		fixturesDir = filepath.Join("..", "fixtures")
		releasesDir = filepath.Join(fixturesDir, "releases")
		outDir      = t.TempDir()
	)

	cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, "protoc-gen-release"), ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("unable to build plugin executable: %v", err)
	}

	for _, preview := range []bool{false, true} {
		for release := 0; release < 4; release++ {
			outPath := filepath.Join(outDir, testName(release, preview))
			if err := os.MkdirAll(outPath, 0o755); err != nil {
				t.Errorf("failed to create output directory: %v", err)
			}
			cmd := exec.Command(
				"protoc",
				"-I..",
				"--release_out="+outPath,
				fmt.Sprintf("--release_opt=release=%v", release),
				fmt.Sprintf("--release_opt=preview=%v", preview),
				"--release_opt=go_package=github.com/devnev/proto-releases:./fixtures/releases/"+testName(release, preview),
				"--release_opt=package=com.github.devnev.proto_releases:.fixtures.releases."+testName(release, preview),
				"fixtures/core.proto",
				"fixtures/imported.proto",
				"fixtures/subpackage/subimport.proto",
			)
			cmd.Env = append(os.Environ(), "PATH="+binDir)
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				t.Errorf("failed to compile protos with %q: %v", cmd.Args, err)
			}
		}
	}
	if t.Failed() {
		return
	}
	cmd = exec.Command("diff", "--unified", "--recursive", "--exclude=*.go", outDir, releasesDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("error from %q: %v", cmd.Args, err)
	}
}

func testName(release int, preview bool) string {
	if release == 0 {
		return "alpha"
	}
	prefix := "stable"
	if preview {
		prefix = "beta"
	}
	return fmt.Sprintf("%v%v", prefix, release)
}
