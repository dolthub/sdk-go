package response

type VersionsResponse struct {
	Versions []Version
}

type Version struct {
	Version     string `json:"version"`
	ReleaseDate string `json:"release_date"`
}
