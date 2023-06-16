package core_request

import (
	"gitlab.com/l3178/sdk-go/core/auth"
)

type ChangePasswordRequest struct {
	auth.BasicAuth
	NewPassword string `json:"new_password,omitempty"`
}
