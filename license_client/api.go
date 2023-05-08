package license_client

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.com/l3178/sdk-go/core/client"
	. "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

func (c *LicenseClient) ActivateLicense(request ActivationRequest) client.Response[LicenseResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post("activate_license", nil, request)
	return client.NewResponse[LicenseResponse](body, err)
}

func (c *LicenseClient) DeactivateLicense(request LicenseRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post("deactivate_license", nil, request)
	return err
}

func (c *LicenseClient) CheckLicense(request ActivationRequest) client.Response[CheckResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get("check_license", nil, request)
	return client.NewResponse[CheckResponse](body, err)
}

func (c *LicenseClient) GenerateOfflineLicenseRequest(request ActivationRequest, requestType OfflineRequestType) client.Response[OfflineRequest] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	date := time.Now()
	val := OfflineRequest{
		ActivationRequest: request,
		Date:              date.Format(time.RFC1123),
		Request:           requestType,
		Signature:         c.OfflineLicenseSignature(date, request.LicenseRequest.Key),
		RequestId:         uuid.New().String(),
	}
	payload, err := json.Marshal(&val)
	return client.Response[OfflineRequest]{
		Payload: payload,
		Value:   val,
		Error:   err,
	}
}

func (c *LicenseClient) ActivateOffline(request OfflineRequest) client.Response[LicenseResponse] {
	if request.Request != OfflineActivationRequest {
		err := errors.New("Activate offline request needs to be of activate type")
		return client.ErrorResponse[LicenseResponse](err)
	}
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post("activate_offline", nil, request)
	return client.NewResponse[LicenseResponse](body, err)
}

func (c *LicenseClient) DeactivateOffline(request OfflineRequest) error {
	if request.Request != OfflineDeactivationRequest {
		return errors.New("Deactivate offline request needs to be of deactivate type")
	}
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post("deactivate_offline", nil, request)
	return err
}

func (c *LicenseClient) AddConsumption(request ConsumptionRequest) client.Response[ConsumptionResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post("add_consumption", nil, request)
	return client.NewResponse[ConsumptionResponse](body, err)
}

func (c *LicenseClient) AddFeatureConsumption(request FeatureConsumptionRequest) client.Response[FeatureConsumptionResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post("add_feature_consumption", nil, request)
	return client.NewResponse[FeatureConsumptionResponse](body, err)
}

func (c *LicenseClient) TrialKey(request TrialLicenseRequest) client.Response[TrialKeyResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get("trial_key", nil, request)
	return client.NewResponse[TrialKeyResponse](body, err)
}

func (c *LicenseClient) ProductDetails() client.Response[ProductDetails] {
	request := ProductDetailsRequest{Product: c.ProductCode}
	body, err := c.c.Get("product_details", nil, request)
	return client.NewResponse[ProductDetails](body, err)
}

func (c *LicenseClient) TrackDeviceVariables(request DeviceVariablesRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post("track_device_variables", nil, request)
	return err
}

func (c *LicenseClient) GetDeviceVariables(request LicenseRequest) client.Response[[]DeviceVariable] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get("get_device_variables", nil, request)
	return client.NewResponse[[]DeviceVariable](body, err)
}

func (c *LicenseClient) FloatingBorrow(request FloatingBorrowRequest) client.Response[FloatingBorrowResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post("floating/borrow", nil, request)
	return client.NewResponse[FloatingBorrowResponse](body, err)
}

func (c *LicenseClient) FloatingRelease(request LicenseRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post("floating/release", nil, request)
	return err
}

func (c *LicenseClient) ChangePassword(request ChangePasswordRequest) error {
	_, err := c.c.Post("change_password", nil, request)
	return err
}

func (c *LicenseClient) Versions(request LicenseRequest) client.Response[[]Version] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get("versions", nil, request)
	return client.NewResponse[[]Version](body, err)
}

func (c *LicenseClient) InstallationFile(request LicenseRequest) client.Response[InstallationFileResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get("installation_file", nil, request)
	return client.NewResponse[InstallationFileResponse](body, err)
}

func (c *LicenseClient) CustomerLicenseUsers(request CustomerLicenseUsersRequest) client.Response[CustomerLicenseUsersResponse] {
	request.Product = c.ProductCode
	body, err := c.c.Get("customer_license_users", nil, request)
	return client.NewResponse[CustomerLicenseUsersResponse](body, err)
}

func (c *LicenseClient) SSOUrl(request SSOUrlRequest) client.Response[SSOUrlResponse] {
	request.Product = c.ProductCode
	body, err := c.c.Get("sso_url", nil, request)
	return client.NewResponse[SSOUrlResponse](body, err)
}
