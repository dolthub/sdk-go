package management_models

import (
	"time"
)

type Device struct {
	Id               int
	DeviceVariables  []BackOfficeDeviceVariable
	CreatedAt        time.Time
	UpdatedAt        time.Time
	HardwareId       string
	TimeActivated    time.Time
	TimeDeactivated  time.Time
	ValidityPeriod   time.Time
	DeviceActive     bool
	Type             string
	Os               string
	SdkBuildVersion  string
	LastCheck        time.Time
	AppVer           string
	Hostname         string
	Ip               string
	ExternalIp       string
	MacAddress       string
	IsVm             bool
	VmInfo           string
	FloatingLastSeen time.Time
	FloatingInUse    bool
	Blacklisted      bool
	License          DeviceLicense
}
