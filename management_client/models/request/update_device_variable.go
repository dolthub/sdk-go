package management_request

type UpdateDeviceVariableRequest struct {
	Variable string `json:"variable,omitempty"`
	Value    string `json:"value,omitempty"`
}
