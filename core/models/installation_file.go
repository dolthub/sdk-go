package core_models

import (
	"time"
)

type InstallationFileResponse struct {
	InstallationFile string
	Version          string
	RequiresVersion  string
	HashMd5          string
	ReleaseDate      time.Time
}
