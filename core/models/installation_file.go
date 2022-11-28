package core_models

import (
	"time"
)

type InstallationFileResponse struct {
	InstallationFile string    `json:"installation_file,omitempty"`
	Version          string    `json:"version,omitempty"`
	RequiresVersion  string    `json:"requires_version,omitempty"`
	HashMd5          string    `json:"hash_md_5,omitempty"`
	ReleaseDate      time.Time `json:"release_date"`
}
