package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// RootDir returns the absolute paths to the root directory of the repository.
func RootDir() (string, error) {
	buf, err := execGit("rev-parse", "--show-toplevel")
	if err != nil {
		return "", fmt.Errorf("failed to find path to root directory of git repositroy: %w", err)
	}
	return strings.TrimSpace(buf.String()), nil
}

func SetHooksPath(p string) error {
	_, err := execGit("config", "core.hooksPath", p)
	if err != nil {
		return fmt.Errorf("failed to set core.hookPath: %w", err)
	}
	return nil
}

func UnsetHooksPath() error {
	_, err := execGit("config", "--unset", "core.hooksPath")
	if err != nil {
		return fmt.Errorf("failed to unset core.hookPath: %w", err)
	}
	return nil
}

func StagedFiles() ([]string, error) {
	buf, err := execGit("diff", "--name-only", "--staged", "--diff-filter=ACMR")
	if err != nil {
		return nil, fmt.Errorf("failed to get staged files: %w", err)
	}
	return strings.Split(buf.String(), "\n"), nil
}

func execGit(args ...string) (*bytes.Buffer, error) {
	cmd := exec.Command("git", args...)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		argsStr := strings.Join(args, " ")
		return nil, fmt.Errorf("failed to run 'git %s', stderr: %s, error: %w", argsStr, stderr.String(), err)
	}
	return stdout, nil
}
