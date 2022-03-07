package request

type DeviceVariablesRequest struct {
	LicenseRequest

	Variables map[string]string `json:"variables"`
}
