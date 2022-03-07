package management_models

type DeviceLicense struct {
	Id           int64
	LicenseKey   string
	LicenseUsers []LicenseUser
}
