package management_request

import management_models "gitlab.com/l3178/sdk-go/management_client/models"

type CreateInstallationFileRequest struct {
	Environment management_models.Environment
	HashMd5     string
	FullLink    string
	Product     int64
	ReleaseDate string
	Version     string
}
