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

	binDir := t.TempDir()
	outDir := t.TempDir()

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
				"-I"+filepath.Join("..", "testdata"),
				"-I..",
				"--release_out="+outPath,
				fmt.Sprintf("--release_opt=release=%v,preview=%v", release, preview),
				"release-option-combinations.proto",
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
	cmd = exec.Command("diff", "-ur", outDir, filepath.Join("..", "testdata", "golden"))
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
