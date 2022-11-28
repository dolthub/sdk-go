package core_models

type VersionsResponse struct {
	Versions []Version
}

type Version struct {
	Version     string `json:"version,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
}
