package core_models

type CheckResponse struct {
	LicenseResponse
	InstallationFileResponse

	LicenseActive  bool `json:"license_active,omitempty"`
	LicenseEnabled bool `json:"license_enabled,omitempty"`
	IsExpired      bool `json:"is_expired,omitempty"`
}
