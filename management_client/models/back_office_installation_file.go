package management_models

import (
	"time"
)

type BackOfficeInstallationFile struct {
	Id          int       `json:"id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     string    `json:"version,omitempty"`
	FullLink    string    `json:"full_link,omitempty"`
	Filename    string    `json:"filename,omitempty"`
	ReleaseDate string    `json:"release_date,omitempty"`
	Enabled     bool      `json:"enabled,omitempty"`
	HashMd5     string    `json:"hash_md_5,omitempty"`
	Environment string    `json:"environment,omitempty"`
	Product     int64     `json:"product,omitempty"`
}
