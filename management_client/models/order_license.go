package management_models

import (
	core_models "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

type OrderLicense struct {
	User                    UserBasedLicense
	Key                     string
	MaxActivations          int
	ValidDuration           string
	EnableMaintenancePeriod bool
	MaintenanceDuration     string
	LicenseType             core_models.LicenseType
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
