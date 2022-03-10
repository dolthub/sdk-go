package core_request

type ActivationRequest struct {
	LicenseRequest

	OsVersion  string `json:"os_version,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	IP         string `json:"ip,omitempty"`
	AppVersion string `json:"app_version,omitempty"`
	SdkVersion string `json:"sdk_version,omitempty"`
	IsVM       string `json:"is_vm,omitempty"`
	VmInfo     string `json:"vm_info,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
}

func (r LicenseRequest) ToActivationRequest() ActivationRequest {
	return ActivationRequest{LicenseRequest: r}
}
