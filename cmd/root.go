package cmd

import (
	"os"

	"github.com/cszatma/go-fish/config"
	"github.com/cszatma/go-fish/fatal"
	"github.com/cszatma/go-fish/util"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var (
	configPath string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:     "go-fish",
	Version: version,
	Short:   "go-fish is a CLI for easily creating git hooks 🎣",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&configPath, "path", "p", "go-fish.yml", "path to config file")

	cobra.OnInitialize(func() {
		util.SetVerboseMode(verbose)
		fatal.ShowStackTraces(verbose)

		util.VerbosePrintf("Reading config file %s\n", configPath)
		if !util.FileOrDirExists(configPath) {
			fatal.Exitf("No such file %s", configPath)
		}

		file, err := os.Open(configPath)
		if err != nil {
			fatal.ExitErrf(err, "Failed to open config file at path", configPath)
		}
		defer file.Close()

		err = config.Init(file)
		if err != nil {
			fatal.ExitErr(err, "Failed reading config file.")
		}
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fatal.ExitErr(err, "Failed executing command.")
	}
}
