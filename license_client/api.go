package license_client

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.com/l3178/sdk-go/core/client"
	. "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

func (c *LicenseClient) ActivateLicense(ctx context.Context, request ActivationRequest) client.Response[LicenseResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post(ctx, "activate_license", nil, request)
	return client.NewResponse[LicenseResponse](body, err)
}

func (c *LicenseClient) DeactivateLicense(ctx context.Context, request LicenseRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post(ctx, "deactivate_license", nil, request)
	return err
}

func (c *LicenseClient) CheckLicense(ctx context.Context, request ActivationRequest) client.Response[CheckResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get(ctx, "check_license", nil, request)
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

func (c *LicenseClient) ActivateOffline(ctx context.Context, request OfflineRequest) client.Response[LicenseResponse] {
	if request.Request != OfflineActivationRequest {
		err := errors.New("Activate offline request needs to be of activate type")
		return client.ErrorResponse[LicenseResponse](err)
	}
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post(ctx, "activate_offline", nil, request)
	return client.NewResponse[LicenseResponse](body, err)
}

func (c *LicenseClient) DeactivateOffline(ctx context.Context, request OfflineRequest) error {
	if request.Request != OfflineDeactivationRequest {
		return errors.New("Deactivate offline request needs to be of deactivate type")
	}
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post(ctx, "deactivate_offline", nil, request)
	return err
}

func (c *LicenseClient) AddConsumption(ctx context.Context, request ConsumptionRequest) client.Response[ConsumptionResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post(ctx, "add_consumption", nil, request)
	return client.NewResponse[ConsumptionResponse](body, err)
}

func (c *LicenseClient) AddFeatureConsumption(ctx context.Context, request FeatureConsumptionRequest) client.Response[FeatureConsumptionResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post(ctx, "add_feature_consumption", nil, request)
	return client.NewResponse[FeatureConsumptionResponse](body, err)
}

func (c *LicenseClient) TrialKey(ctx context.Context, request TrialLicenseRequest) client.Response[TrialKeyResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get(ctx, "trial_key", nil, request)
	return client.NewResponse[TrialKeyResponse](body, err)
}

func (c *LicenseClient) ProductDetails(ctx context.Context) client.Response[ProductDetails] {
	request := ProductDetailsRequest{Product: c.ProductCode}
	body, err := c.c.Get(ctx, "product_details", nil, request)
	return client.NewResponse[ProductDetails](body, err)
}

func (c *LicenseClient) TrackDeviceVariables(ctx context.Context, request DeviceVariablesRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post(ctx, "track_device_variables", nil, request)
	return err
}

func (c *LicenseClient) GetDeviceVariables(ctx context.Context, request LicenseRequest) client.Response[[]DeviceVariable] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get(ctx, "get_device_variables", nil, request)
	return client.NewResponse[[]DeviceVariable](body, err)
}

func (c *LicenseClient) FloatingBorrow(ctx context.Context, request FloatingBorrowRequest) client.Response[FloatingBorrowResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Post(ctx, "floating/borrow", nil, request)
	return client.NewResponse[FloatingBorrowResponse](body, err)
}

func (c *LicenseClient) FloatingRelease(ctx context.Context, request LicenseRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	_, err := c.c.Post(ctx, "floating/release", nil, request)
	return err
}

func (c *LicenseClient) ChangePassword(ctx context.Context, request ChangePasswordRequest) error {
	_, err := c.c.Post(ctx, "change_password", nil, request)
	return err
}

func (c *LicenseClient) Versions(ctx context.Context, request LicenseRequest) client.Response[[]Version] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get(ctx, "versions", nil, request)
	return client.NewResponse[[]Version](body, err)
}

func (c *LicenseClient) InstallationFile(ctx context.Context, request LicenseRequest) client.Response[InstallationFileResponse] {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	body, err := c.c.Get(ctx, "installation_file", nil, request)
	return client.NewResponse[InstallationFileResponse](body, err)
}

func (c *LicenseClient) CustomerLicenseUsers(ctx context.Context, request CustomerLicenseUsersRequest) client.Response[CustomerLicenseUsersResponse] {
	request.Product = c.ProductCode
	body, err := c.c.Get(ctx, "customer_license_users", nil, request)
	return client.NewResponse[CustomerLicenseUsersResponse](body, err)
}

func (c *LicenseClient) SSOUrl(ctx context.Context, request SSOUrlRequest) client.Response[SSOUrlResponse] {
	request.Product = c.ProductCode
	body, err := c.c.Get(ctx, "sso_url", nil, request)
	return client.NewResponse[SSOUrlResponse](body, err)
}
