package management_request

type AssignMultipleUsersRequest struct {
	Emails    []AssignUserData `json:"emails,omitempty"`
	SendEmail bool             `json:"send_email,omitempty"`
}
