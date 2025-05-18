package types

import (
	"database/sql"

	"github.com/{{ .project.github_username }}/{{ .project.folder_name }}/internal/config"
)

type Server struct {
	Config *config.Config
	DB     *sql.DB
}
