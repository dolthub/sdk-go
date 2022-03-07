package management_request

type CreateInstallationFileRequest struct {
	Environment management_models.Environment
	HashMd5     string
	FullLink    string
	Product     int64
	ReleaseDate string
	Version     string
}
