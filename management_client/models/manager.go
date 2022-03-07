package management_models

type Manager struct {
	Id                int64
	Email             string
	IsInitialPassword bool
	InitialPassword   string
	FirstName         string
	LastName          string
	PhoneNumber       string
	NewUser           bool
	Password          string
}
