package util

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// FileOrDirExists returns a bool indicating if a file or directory at the given path exists.
func FileOrDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func Exec(name, dir string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "Exec failed to run %s %s", name, arg)
	}

	return nil
}

func ExecOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	return stdout.String(), errors.Wrapf(err, "exec failed for command %s: %s", name, stderr.String())
}

// IsCI returns true if the current environment is CI.
func IsCI() bool {
	return os.Getenv("CI") != "" ||
		os.Getenv("CONTINUOUS_INTEGRATION") != "" ||
		os.Getenv("BUILD_NUMBER") != "" ||
		os.Getenv("RUN_ID") != ""
}
