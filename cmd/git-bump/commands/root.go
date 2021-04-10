package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/erdaltsksn/cui"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type semver struct {
	Name        string
	Version     string
	Description string
}

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "git-bump",
	Short: "Bump the app version",
	Long: `Bump the app version using git tags that follows the semantic
versioning rules.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(".git"); os.IsNotExist(err) {
			cui.Error("Not a Git repository", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Get prefix
		var versionPrefix string
		versionPrefixConf, err := exec.Command("git", "config", "--get", "git-bump.prefix").Output()
		if err != nil {
			cui.Info("Using default semver prefix")
			versionPrefix = "v"
		} else {
			versionPrefix = strings.Trim(string(versionPrefixConf), "\n")
		}

		// Check whether we need to bump or not
		status, err := exec.Command("git", "tag", "--contains", "HEAD", "--list", versionPrefix+"*.*.*").Output()
		if err != nil {
			cui.Error("Couldn't get the status (Git Tags) at HEAD", err)
		}
		versionMatched := strings.Trim(string(status), "\n")
		if versionMatched != "" {
			cui.Warning("You don't need to bump the version", versionMatched)
		}

		// Get all tags that match semantic versioning
		tags, err := exec.Command("git", "tag", "--list", versionPrefix+"*.*.*", "--sort=-v:refname").Output()
		if err != nil {
			cui.Error("Couldn't get the Git Tags to bump the version", err)
		}

		// Initiate the first version if there aren't any release before
		if len(tags) == 0 {
			if err := exec.Command("git", "tag", versionPrefix+"0.1.0").Run(); err != nil {
				cui.Error("Couldn't get the Git Tags to bump the version", err)
			}

			cui.Success("The Semantic Version is initiated", fmt.Sprintf("Current Version: %s0.1.0", versionPrefix))
		}

		// Calculate the current and the next version
		currentVersion := strings.Split(strings.Trim(string(tags), "\n"), "\n")[0]
		nextVersion := strings.Split(currentVersion, ".")
		var currentMajor int
		if versionPrefix != "" {
			currentMajor, _ = strconv.Atoi(strings.Split(nextVersion[0], "")[1])
		} else {
			currentMajor, _ = strconv.Atoi(nextVersion[0])
		}
		currentMinor, _ := strconv.Atoi(nextVersion[1])
		currentPatch, _ := strconv.Atoi(nextVersion[2])
		nextMajor := fmt.Sprintf("%s%d.%d.%d", versionPrefix, currentMajor+1, 0, 0)
		nextMinor := fmt.Sprintf("%s%d.%d.%d", versionPrefix, currentMajor, currentMinor+1, 0)
		nextPatch := fmt.Sprintf("%s%d.%d.%d", versionPrefix, currentMajor, currentMinor, currentPatch+1)

		// Build up CLI UI
		semvers := []semver{
			{
				Name:        "\U0001F4A5 Major",
				Version:     nextMajor,
				Description: "MAJOR version when you make incompatible API changes.",
			},
			{
				Name:        "\U0001f389 Minor",
				Version:     nextMinor,
				Description: "MINOR version when you add functionality in a backwards compatible manner.",
			},
			{
				Name:        "\U0001F41B Patch",
				Version:     nextPatch,
				Description: "PATCH version when you make backwards compatible bug fixes.",
			},
		}

		cui.Info("Previous Version: " + currentVersion)
		prompt := promptui.Select{
			Label: "How do you want to bump it",
			Items: semvers,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}?",
				Active:   "> {{ .Name | cyan }} ({{ .Version | red }})",
				Inactive: "  {{ .Name | faint }} ({{ .Version | faint }})",
				Selected: "* {{ .Name | cyan }} ({{ .Version | red }})",
				Details:  "---------- Details ----------\n{{ .Description | faint }}",
			},
			Size: 3,
		}

		selected, _, err := prompt.Run()
		if err != nil {
			cui.Error("Interactive UI failed", err)
		}
		bumpedVersion := semvers[selected].Version

		// Bump the version according to selected
		if err := exec.Command("git", "tag", bumpedVersion).Run(); err != nil {
			cui.Error("Couldn't get the Git Tags to bump the version", err)
		}

		// Success
		cui.Success(
			"The Semantic Version is bumped",
			fmt.Sprintf("Current Version: %s", bumpedVersion),
		)
	},
}
