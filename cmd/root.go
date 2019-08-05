package cmd

import (
	"fmt"

	"github.com/cszatma/go-fish/config"
	"github.com/cszatma/go-fish/util"
	p "github.com/cszatma/printer"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var RootCmd = &cobra.Command{
	Use:              "go-fish",
	Version:          version,
	Short:            "go-fish is a CLI for easily creating git hooks 🎣",
	PersistentPreRun: setup,
}

var path string
var force bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&util.Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&path, "path", "p", ".", "path to config file")
}

func setup(cmd *cobra.Command, args []string) {
	configPath := fmt.Sprintf("%s/go-fish.yml", path)
	util.VerbosePrintf("Searching for config at path %s\n", configPath)
	err := config.Init(configPath)

	if err != nil {
		fmt.Println(p.Yellow("Unable to read config file, using defaults"))
	}
}
