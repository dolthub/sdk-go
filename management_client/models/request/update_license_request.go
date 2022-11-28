package management_request

type UpdateLicenseRequest struct {
	MaxActivations          int    `json:"max_activations,omitempty"`
	MaxConsumptions         int    `json:"max_consumptions,omitempty"`
	AllowOverages           bool   `json:"allow_overages,omitempty"`
	MaxOverages             int    `json:"max_overages,omitempty"`
	ResetConsumption        bool   `json:"reset_consumption,omitempty"`
	ConsumptionPeriod       string `json:"consumption_period,omitempty"`
	IsTrial                 bool   `json:"is_trial,omitempty"`
	TrialDays               int    `json:"trial_days,omitempty"`
	EnableMaintenancePeriod bool   `json:"enable_maintenance_period,omitempty"`
	MaintenanceDuration     string `json:"maintenance_duration,omitempty"`
	IsFloating              bool   `json:"is_floating,omitempty"`
	IsFloatingCloud         bool   `json:"is_floating_cloud,omitempty"`
	FloatingUsers           int    `json:"floating_users,omitempty"`
	ValidDuration           string `json:"valid_duration,omitempty"`
	ValidityPeriod          string `json:"validity_period,omitempty"`
	PreventVm               bool   `json:"prevent_vm,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	TotalConsumptions       int    `json:"total_consumptions,omitempty"`
}
