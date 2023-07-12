package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "Search files",
	Short: "A CLI tool is used to search files in system.",
}

func init() {
	rootCmd.AddCommand(rootCmd)
}
