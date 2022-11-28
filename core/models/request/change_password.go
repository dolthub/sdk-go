package core_request

type ChangePasswordRequest struct {
	PasswordAuth
	NewPassword string `json:"new_password,omitempty"`
}
