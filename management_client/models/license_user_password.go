package management_models

type LicenseUserPassword struct {
	Email       string `json:"email,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}
