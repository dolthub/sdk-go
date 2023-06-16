package core_request

import (
	"gitlab.com/l3178/sdk-go/core/auth"
)

type LicenseRequest struct {
	HardwareId string `json:"hardware_id,omitempty"`
	Product    string `json:"product,omitempty"`
	auth.Auth
}
