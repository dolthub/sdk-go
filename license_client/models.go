package license_client

import (
	"encoding/base64"
	"encoding/json"
	"gitlab.com/l3178/sdk-go/core/client"
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
	data := client.SnakeCaseEncode(req)
	jsn, err := json.Marshal(data)
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
	Id        int
	Value     string
	Variable  string
	CreatedAt time.Time
	DeviceId  int64
}

func (d *DeviceVariable) UnmarshalJSON(bytes []byte) error {
	type Alias struct {
		Id        int    `json:"id"`
		Value     string `json:"value"`
		Variable  string `json:"variable"`
		CreatedAt string `json:"created_at"`
		DeviceId  int64  `json:"device_id"`
	}
	var alias Alias

	err := json.Unmarshal(bytes, &alias)
	if err != nil {
		return err
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05Z", alias.CreatedAt)
	if err != nil {
		return err
	}

	d.Id = alias.Id
	d.Value = alias.Value
	d.Variable = alias.Variable
	d.CreatedAt = createdAt
	d.DeviceId = alias.DeviceId
	return nil
}

type FloatingBorrowRequest struct {
	LicenseRequest

	BorrowedUntil time.Time
}

func (r *FloatingBorrowRequest) MarshalJSON() ([]byte, error) {
	type Alias FloatingBorrowRequest
	return json.Marshal(&struct {
		*Alias
		BorrowedUntil string `json:"borrowed_until"`
	}{
		Alias:         (*Alias)(r),
		BorrowedUntil: r.BorrowedUntil.Format("2006-01-02T15:04:05Z"),
	})
}

type FloatingBorrowResponse struct {
	BorrowedUntil time.Time
	MaxBorrowTime int
	DeviceId      int64
	LicenseId     int64
}

func (r *FloatingBorrowResponse) UnmarshalJSON(bytes []byte) error {
	type Alias struct {
		BorrowedUntil string `json:"borrowed_until"`
		MaxBorrowTime int    `json:"max_borrow_time"`
		DeviceId      int64  `json:"device_id"`
		LicenseId     int64  `json:"license_id"`
	}
	var alias Alias

	err := json.Unmarshal(bytes, &alias)
	if err != nil {
		return err
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", alias.BorrowedUntil)
	if err != nil {
		return err
	}
	r.BorrowedUntil = ts
	r.MaxBorrowTime = alias.MaxBorrowTime
	r.DeviceId = alias.DeviceId
	r.LicenseId = alias.LicenseId
	return nil
}

type ChangePasswordRequest struct {
	core_request.PasswordAuth
	NewPassword string
}

type CustomerLicenseUsersRequest struct {
	Product  string
	Customer string
}

type CustomerLicenseUsersResponse struct {
	core_models.Customer
	Users []core_models.User
}

type SSOUrlRequest struct {
	Product             string
	CustomerAccountCode string
}

type SSOUrlResponse struct {
	Url string
}
