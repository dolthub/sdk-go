package management_request

import management_models "gitlab.com/l3178/sdk-go/management_client/models"

type CreateInstallationFileRequest struct {
	Environment management_models.Environment `json:"environment,omitempty"`
	HashMd5     string                        `json:"hash_md_5,omitempty"`
	FullLink    string                        `json:"full_link,omitempty"`
	Product     int64                         `json:"product,omitempty"`
	ReleaseDate string                        `json:"release_date,omitempty"`
	Version     string                        `json:"version,omitempty"`
}
