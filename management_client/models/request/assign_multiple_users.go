package management_request

type AssignMultipleUsersRequest struct {
	Emails    []AssignUserData
	SendEmail bool
}
