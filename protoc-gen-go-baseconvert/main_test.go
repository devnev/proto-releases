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
		repoRoot    = ".."
		releasesDir = filepath.Join(repoRoot, "fixtures", "releases")
		outDir      = t.TempDir()
	)

	cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, "protoc-gen-go-baseconvert"), ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("unable to build plugin executable: %v", err)
	}

	for _, preview := range []bool{false, true} {
		for release := 0; release < 4; release++ {
			name := testName(release, preview)
			relPath := filepath.Join(releasesDir, name)
			outPath := filepath.Join(outDir, name)
			if err := os.MkdirAll(outPath, 0o755); err != nil {
				t.Fatalf("unable to create %q: %v", outPath, err)
			}

			cmd = exec.Command(
				"protoc",
				"-I"+relPath,
				"--go_out="+outPath,
				"--go_opt=paths=source_relative",
				"--go-grpc_out="+outPath,
				"--go-grpc_opt=paths=source_relative",
				"--go-baseconvert_out="+outPath,
				"--go-baseconvert_opt=base_go_package=github.com/devnev/proto-releases/fixtures/releases/"+name+":../../..",
				"--go-baseconvert_opt=paths=source_relative",
				"fixtures/core.proto",
				"fixtures/imported.proto",
				"fixtures/subpackage/subimport.proto",
			)
			cmd.Env = append(os.Environ(), "PATH="+binDir+":"+os.Getenv("PATH"))
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("failed to run protoc: %v", err)
			}
		}
	}

	cmd = exec.Command("diff", "--unified", "--recursive", "--exclude=*.proto", outDir, releasesDir)
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
