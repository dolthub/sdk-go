package response

import (
	"time"
)

type InstallationFileResponse struct {
	InstallationFile string    `json:"installation_file"`
	Version          string    `json:"version"`
	RequiresVersion  string    `json:"requires_version"`
	HashMd5          string    `json:"hash_md_5"`
	ReleaseDate      time.Time `json:"release_date"`
}
