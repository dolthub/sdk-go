package management_request

import (
	"licensespring/go-sdk/management_client/models"
)

type CreateInstallationFileRequest struct {
	Environment management_models.Environment
	HashMd5     string
	FullLink    string
	Product     int64
	ReleaseDate string
	Version     string
}
