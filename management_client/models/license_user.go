package management_models

import (
	"time"
)

type LicenseUser struct {
	Id                int64
	LastLogin         time.Time
	Email             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	IsActive          bool
	IsStaff           bool
	FirstName         string
	LastName          string
	PhoneNumber       string
	SubscribedToEmail bool
	AcquiredConsent   time.Time
	IsManager         bool
	InitialPassword   string
	IsInitialPassword bool
}
