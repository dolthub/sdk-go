package management_response

type MultipleAssignmentResponse struct {
	Email    string `json:"email,omitempty"`
	NewUser  bool   `json:"new_user,omitempty"`
	Password string `json:"password,omitempty"`
}
