package core_models

type TrialKeyResponse struct {
	LicenseType LicenseType `json:"license_type,omitempty"`
	IsTrial     bool        `json:"is_trial,omitempty"`
	License     string      `json:"license,omitempty"`
	LicenseUser string      `json:"license_user,omitempty"`
}
