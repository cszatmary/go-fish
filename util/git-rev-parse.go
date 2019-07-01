package util

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	p "github.com/cszatma/printer"
)

// GitRevParse finds the absolute paths to the
// project root directory and the git directory.
func GitRevParse() (topLevelDir, gitDir string, err error) {
	VerbosePrintln("Finding git directory")

	cmd := exec.Command("git", "rev-parse", "--show-toplevel", "--absolute-git-dir")

	cmdOut := &bytes.Buffer{}
	cmdErr := &bytes.Buffer{}
	cmd.Stdout = cmdOut
	cmd.Stderr = cmdErr

	err = cmd.Run()

	if err != nil {
		VerboseFprintln(os.Stderr, cmdErr.String())
		VerboseFprintln(os.Stderr, err)
		p.Eprintln(p.Red("Error: Unabled to find git directory"))
		return "", "", err
	}

	result := cmdOut.String()
	VerbosePrintln(result)

	split := strings.Split(result, "\n")

	return split[0], split[1], nil
}
