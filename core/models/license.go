package core_models

import (
	"time"
)

type LicenseResponse struct {
	Id                int64     `json:"id,omitempty"`
	LicenseSignature  string    `json:"license_signature,omitempty"`
	ValidityPeriod    time.Time `json:"validity_period"`
	MaintenancePeriod time.Time `json:"maintenance_period"`

	LicenseType LicenseType `json:"license_type,omitempty"`
	LicenseKey  string      `json:"license_key,omitempty"`

	IsTrial bool `json:"is_trial,omitempty"`

	ProductFeatures []ProductFeature `json:"product_features,omitempty"`

	CustomFields []CustomFields `json:"custom_fields,omitempty"`

	MaxActivations int `json:"max_activations,omitempty"`
	TimesActivated int `json:"times_activated,omitempty"`

	Customer Customer `json:"customer"`

	MaxConsumptions   int    `json:"max_consumptions,omitempty"`
	TotalConsumptions int    `json:"total_consumptions,omitempty"`
	AllowOverages     bool   `json:"allow_overages,omitempty"`
	MaxOverages       int    `json:"max_overages,omitempty"`
	ResetConsumption  bool   `json:"reset_consumption,omitempty"`
	ConsumptionPeriod string `json:"consumption_period,omitempty"`

	PreventVm            bool `json:"prevent_vm,omitempty"`
	IsFloatingCloud      bool `json:"is_floating_cloud,omitempty"`
	IsFloating           bool `json:"is_floating,omitempty"`
	FloatingTimeout      int  `json:"floating_timeout,omitempty"`
	FloatingUsers        int  `json:"floating_users,omitempty"`
	FloatingInUseDevices int  `json:"floating_in_use_devices,omitempty"`
	FloatingInUse        bool `json:"floating_in_use,omitempty"`

	User User `json:"user"`

	ProductDetails ProductDetails `json:"product_details"`
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
