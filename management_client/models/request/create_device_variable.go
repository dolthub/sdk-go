package management_request

type CreateDeviceVariableRequest struct {
	Variable string `json:"variable,omitempty"`
	Value    string `json:"value,omitempty"`
	Device   int64  `json:"device,omitempty"`
}
