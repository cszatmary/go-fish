package main

import (
	"github.com/cszatma/go-fish/cmd"
	p "github.com/cszatma/printer"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		p.ExitFailure(err.Error())
	}
}
