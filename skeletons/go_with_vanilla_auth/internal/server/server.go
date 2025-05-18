package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/{{ .project.github_username }}/{{ .project.folder_name }}/internal/config"
	"github.com/{{ .project.github_username }}/{{ .project.folder_name }}/internal/database"
	"github.com/{{ .project.github_username }}/{{ .project.folder_name }}/internal/types"
)

func NewServer(config *config.Config) error {
	// Create a global logger further will be used in whole application

	// Create a global db connection which will be used in whole application
	db, err := database.GetDB(config)
	if err != nil {
		return err
	}

	server := types.Server{
		Config: config,
		DB:     db,
	}

	router := NewRouter(&server)

	log.Printf("Server started on http://localhost:%d", config.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), router)
	if err != nil {
		return err
	}

	return nil
}
