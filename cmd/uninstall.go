package cmd

import (
	"fmt"
	"os"

	"github.com/cszatma/go-fish/util"
	p "github.com/cszatma/printer"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the git hooks",
	Run:   runUninstall,
}

func init() {
	RootCmd.AddCommand(uninstallCmd)
}

func runUninstall(cmd *cobra.Command, args []string) {
	fmt.Println(p.Cyan("Uninstalling git hooks..."))
	_, gitDir, err := util.GitRevParse()

	if err != nil {
		p.ExitFailure(p.Red("Error: Unable to find git directory"))
	}

	uninstall(gitDir)

	fmt.Println(p.Green("Successfully uninstalled Git hooks! 🎣"))
}

func uninstall(gitDir string) {
	hooksPath := gitDir + "/hooks"
	err := util.RemoveHooks(hooksPath)

	if err != nil {
		util.VerboseFprintln(os.Stderr, err)
		p.ExitFailure(p.Red("Error: Unable to remove hooks"))
	}
}
