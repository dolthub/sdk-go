package management_request

type CreateDeviceVariableRequest struct {
	Variable string
	Value    string
	Device   int64
}
