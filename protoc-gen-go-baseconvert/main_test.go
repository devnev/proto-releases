package main

import (
	"errors"
	"io/fs"
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
		testdataDir = filepath.Join("..", "testdata")
		goldenDir   = filepath.Join(testdataDir, "golden")
		outDir      = t.TempDir()
	)

	cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, "protoc-gen-go-baseconvert"), ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("unable to build plugin executable: %v", err)
	}

	cmd = exec.Command(
		"protoc",
		"-I"+testdataDir,
		"-I..",
		"--go_out="+outDir,
		"--go_opt=paths=source_relative",
		"release-option-combinations.proto",
	)
	cmd.Env = append(os.Environ(), "PATH="+binDir+":"+os.Getenv("PATH"))
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("unable to generate base proto")
	}

	if err := filepath.WalkDir(goldenDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".proto" {
			return err
		}
		relPath, err := filepath.Rel(goldenDir, path)
		if err != nil {
			return err
		}
		cmd = exec.Command(
			"protoc",
			"-I"+goldenDir,
			"-I..",
			"--go_out="+outDir,
			"--go_opt=paths=source_relative",
			"--go-baseconvert_out="+outDir,
			"--go-baseconvert_opt=base=github.com/devnev/proto-releases/testdata/golden,paths=source_relative",
			relPath,
		)
		cmd.Env = append(os.Environ(), "PATH="+binDir+":"+os.Getenv("PATH"))
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}); err != nil {
		t.Errorf("failed to find files: %v", err)
	}

	cmd = exec.Command("diff", "--unified", "--recursive", "--exclude=*.proto", outDir, goldenDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("error from %q: %v", cmd.Args, err)
	}
}
