package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cszatma/go-fish/config"
	"github.com/cszatma/go-fish/fatal"
	"github.com/cszatma/go-fish/git"
	"github.com/cszatma/go-fish/util"
	p "github.com/cszatma/printer"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the specified git hook",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hookName := args[0]
		util.VerbosePrintf("Running hook: %s\n", hookName)

		conf := config.All()
		hook, exists := conf.Hooks[hookName]

		if !exists {
			util.VerbosePrintf("No action defined for %s in config, skipping\n", hookName)
			return
		}

		util.VerbosePrintln("Finding root directory of git repo")
		rootDir, _, err := git.RootDir()
		if err != nil {
			fatal.ExitErr(err, "Unable to find git directory")
		}

		fmt.Printf(p.Cyan("🎣 go-fish > %s\n"), hookName)

		hookArgs := strings.Fields(hook.Run)
		if len(hookArgs) == 0 {
			fatal.ExitErr(err, "Run field cannot be empty")
		}

		fmt.Printf("Running: %s\n", hook.Run)
		err = util.Exec(hookArgs[0], rootDir, hookArgs[1:]...)
		if err != nil {
			fmt.Println("Hook failed with error:")
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
