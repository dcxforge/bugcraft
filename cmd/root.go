package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bugcraft",
	Short: "A cozy terminal break game for developers",
	Long:  "BugCraft combines cron-style farming, stack crafting, regex puzzles, and ASCII bug hunts.",
}

func Execute() error {
	return rootCmd.Execute()
}
