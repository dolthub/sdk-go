package management_request

type SearchDeviceVariablesRequest struct {
	Limit  int   `json:"limit,omitempty"`
	Offset int   `json:"offset,omitempty"`
	Device int64 `json:"device,omitempty"`
}
