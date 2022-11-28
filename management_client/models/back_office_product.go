package management_models

import (
	core_models "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

type BackOfficeProduct struct {
	Id                          int                          `json:"id,omitempty"`
	UpgradeFrom                 []int64                      `json:"upgrade_from,omitempty"`
	ProductFeatures             []ProductFeature             `json:"product_features,omitempty"`
	CustomFields                []BackOfficeCustomField      `json:"custom_fields,omitempty"`
	BackOfficeInstallationFiles []BackOfficeInstallationFile `json:"back_office_installation_files,omitempty"`
	CreatedAt                   time.Time                    `json:"created_at"`
	UpdatedAt                   time.Time                    `json:"updated_at"`
	ProductName                 string                       `json:"product_name,omitempty"`
	ShortCode                   string                       `json:"short_code,omitempty"`
	Active                      bool                         `json:"active,omitempty"`
	ValidDuration               string                       `json:"valid_duration,omitempty"`
	AllowTrial                  bool                         `json:"allow_trial,omitempty"`
	TrialDays                   int                          `json:"trial_days,omitempty"`
	MaxActivations              int                          `json:"max_activations,omitempty"`
	HardwareIdRequired          bool                         `json:"hardware_id_required,omitempty"`
	IsUpgrade                   bool                         `json:"is_upgrade,omitempty"`
	SubscriptionDuration        string                       `json:"subscription_duration,omitempty"`
	EnableMaintenancePeriod     bool                         `json:"enable_maintenance_period,omitempty"`
	MaintenanceDuration         string                       `json:"maintenance_duration,omitempty"`
	IsFloating                  bool                         `json:"is_floating,omitempty"`
	IsFloatingCloud             bool                         `json:"is_floating_cloud,omitempty"`
	FloatingUsers               int                          `json:"floating_users,omitempty"`
	FloatingTimeout             int                          `json:"floating_timeout,omitempty"`
	DefaultLicenseType          core_models.LicenseType      `json:"default_license_type,omitempty"`
	IsUserLocked                bool                         `json:"is_user_locked,omitempty"`
	IsNodeLocked                bool                         `json:"is_node_locked,omitempty"`
	MaxConsumptions             int                          `json:"max_consumptions,omitempty"`
	AuthorizationMethod         string                       `json:"authorization_method,omitempty"`
	PreventVm                   bool                         `json:"prevent_vm,omitempty"`
	AllowOverages               bool                         `json:"allow_overages,omitempty"`
	MaxOverages                 int                          `json:"max_overages,omitempty"`
	ResetConsumption            bool                         `json:"reset_consumption,omitempty"`
	ConsumptionPeriod           string                       `json:"consumption_period,omitempty"`
	IsArchived                  bool                         `json:"is_archived,omitempty"`
	Company                     int                          `json:"company,omitempty"`
}
