package response

import (
	"licensespring/go-sdk/core/models"
)

type TrialKeyResponse struct {
	LicenseType models.LicenseType `json:"license_type"`
	IsTrial     bool               `json:"is_trial"`
	License     string             `json:"license"`
	LicenseUser string             `json:"license_user"`
}
