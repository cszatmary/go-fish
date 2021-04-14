package hooks

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cszatmary/go-fish/git"
)

//go:embed go-fish.sh
var gofishScript []byte

func Install() error {
	dir, err := hooksDir()
	if err != nil {
		return fmt.Errorf("failed to install hooks: %w", err)
	}
	gofishDir := filepath.Join(dir, "go-fish")
	if err := os.MkdirAll(gofishDir, 0o755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", gofishDir, err)
	}

	// Create gitignore for gofish stuff
	gitignorePath := filepath.Join(dir, ".gitignore")
	if err := os.WriteFile(gitignorePath, []byte("go-fish\n"), 0o644); err != nil {
		return fmt.Errorf("failed to create %s: %w", gitignorePath, err)
	}
	// Create .hooks/go-fish/go-fish.sh
	scriptPath := filepath.Join(gofishDir, "go-fish.sh")
	if err := os.WriteFile(scriptPath, gofishScript, 0o755); err != nil {
		return fmt.Errorf("failed to create %s: %w", scriptPath, err)
	}
	if err := git.SetHooksPath(dir); err != nil {
		return fmt.Errorf("failed to install hooks: %w", err)
	}
	return nil
}

func Uninstall() error {
	if err := git.UnsetHooksPath(); err != nil {
		return fmt.Errorf("failed to uninstall git hooks: %w", err)
	}
	return nil
}

// Create creates a new hook script template.
// It also checks that name is a valid git hook name.
func Create(name string) error {
	if ok := hooks[name]; !ok {
		return fmt.Errorf("%s is not a valid git hook", name)
	}
	dir, err := hooksDir()
	if err != nil {
		return fmt.Errorf("failed to create hook %s: %w", name, err)
	}
	fp := filepath.Join(dir, name)
	const tmpl = `#!/bin/sh
. "$(dirname "$0")/go-fish/go-fish.sh"

# Add hook commands

`
	if err := os.WriteFile(fp, []byte(tmpl), 0o755); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fp, err)
	}
	return nil
}

// hooksDir returns the path to the .hooks directory in the root of the git repo.
func hooksDir() (string, error) {
	rootDir, err := git.RootDir()
	if err != nil {
		return "", fmt.Errorf("failed to get hooks directory: %w", err)
	}
	return filepath.Join(rootDir, ".hooks"), nil
}

// List of hooks supported by git, used for validation
var hooks = map[string]bool{
	"applypatch-msg":     true,
	"pre-applypatch":     true,
	"post-applypatch":    true,
	"pre-commit":         true,
	"prepare-commit-msg": true,
	"commit-msg":         true,
	"post-commit":        true,
	"pre-rebase":         true,
	"post-checkout":      true,
	"post-merge":         true,
	"pre-push":           true,
	"pre-receive":        true,
	"update":             true,
	"post-receive":       true,
	"post-update":        true,
	"push-to-checkout":   true,
	"pre-auto-gc":        true,
	"post-rewrite":       true,
	"sendemail-validate": true,
}
