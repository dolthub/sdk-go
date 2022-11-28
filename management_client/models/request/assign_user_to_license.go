package management_request

type AssignUserToLicenseRequest struct {
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password,omitempty"`
	IsManager   bool   `json:"is_manager,omitempty"`
}
