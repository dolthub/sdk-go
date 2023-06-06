package management_models

import (
	"time"
)

type LicenseUser struct {
	Id                int       `json:"id,omitempty"`
	LastLogin         time.Time `json:"last_login"`
	Email             string    `json:"email,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsActive          bool      `json:"is_active,omitempty"`
	IsStaff           bool      `json:"is_staff,omitempty"`
	FirstName         string    `json:"first_name,omitempty"`
	LastName          string    `json:"last_name,omitempty"`
	PhoneNumber       string    `json:"phone_number,omitempty"`
	SubscribedToEmail bool      `json:"subscribed_to_email,omitempty"`
	AcquiredConsent   time.Time `json:"acquired_consent"`
	IsManager         bool      `json:"is_manager,omitempty"`
	InitialPassword   string    `json:"initial_password,omitempty"`
	IsInitialPassword bool      `json:"is_initial_password,omitempty"`
}
