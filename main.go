package main

import "github.com/erdaltsksn/git-bump/cmd"

var version = "unknown"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
