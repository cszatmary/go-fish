package cmd

import (
	"fmt"
	"os"

	"github.com/cszatma/go-fish/config"
	"github.com/cszatma/go-fish/util"
	p "github.com/cszatma/printer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the git hooks",
	Run:   runInstall,
}

func init() {
	RootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) {
	fmt.Println(p.Cyan("Installing git hooks..."))
	topLevelDir, gitDir, err := util.GitRevParse()

	if err != nil {
		p.ExitFailure(p.Red("Error: Unable to find git directory"))
	}

	path, err := os.Executable()

	if err != nil {
		p.ExitFailure(p.Red("Error: Unable to find path of go-fish"))
	}

	install(topLevelDir, gitDir, path)

	fmt.Println(p.Green("Successfully installed Git hooks! Enjoy! 🎣"))
}

func install(topLevelDir, gitDir, goFishPath string) {
	skip := os.Getenv("GOFISH_SKIP_INSTALL")
	if skip == "true" || skip == "1" {
		fmt.Println("GOFISH_SKIP_INSTALL is set, skipping Git hooks installation")
		return
	}

	if util.IsCI() && config.All().SkipCI {
		fmt.Println("CI detected, skipping Git hooks installation")
	}

	hooksPath := gitDir + "/hooks"
	if _, err := os.Stat(hooksPath); os.IsNotExist(err) {
		err = os.Mkdir(hooksPath, 0755)
		if err != nil {
			p.ExitFailure(p.Red("Error: Unable to create hooks directory"))
		}
	}

	script, err := util.RenderScript(topLevelDir + "/go-fish")

	if err != nil {
		p.ExitFailure(p.Red("Error: Unable to generate hook script"))
	}

	err = util.CreateHooks(hooksPath, script)

	if err != nil {
		util.VerboseFprintln(os.Stderr, err)
		p.ExitFailure(p.Red("Error: Unable to create hooks"))
	}
}
