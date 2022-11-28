package management_response

type LicenseAssignmentResponse struct {
	NewPassword string `json:"new_password,omitempty"`
	Message     string `json:"message,omitempty"`
}
