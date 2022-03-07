package management_request

type UpdateLicenseRequest struct {
	MaxActivations          int
	MaxConsumptions         int
	AllowOverages           bool
	MaxOverages             int
	ResetConsumption        bool
	ConsumptionPeriod       string
	IsTrial                 bool
	TrialDays               int
	EnableMaintenancePeriod bool
	MaintenanceDuration     string
	IsFloating              bool
	IsFloatingCloud         bool
	FloatingUsers           int
	ValidDuration           string
	ValidityPeriod          string
	PreventVm               bool
	Enabled                 bool
	TotalConsumptions       int
}
