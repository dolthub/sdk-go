package core_models

type VersionsResponse struct {
	Versions []Version
}

type Version struct {
	Version     string
	ReleaseDate string
}
