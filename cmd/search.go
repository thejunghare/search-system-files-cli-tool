package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "Search",
	Short: "Search system files",
	Run:   runCmd,
}

var filePath string

func runCmd(cmd *cobra.Command, args []string) {
	if filePath == "" {
		fmt.Println("Please provide a file path")
		return
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", filePath)
	} else {
		absPath, _ := filepath.Abs(filePath)
		fmt.Printf("File %s exists at %s\n", filePath, absPath)
	}
}

func init() {
	searchCmd.Flags().StringVarP(&filePath, "path", "p", "", "File path")
	RootCmd.AddCommand(searchCmd)
}
