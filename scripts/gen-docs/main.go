package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/erdaltsksn/git-bump/commands"
)

func main() {
	commands.GetRootCmd().DisableAutoGenTag = true

	if err := doc.GenMarkdownTree(commands.GetRootCmd(), "./docs"); err != nil {
		log.Fatal(err)
	}
}
