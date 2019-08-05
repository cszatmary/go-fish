package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cszatma/go-fish/config"
	"github.com/cszatma/go-fish/util"
	p "github.com/cszatma/printer"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the specified git hook",
	Args:  cobra.MinimumNArgs(1),
	Run:   runRun,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

// Jokes name but it will have to do
func runRun(cmd *cobra.Command, args []string) {
	hookName := args[0]
	util.VerbosePrintf("Running hook: %s\n", hookName)

	conf := config.All()
	hook, exists := conf.Hooks[hookName]

	if !exists {
		util.VerbosePrintf("No action defined for %s in config, skipping\n", hookName)
		os.Exit(0)
	}

	topLevelDir, _, err := util.GitRevParse()
	if err != nil {
		p.ExitFailure(p.Red("Error: Unable to find root directory"))
	}

	fmt.Printf(p.Cyan("🎣 go-fish > %s\n"), hookName)

	hookArgs := strings.Fields(hook.Run)
	if len(hookArgs) == 0 {
		p.ExitFailure(p.Red("run key cannot be empty"))
	}
	fmt.Printf("Running: %s\n", hook.Run)

	execCmd := exec.Command(hookArgs[0], hookArgs[1:]...)
	execCmd.Dir = topLevelDir

	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	err = execCmd.Run()
	if err != nil {
		fmt.Println("Hook failed with error:")
		fmt.Println(err)
		os.Exit(1)
	}
}
