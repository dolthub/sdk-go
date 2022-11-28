package management_models

type Manager struct {
	Id                int64  `json:"id,omitempty"`
	Email             string `json:"email,omitempty"`
	IsInitialPassword bool   `json:"is_initial_password,omitempty"`
	InitialPassword   string `json:"initial_password,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	PhoneNumber       string `json:"phone_number,omitempty"`
	NewUser           bool   `json:"new_user,omitempty"`
	Password          string `json:"password,omitempty"`
}
