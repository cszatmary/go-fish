package hooks

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInstallUninstall(t *testing.T) {
	rootDir := setupRepo(t)
	if err := Install(); err != nil {
		t.Fatalf("got install error %v", err)
	}
	buf, err := execGit("config", "core.hooksPath")
	if err != nil {
		t.Fatalf("failed to get hooks path: %v", err)
	}
	got := strings.TrimSpace(buf.String())
	want := filepath.Join(rootDir, ".hooks")
	if got != want {
		t.Fatalf("got hooks path %q, want %q", got, want)
	}
	if err := Uninstall(); err != nil {
		t.Fatalf("got uninstall error %v", err)
	}
	if _, err := execGit("config", "core.hooksPath"); err == nil {
		t.Fatalf("want non-nil error for hooks path")
	}
}

func TestInstallNoRepo(t *testing.T) {
	td := t.TempDir()
	if err := os.Chdir(td); err != nil {
		t.Fatalf("failed to chdir to %s: %v", td, err)
	}
	if err := Install(); err == nil {
		t.Error("want non-nil error")
	}
}

func TestCreate(t *testing.T) {
	rootDir := setupRepo(t)
	if err := Install(); err != nil {
		t.Fatalf("got install error %v", err)
	}
	if err := Create("pre-commit"); err != nil {
		t.Fatalf("got create error %v", err)
	}
	p := filepath.Join(rootDir, ".hooks", "pre-commit")
	if _, err := os.Stat(p); err != nil {
		t.Errorf("want hook script to exist, got error %v", err)
	}
}

func TestCreateInvalidHook(t *testing.T) {
	err := Create("pre-foo")
	if err == nil {
		t.Fatalf("want non-nil error")
	}
	want := "pre-foo is not a valid git hook"
	if err.Error() != want {
		t.Errorf("got error %v, want %s", err, want)
	}
}

func setupRepo(t *testing.T) string {
	t.Helper()
	td := t.TempDir()
	// On macOS the returned path contains a symlink which git evaluates leading
	// to differences in paths. Resolve the symlink to avoid that.
	rootDir, err := filepath.EvalSymlinks(td)
	if err != nil {
		t.Fatalf("failed to eval symlinks on %s: %v", td, err)
	}
	if err := os.Chdir(rootDir); err != nil {
		t.Fatalf("failed to chdir to %s: %v", rootDir, err)
	}
	if _, err := execGit("init"); err != nil {
		t.Fatalf("failed to create git repo: %v", err)
	}
	return rootDir
}
