package management_models

type DeviceLicense struct {
	Id           int64         `json:"id,omitempty"`
	LicenseKey   string        `json:"license_key,omitempty"`
	LicenseUsers []LicenseUser `json:"license_users,omitempty"`
}
