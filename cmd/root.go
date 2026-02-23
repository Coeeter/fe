package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "fe [directory]",
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil && len(args) == 0 {
			cmd.PrintErrf("Error getting current working directory: %v\n", err)
			return
		}

		if len(args) > 0 {
			dir = args[0]
		}

		isAbs := filepath.IsAbs(dir)
		if !isAbs {
			absDir, err := filepath.Abs(dir)
			if err != nil {
				cmd.PrintErrf("Error resolving absolute path: %v\n", err)
			}
			dir = absDir
		}

		cmd.Printf("Processing directory: %s\n", dir)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
