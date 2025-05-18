package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{ .project.folder_name }}",
	Short: "{{ .project.name }} CLI an network traffic monitoring tool",
	Long:  "{{ .project.name }} CLI an network traffic monitoring tool",
}

var configPath string

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
