package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/{{ .project.github_username }}/{{ .project.folder_name }}/internal/config"
	"github.com/{{ .project.github_username }}/{{ .project.folder_name }}/internal/server"
)

func runServer(configPath string) {
	// Generate application configuration
	config, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the {{ .project.name }} service",
	Run: func(cmd *cobra.Command, args []string) {
		runServer(configPath)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&configPath, "config", "", "Path to the config file (required)")
	startCmd.MarkFlagRequired("config")
}
