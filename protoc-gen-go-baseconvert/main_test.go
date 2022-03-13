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
		fixturesDir = filepath.Join("..", "fixtures")
		releasesDir = filepath.Join(fixturesDir, "releases")
		outDir      = t.TempDir()
	)

	cmd := exec.Command("go", "build", "-o", filepath.Join(binDir, "protoc-gen-go-baseconvert"), ".")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("unable to build plugin executable: %v", err)
	}

	if err := filepath.WalkDir(releasesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".proto" {
			return err
		}
		relPath, err := filepath.Rel(releasesDir, path)
		if err != nil {
			return err
		}
		cmd = exec.Command(
			"protoc",
			"-I"+releasesDir,
			"-I..",
			"--go_out="+outDir,
			"--go_opt=paths=source_relative",
			"--go-grpc_out="+outDir,
			"--go-grpc_opt=paths=source_relative",
			"--go-baseconvert_out="+outDir,
			"--go-baseconvert_opt=base=github.com/devnev/proto-releases/fixtures,paths=source_relative",
			relPath,
		)
		cmd.Env = append(os.Environ(), "PATH="+binDir+":"+os.Getenv("PATH"))
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}); err != nil {
		t.Errorf("failed to find files: %v", err)
	}

	cmd = exec.Command("diff", "--unified", "--recursive", "--exclude=*.proto", outDir, releasesDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("error from %q: %v", cmd.Args, err)
	}
}
