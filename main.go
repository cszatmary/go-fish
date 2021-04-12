package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/cszatmary/go-fish/hooks"
)

var version = "master"

func main() {
	// Main flags
	flags := flag.NewFlagSet("go-fish", flag.ExitOnError)
	versionFlag := flags.Bool("version", false, "version for go-fish")
	flags.Usage = func() {
		var b strings.Builder
		b.WriteString(`go-fish is a CLI for easily creating git hooks 🎣

Usage:
  go-fish [command]

Available Commands:
  install       Install git hooks
  uninstall     Uninstall git hooks
  create [hook] Create a git hook script

Flags:
`)
		flags.SetOutput(&b)
		flags.PrintDefaults()
		fmt.Fprint(os.Stderr, b.String())
	}

	// Ignore error because it is set to ExitOnError
	_ = flags.Parse(os.Args[1:])
	if *versionFlag {
		fmt.Printf("go-fish version %s\n", version)
		os.Exit(0)
	}

	logger := log.New(os.Stderr, "", 0)
	if flags.NArg() == 0 {
		logger.Print("Error: No command specified\n\n")
		flags.Usage()
		os.Exit(1)
	}

	var runFn func()
	switch flags.Arg(0) {
	case "install":
		runFn = func() {
			if err := hooks.Install(); err != nil {
				logger.Fatalf("Failed to install hooks\nError: %v", err)
			}
			logger.Println("Git hooks installed")
		}
	case "uninstall":
		runFn = func() {
			if err := hooks.Uninstall(); err != nil {
				logger.Fatalf("Failed to uninstall hooks\nError: %v", err)
			}
			logger.Println("Git hooks uninstalled")
		}
	case "create":
		runFn = func() {
			hookName := flags.Arg(1)
			if hookName == "" {
				logger.Fatalln("Error: hook name is required")
			}
			if err := hooks.Create(hookName); err != nil {
				logger.Fatalf("Failed to create %s hook\nError: %v", hookName, err)
			}
			logger.Printf("Created %s hook\n", hookName)
		}
	default:
		logger.Fatalf("Error: unknown command %q\nRun 'go-fish --help' for usage.\n", flags.Arg(0))
	}
	runFn()
}

func init() {
	if info, available := debug.ReadBuildInfo(); available {
		version = info.Main.Version
	}
}
