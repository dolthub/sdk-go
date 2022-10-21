package core_models

import (
	"time"
)

type LicenseResponse struct {
	LicenseSignature  string
	ValidityPeriod    time.Time
	MaintenancePeriod time.Time

	LicenseType LicenseType
	LicenseKey  string

	IsTrial bool

	ProductFeatures []ProductFeature

	CustomFields []CustomFields

	MaxActivations int
	TimesActivated int

	Customer Customer

	MaxConsumptions   int
	TotalConsumptions int
	AllowOverages     bool
	MaxOverages       int
	ResetConsumption  bool
	ConsumptionPeriod string

	PreventVm            bool
	IsFloatingCloud      bool
	IsFloating           bool
	FloatingTimeout      int
	FloatingUsers        int
	FloatingInUseDevices int
	FloatingInUse        int

	User User

	ProductDetails ProductDetails
}

type ProductFeature struct {
	Code              string
	Name              string
	FeatureType       string
	ExpiryDate        string
	MaxConsumption    int
	TotalConsumptions int
	AllowOverages     bool
	MaxOverages       int
}

type CustomFields struct {
	Name  string
	Value string
}

type Customer struct {
	Email       string
	FirstName   string
	LastName    string
	CompanyName string
	Phone       string
	Reference   string
	Address     string
	Postcode    string
	City        string
	Country     string
	State       string
}

type User struct {
	Id                int
	Email             string
	IsActive          bool
	FirstName         string
	LastName          string
	PhoneNumber       string
	IsInitialPassword bool
	InitialPassword   string
	LicenseId         int64
	OrderStoreId      string
	OrderId           int64
}

type ProductDetails struct {
	ProductId           int64
	ProductName         string
	ShortCode           string
	TrialDays           int
	AuthorizationMethod string
	AllowOverages       bool
	MaxOverages         int
}
