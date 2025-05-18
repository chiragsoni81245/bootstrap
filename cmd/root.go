package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap CLI an project bootstrapping tool",
	Long:  "Bootstrap CLI an project bootstrapping tool",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
