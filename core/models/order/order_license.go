package order

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models"
	"time"
)

type OrderLicense struct {
	User                    UserBasedLicense
	Key                     string
	MaxActivations          int
	ValidDuration           string
	EnableMaintenancePeriod bool
	MaintenanceDuration     string
	LicenseType             models.LicenseType
	ValidityPeriod          time.Time
	ProductFeatures         []Feature
	IsTrial                 bool
	MaxConsumptions         int
	AllowOverages           bool
	MaxOverages             int
	ResetConsumption        bool
	ConsumptionPeriod       string
	PreventVm               bool
	IsFloating              bool
	IsFloatingCloud         bool
	FloatingUsers           int
}
