package management_models

import (
	"time"
)

type Device struct {
	Id               int                        `json:"id,omitempty"`
	DeviceVariables  []BackOfficeDeviceVariable `json:"device_variables,omitempty"`
	CreatedAt        time.Time                  `json:"created_at"`
	UpdatedAt        time.Time                  `json:"updated_at"`
	HardwareId       string                     `json:"hardware_id,omitempty"`
	TimeActivated    time.Time                  `json:"time_activated"`
	TimeDeactivated  time.Time                  `json:"time_deactivated"`
	ValidityPeriod   time.Time                  `json:"validity_period"`
	DeviceActive     bool                       `json:"device_active,omitempty"`
	Type             string                     `json:"type,omitempty"`
	Os               string                     `json:"os,omitempty"`
	SdkBuildVersion  string                     `json:"sdk_build_version,omitempty"`
	LastCheck        time.Time                  `json:"last_check"`
	AppVer           string                     `json:"app_ver,omitempty"`
	Hostname         string                     `json:"hostname,omitempty"`
	Ip               string                     `json:"ip,omitempty"`
	ExternalIp       string                     `json:"external_ip,omitempty"`
	MacAddress       string                     `json:"mac_address,omitempty"`
	IsVm             bool                       `json:"is_vm,omitempty"`
	VmInfo           string                     `json:"vm_info,omitempty"`
	FloatingLastSeen time.Time                  `json:"floating_last_seen"`
	FloatingInUse    bool                       `json:"floating_in_use,omitempty"`
	Blacklisted      bool                       `json:"blacklisted,omitempty"`
	License          DeviceLicense              `json:"license"`
}
