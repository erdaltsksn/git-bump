package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/erdaltsksn/cui"
	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type semver struct {
	Name        string
	Version     string
	Description string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-bump",
	Short: "Bump the app version",
	Long: `Bump the app version using git tags. This application follows the
semantic versioning rules.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check whether we need to bump or not
		status, err := exec.Command("git", "tag", "--contains", "HEAD", "--list", "v*.*.*").Output()
		if err != nil {
			cui.Error("Couldn't get the status (Git Tags) at HEAD", err)
		}
		versionMatched := strings.Trim(string(status), "\n")
		if versionMatched != "" {
			cui.Warning("You don't need to bump the version", versionMatched)
		}

		// Get the tags that match semantic versioning
		tags, err := exec.Command("git", "tag", "--list", "v*.*.*", "--sort=-v:refname").Output()
		if err != nil {
			cui.Error("Couldn't get the Git Tags to bump the version", err)
		}

		// Initiate the first version if it is not exists
		if len(tags) == 0 {
			if err := exec.Command("git", "tag", "v0.1.0").Run(); err != nil {
				cui.Error("Couldn't get the Git Tags to bump the version", err)
			}

			cui.Success("The Semantic Version is initiated", "Current Version: v0.1.0")
		}

		// Calculate the current and the next versions
		latestVersion := strings.Split(strings.Trim(string(tags), "\n"), "\n")[0]
		currentVersion := strings.Split(latestVersion, ".")
		major, _ := strconv.Atoi(strings.Split(currentVersion[0], "")[1])
		minor, _ := strconv.Atoi(currentVersion[1])
		patch, _ := strconv.Atoi(currentVersion[2])
		nextMajor := fmt.Sprintf("v%d.%d.%d", major+1, 0, 0)
		nextMinor := fmt.Sprintf("v%d.%d.%d", major, minor+1, 0)
		nextPatch := fmt.Sprintf("v%d.%d.%d", major, minor, patch+1)

		// Build up CLI UI
		semvers := []semver{
			{Name: "\U0001F4A5 Major", Version: nextMajor,
				Description: "MAJOR version when you make incompatible API changes."},
			{Name: "\U0001f389 Minor", Version: nextMinor,
				Description: "MINOR version when you add functionality in a backwards compatible manner."},
			{Name: "\U0001F41B Patch", Version: nextPatch,
				Description: "PATCH version when you make backwards compatible bug fixes."},
		}

		fmt.Println(color.Warn.Sprint("Current Version: "), color.Info.Sprint(latestVersion))
		prompt := promptui.Select{
			Label: "How do you want to bump it",
			Items: semvers,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}?",
				Active:   "> {{ .Name | cyan }} ({{ .Version | red }})",
				Inactive: "  {{ .Name | faint }} ({{ .Version | faint }})",
				Selected: "* {{ .Name | cyan }} ({{ .Version | red }})",
				Details: `
---------- Details ----------
{{ .Description | faint }}`,
			},
			Size: 3,
		}

		selected, _, err := prompt.Run()
		if err != nil {
			cui.Error("Interactive UI failed", err)
		}

		// Bump the version according to selected
		if err := exec.Command("git", "tag", semvers[selected].Version).Run(); err != nil {
			cui.Error("Couldn't get the Git Tags to bump the version", err)
		}

		// Success
		cui.Success("The Semantic Version is bumped to", semvers[selected].Version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cui.Error("Something went wrong", err)
	}
}

// GetRootCmd returns the instance of root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}
