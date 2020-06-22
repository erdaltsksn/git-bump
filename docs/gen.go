package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/erdaltsksn/git-bump/cmd"
)

func main() {
	err := doc.GenMarkdownTree(cmd.GetRootCmd(), "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
