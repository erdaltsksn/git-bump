package main

import (
	"github.com/erdaltsksn/cui"

	"github.com/erdaltsksn/git-bump/cmd/git-bump/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		cui.Error("Something went wrong", err)
	}
}
