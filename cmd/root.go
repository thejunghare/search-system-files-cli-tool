package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "Search files",
	Short: "A CLI tool is used to search files in system.",
}
