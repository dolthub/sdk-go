package management_models

import (
	"time"
)

type BackOfficeInstallationFile struct {
	Id          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Version     string
	FullLink    string
	Filename    string
	ReleaseDate string
	Enabled     bool
	HashMd5     string
	Environment string
	Product     int64
}
