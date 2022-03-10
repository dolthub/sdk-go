package core_models

type TrialKeyResponse struct {
	LicenseType LicenseType `json:"license_type"`
	IsTrial     bool        `json:"is_trial"`
	License     string      `json:"license"`
	LicenseUser string      `json:"license_user"`
}
