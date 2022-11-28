package core_request

type DeviceVariablesRequest struct {
	LicenseRequest

	Variables map[string]string `json:"variables,omitempty"`
}
