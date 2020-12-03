package cmd

import (
	"github.com/TouchBistro/goutils/fatal"
	"github.com/cszatmary/go-fish/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

type rootOptions struct {
	verbose bool
}

var rootOpts rootOptions

var rootCmd = &cobra.Command{
	Use:     "go-fish",
	Version: version,
	Short:   "go-fish is a CLI for easily creating git hooks 🎣",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if rootOpts.verbose {
			log.SetLevel(log.DebugLevel)
		}
		fatal.ShowStackTraces(rootOpts.verbose)
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
		})

		err := config.Init()
		if err != nil {
			fatal.ExitErr(err, "Failed to initialize config")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootOpts.verbose, "verbose", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fatal.ExitErr(err, "Failed executing command.")
	}
}
