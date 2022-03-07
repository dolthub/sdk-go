package response

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models"
	"time"
)

type LicenseResponse struct {
	LicenseSignature  string    `json:"license_signature"`
	ValidityPeriod    time.Time `json:"validity_period"`
	MaintenancePeriod time.Time `json:"maintenance_period"`

	LicenseType models.LicenseType `json:"license_type"`
	LicenseKey  string             `json:"license_key"`

	IsTrial bool `json:"is_trial"`

	ProductFeatures []ProductFeature `json:"product_features"`

	CustomFields []CustomFields `json:"custom_fields"`

	MaxActivations int `json:"max_activations"`
	TimesActivated int `json:"times_activated"`

	Customer Customer `json:"customer"`

	MaxConsumptions   int    `json:"max_consumptions"`
	TotalConsumptions int    `json:"total_consumptions"`
	AllowOverages     bool   `json:"allow_overages"`
	MaxOverages       int    `json:"max_overages"`
	ResetConsumption  bool   `json:"reset_consumption"`
	ConsumptionPeriod string `json:"consumption_period"`

	PreventVm            bool `json:"prevent_vm"`
	IsFloatingCloud      bool `json:"is_floating_cloud"`
	IsFloating           bool `json:"is_floating"`
	FloatingTimeout      int  `json:"floating_timeout"`
	FloatingUsers        int  `json:"floating_users"`
	FloatingInUseDevices int  `json:"floating_in_use_devices"`
	FloatingInUse        int  `json:"floating_in_use"`

	User User `json:"user"`

	ProductDetails ProductDetails `json:"product_details"`
}

type ProductFeature struct {
	Code              string `json:"code"`
	Name              string `json:"name"`
	FeatureType       string `json:"feature_type"`
	ExpiryDate        string `json:"expiry_date"`
	MaxConsumption    int    `json:"max_consumption"`
	TotalConsumptions int    `json:"total_consumptions"`
	AllowOverages     bool   `json:"allow_overages"`
	MaxOverages       int    `json:"max_overages"`
}

type CustomFields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Customer struct {
	Email       string `json:"email"`
	CompanyName string `json:"company_name"`
	Reference   string `json:"reference"`
	Phone       string `json:"phone"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	IsInitialPassword bool   `json:"is_initial_password"`
}

type ProductDetails struct {
	ProductId           int64  `json:"product_id"`
	ProductName         string `json:"product_name"`
	ShortCode           string `json:"short_code"`
	TrialDays           int    `json:"trial_days"`
	AuthorizationMethod string `json:"authorization_method"`
	AllowOverages       bool   `json:"allow_overages"`
	MaxOverages         int    `json:"max_overages"`
}
