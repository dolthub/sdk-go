package core_request

type LicenseRequest struct {
	HardwareId string `json:"hardware_id,omitempty"`
	Product    string `json:"product,omitempty"`
	Auth
}
