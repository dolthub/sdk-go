package core_models

type TrialKeyResponse struct {
	LicenseType LicenseType
	IsTrial     bool
	License     string
	LicenseUser string
}
