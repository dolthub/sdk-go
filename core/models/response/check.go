package response

type CheckResponse struct {
	LicenseResponse
	InstallationFileResponse

	LicenseActive  bool `json:"license_active"`
	LicenseEnabled bool `json:"license_enabled"`
	IsExpired      bool `json:"is_expired"`
}
