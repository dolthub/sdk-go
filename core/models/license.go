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
	Code              string `json:"code,omitempty"`
	Name              string `json:"name,omitempty"`
	FeatureType       string `json:"feature_type,omitempty"`
	ExpiryDate        string `json:"expiry_date,omitempty"`
	MaxConsumption    int    `json:"max_consumption,omitempty"`
	TotalConsumptions int    `json:"total_consumptions,omitempty"`
	AllowOverages     bool   `json:"allow_overages,omitempty"`
	MaxOverages       int    `json:"max_overages,omitempty"`
	ResetConsumption  bool   `json:"reset_consumption"`
	ConsumptionPeriod string `json:"consumption_period"`
}

type CustomFields struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Customer struct {
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Reference   string `json:"reference,omitempty"`
	Address     string `json:"address,omitempty"`
	Postcode    string `json:"postcode,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	State       string `json:"state,omitempty"`
}

type User struct {
	Id                int    `json:"id,omitempty"`
	Email             string `json:"email,omitempty"`
	IsActive          bool   `json:"is_active,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	PhoneNumber       string `json:"phone_number,omitempty"`
	IsInitialPassword bool   `json:"is_initial_password,omitempty"`
	InitialPassword   string `json:"initial_password,omitempty"`
	LicenseId         int64  `json:"license_id,omitempty"`
	OrderStoreId      string `json:"order_store_id,omitempty"`
	OrderId           int64  `json:"order_id,omitempty"`
}

type ProductDetails struct {
	ProductId           int64  `json:"product_id,omitempty"`
	ProductName         string `json:"product_name,omitempty"`
	ShortCode           string `json:"short_code,omitempty"`
	TrialDays           int    `json:"trial_days,omitempty"`
	AuthorizationMethod string `json:"authorization_method,omitempty"`
	AllowOverages       bool   `json:"allow_overages,omitempty"`
	MaxOverages         int    `json:"max_overages,omitempty"`
}
