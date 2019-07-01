package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the specified git hook",
	Run:   runRun,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

// Jokes name but it will have to do
func runRun(cmd *cobra.Command, args []string) {
	// hook := args[0]
	// fmt.Printf(printer.Cyan("Running hook: %s\n"))
	fmt.Println(args)

	os.Exit(1)
}
