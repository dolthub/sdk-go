package management_request

type AssignUserData struct {
	Email     string `json:"email,omitempty"`
	IsManager bool   `json:"is_manager,omitempty"`
}
