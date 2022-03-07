package request

type LicenseRequest struct {
	HardwareId string `json:"hardware_id"`
	Product    string `json:"product"`
	Auth
}
