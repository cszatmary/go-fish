package git

import (
	"strings"

	"github.com/cszatma/go-fish/util"
	"github.com/pkg/errors"
)

func RootDir() (string, string, error) {
	result, err := util.ExecOutput("git", "rev-parse", "--show-toplevel", "--absolute-git-dir")
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to find root dir of git repo")
	}

	parts := strings.Split(result, "\n")
	return parts[0], parts[1], nil
}
