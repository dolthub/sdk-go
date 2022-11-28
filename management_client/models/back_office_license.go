package management_models

import (
	core_models "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

type BackOfficeLicense struct {
	Id                      int64                      `json:"id,omitempty"`
	Status                  string                     `json:"status,omitempty"`
	FloatingInUseDevices    int                        `json:"floating_in_use_devices,omitempty"`
	Customer                BackOfficeCustomer         `json:"customer"`
	ProductFeatures         []BackOfficeLicenseFeature `json:"product_features,omitempty"`
	LicenseUsers            []LicenseUser              `json:"license_users,omitempty"`
	LicenseUser             LicenseUser                `json:"license_user"`
	OrderStoreId            string                     `json:"order_store_id,omitempty"`
	ActiveUpToDateDevices   int                        `json:"active_up_to_date_devices,omitempty"`
	TotalActiveDevices      int                        `json:"total_active_devices,omitempty"`
	Product                 BackOfficeLicenseProduct   `json:"product"`
	CreatedAt               time.Time                  `json:"created_at"`
	UpdatedAt               time.Time                  `json:"updated_at"`
	LicenseKey              string                     `json:"license_key,omitempty"`
	Active                  bool                       `json:"active,omitempty"`
	Enabled                 bool                       `json:"enabled,omitempty"`
	TimeActivated           time.Time                  `json:"time_activated"`
	TimeDisabled            time.Time                  `json:"time_disabled"`
	MaxActivtions           int                        `json:"max_activtions,omitempty"`
	TimesActivated          int                        `json:"times_activated,omitempty"`
	ValidDuration           string                     `json:"valid_duration,omitempty"`
	ValidityPeriod          string                     `json:"validity_period,omitempty"`
	LicenseType             core_models.LicenseType    `json:"license_type,omitempty"`
	EnableMaintenancePeriod bool                       `json:"enable_maintenance_period,omitempty"`
	MaintenanceDuration     string                     `json:"maintenance_duration,omitempty"`
	MaintenancePeriod       string                     `json:"maintenance_period,omitempty"`
	IsTrial                 bool                       `json:"is_trial,omitempty"`
	MaxConsumptions         int                        `json:"max_consumptions,omitempty"`
	TotalConsumptions       int                        `json:"total_consumptions,omitempty"`
	SubscriptionId          string                     `json:"subscription_id,omitempty"`
	LastCheck               time.Time                  `json:"last_check"`
	PreventVm               bool                       `json:"prevent_vm,omitempty"`
	AllowOverages           bool                       `json:"allow_overages,omitempty"`
	MaxOverages             int                        `json:"max_overages,omitempty"`
	ResetConsumption        bool                       `json:"reset_consumption,omitempty"`
	ConsumptionPeriod       string                     `json:"consumption_period,omitempty"`
	TrialDays               int                        `json:"trial_days,omitempty"`
	IsFloating              bool                       `json:"is_floating,omitempty"`
	IsFloatingCloud         bool                       `json:"is_floating_cloud,omitempty"`
	FloatingUsers           int                        `json:"floating_users,omitempty"`
	Note                    string                     `json:"note,omitempty"`
	MaxLicenseUsers         int                        `json:"max_license_users,omitempty"`
	Order                   int64                      `json:"order,omitempty"`
	OrderItem               int64                      `json:"order_item,omitempty"`
	DisabledUser            int                        `json:"disabled_user,omitempty"`
}
