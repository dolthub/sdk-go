package core_models

type CheckResponse struct {
	LicenseResponse
	InstallationFileResponse

	LicenseActive  bool
	LicenseEnabled bool
	IsExpired      bool
}
