package management_request

type AssignUserToLicenseRequest struct {
	Email       string
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	IsManager   bool
}
