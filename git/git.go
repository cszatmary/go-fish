package git

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/TouchBistro/goutils/command"
	"github.com/pkg/errors"
)

// RootDir returns the absolute paths to the root directory of the repo
// and the .git directory.
func RootDir() (string, error) {
	buf := &bytes.Buffer{}
	args := []string{"rev-parse", "--show-toplevel"}
	err := command.Exec("git", args, "git-rev-parse", func(cmd *exec.Cmd) {
		cmd.Stdout = buf
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to find path to root dir of git repo")
	}

	return strings.TrimSpace(buf.String()), nil
}

// GitDir returns the absolute path to the .git directory.
func GitDir() (string, error) {
	buf := &bytes.Buffer{}
	args := []string{"rev-parse", "--absolute-git-dir"}
	err := command.Exec("git", args, "git-rev-parse", func(cmd *exec.Cmd) {
		cmd.Stdout = buf
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to find path to .git dir")
	}

	return strings.TrimSpace(buf.String()), nil
}

// StagedFiles returns a slice of paths to staged files.
func StagedFiles() ([]string, error) {
	buf := &bytes.Buffer{}
	args := []string{"diff", "--name-only", "--staged", "--diff-filter=ACMR"}
	err := command.Exec("git", args, "git-diff", func(cmd *exec.Cmd) {
		cmd.Stdout = buf
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get staged files")
	}

	files := strings.Split(buf.String(), "\n")
	return files, nil
}
