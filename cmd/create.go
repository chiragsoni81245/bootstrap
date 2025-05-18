package cmd

import (
	"log"

	"github.com/chiragsoni81245/bootstrap/internal/config"
	"github.com/chiragsoni81245/bootstrap/internal/generator"
	"github.com/spf13/cobra"
)

func create(templateName string, configPath string) {
	// Generate application configuration
	config, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

    err = generator.CreateNewProject(templateName, *config)
	if err != nil {
		log.Fatal(err)
	}
}

var templateName string
var configPath string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the boilerplate project with given configurations",
	Run: func(cmd *cobra.Command, args []string) {
		create(templateName, configPath)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&templateName, "template-name", "", "Name of the template you want to use (required)")
	createCmd.MarkFlagRequired("template-name")
	createCmd.Flags().StringVar(&configPath, "config", "", "Path to the config file (required)")
	createCmd.MarkFlagRequired("config")
}
