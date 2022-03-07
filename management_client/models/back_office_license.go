package management_models

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models"
	"time"
)

type BackOfficeLicense struct {
	Id                      int64
	Status                  string
	FloatingInUseDevices    int
	Customer                BackOfficeCustomer
	ProductFeatures         []BackOfficeLicenseFeature
	LicenseUsers            []LicenseUser
	LicenseUser             LicenseUser
	OrderStoreId            string
	ActiveUpToDateDevices   int
	TotalActiveDevices      int
	Product                 BackOfficeLicenseProduct
	CreatedAt               time.Time
	UpdatedAt               time.Time
	LicenseKey              string
	Active                  bool
	Enabled                 bool
	TimeActivated           time.Time
	TimeDisabled            time.Time
	MaxActivtions           int
	TimesActivated          int
	ValidDuration           string
	ValidityPeriod          string
	LicenseType             models.LicenseType
	EnableMaintenancePeriod bool
	MaintenanceDuration     bool
	MaintenancePeriod       string
	IsTrial                 bool
	MaxConsumptions         int
	TotalConsumptions       int
	SubscriptionId          string
	LastCheck               time.Time
	PreventVm               bool
	AllowOverages           bool
	MaxOverages             int
	ResetConsumption        bool
	ConsumptionPeriod       string
	TrialDays               int
	IsFloating              bool
	IsFloatingCloud         bool
	FloatingUsers           int
	Note                    string
	MaxLicenseUsers         int
	Order                   int64
	OrderItem               int64
	DisabledUser            int
}
