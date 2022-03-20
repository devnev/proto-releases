package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	releases "github.com/devnev/proto-releases"
)

func TestRun(t *testing.T) {
	outDir := t.TempDir()
	t.Logf("generating output to dir %q", outDir)
	includes := []string{".."}
	files := []string{
		"fixtures/core.proto",
		"fixtures/imported.proto",
		"fixtures/subpackage/subimport.proto",
	}
	for _, preview := range []bool{false, true} {
		for release := 0; release < 4; release++ {
			outPath := filepath.Join(outDir, testName(release, preview))
			config := &releases.Config{
				Release: uint64(release),
				Preview: preview,
				GoPackage: &releases.Config_GoPackage{
					SourceRoot:  "github.com/devnev/proto-releases",
					ReleaseRoot: "github.com/devnev/proto-releases/fixtures/releases/" + testName(release, preview),
				},
			}
			run(outPath, config, includes, files)
		}
	}
	cmd := exec.Command("diff", "--unified", "--recursive", "--exclude=*.go", outDir, filepath.Join("..", "fixtures", "releases"))
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
