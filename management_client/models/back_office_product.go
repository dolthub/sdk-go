package management_models

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models"
	"time"
)

type BackOfficeProduct struct {
	Id                          int
	UpgradeFrom                 []int64
	ProductFeatures             []ProductFeature
	CustomFields                []BackOfficeCustomField
	BackOfficeInstallationFiles []BackOfficeInstallationFile
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
	ProductName                 string
	ShortCode                   string
	Active                      bool
	ValidDuration               string
	AllowTrial                  bool
	TrialDays                   int
	MaxActivations              int
	HardwareIdRequired          bool
	IsUpgrade                   bool
	SubscriptionDuration        string
	EnableMaintenancePeriod     bool
	MaintenanceDuration         string
	IsFloating                  bool
	IsFloatingCloud             bool
	FloatingUsers               int
	FloatingTimeout             int
	DefaultLicenseType          models.LicenseType
	IsUserLocked                bool
	IsNodeLocked                bool
	MaxConsumptions             int
	AuthorizationMethod         string
	PreventVm                   bool
	AllowOverages               bool
	MaxOverages                 int
	ResetConsumption            bool
	ConsumptionPeriod           string
	IsArchived                  bool
	Company                     int
}
