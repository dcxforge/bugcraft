package cmd

import (
	"fmt"

	"github.com/dcxforge/bugcraft/internal/app"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start BugCraft",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := app.Init()
		if err != nil {
			return err
		}
		fmt.Println("Welcome to BugCraft")
		fmt.Println("You inherited a tiny farm at the edge of a haunted codebase.")
		fmt.Println("Save directory: ", dir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
