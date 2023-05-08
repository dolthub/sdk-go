package license_client

import (
	"encoding/base64"
	"encoding/json"
	core_models "gitlab.com/l3178/sdk-go/core/models"
	core_request "gitlab.com/l3178/sdk-go/core/models/request"
	"time"
)

type ActivationRequest = core_request.ActivationRequest
type LicenseRequest = core_request.LicenseRequest
type ConsumptionRequest = core_request.ConsumptionRequest
type FeatureConsumptionRequest = core_request.FeatureConsumptionRequest
type TrialLicenseRequest = core_request.TrialLicenseRequest
type ProductDetailsRequest = core_request.ProductDetailsRequest
type DeviceVariablesRequest = core_request.DeviceVariablesRequest
type Auth = core_request.Auth

type OfflineRequest struct {
	ActivationRequest

	Date      string
	Request   OfflineRequestType
	Signature string
	RequestId string
}

func (req *OfflineRequest) Encode() (string, error) {
	jsn, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(jsn)
	return encoded, nil
}

type OfflineRequestType string

const (
	OfflineActivationRequest   = "activation"
	OfflineDeactivationRequest = "deactivation"
)

type DeviceVariable struct {
	Id        int       `json:"id,omitempty"`
	Value     string    `json:"value,omitempty"`
	Variable  string    `json:"variable,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	DeviceId  int64     `json:"device_id,omitempty"`
}

type FloatingBorrowRequest struct {
	LicenseRequest

	BorrowedUntil time.Time `json:"borrowed_until"`
}

type FloatingBorrowResponse struct {
	BorrowedUntil time.Time `json:"borrowed_until"`
	MaxBorrowTime int       `json:"max_borrow_time,omitempty"`
	DeviceId      int64     `json:"device_id,omitempty"`
	LicenseId     int64     `json:"license_id,omitempty"`
}

type ChangePasswordRequest struct {
	core_request.PasswordAuth
	NewPassword string `json:"new_password,omitempty"`
}

type CustomerLicenseUsersRequest struct {
	Product  string `json:"product,omitempty"`
	Customer string `json:"customer,omitempty"`
}

type CustomerLicenseUsersResponse struct {
	core_models.Customer
	Users []core_models.User `json:"users,omitempty"`
}

type SSOUrlRequest struct {
	Product             string `json:"product,omitempty"`
	CustomerAccountCode string `json:"customer_account_code,omitempty"`
}

type SSOUrlResponse struct {
	Url string `json:"url,omitempty"`
}
