package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	outDir := t.TempDir()
	t.Logf("generating output to dir %q", outDir)
	includes := []string{filepath.Join("..", "testdata"), ".."}
	for _, preview := range []bool{false, true} {
		for release := 0; release < 4; release++ {
			outPath := filepath.Join(outDir, testName(release, preview))
			run(outPath, release, preview, includes, []string{"release-option-combinations.proto"})
		}
	}
	cmd := exec.Command("diff", "-ur", outDir, filepath.Join("..", "testdata", "golden"))
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
