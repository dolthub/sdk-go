package management_models

import (
	core_models "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

type OrderLicense struct {
	User                    UserBasedLicense        `json:"user"`
	Key                     string                  `json:"key,omitempty"`
	MaxActivations          int                     `json:"max_activations,omitempty"`
	ValidDuration           string                  `json:"valid_duration,omitempty"`
	EnableMaintenancePeriod bool                    `json:"enable_maintenance_period,omitempty"`
	MaintenanceDuration     string                  `json:"maintenance_duration,omitempty"`
	LicenseType             core_models.LicenseType `json:"license_type,omitempty"`
	ValidityPeriod          time.Time               `json:"validity_period"`
	ProductFeatures         []Feature               `json:"product_features,omitempty"`
	IsTrial                 bool                    `json:"is_trial,omitempty"`
	MaxConsumptions         int                     `json:"max_consumptions,omitempty"`
	AllowOverages           bool                    `json:"allow_overages,omitempty"`
	MaxOverages             int                     `json:"max_overages,omitempty"`
	ResetConsumption        bool                    `json:"reset_consumption,omitempty"`
	ConsumptionPeriod       string                  `json:"consumption_period,omitempty"`
	PreventVm               bool                    `json:"prevent_vm,omitempty"`
	IsFloating              bool                    `json:"is_floating,omitempty"`
	IsFloatingCloud         bool                    `json:"is_floating_cloud,omitempty"`
	FloatingUsers           int                     `json:"floating_users,omitempty"`
}
