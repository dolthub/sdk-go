package license_client

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	. "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

func (c *LicenseClient) ActivateLicense(request ActivationRequest) (resp LicenseResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("activate_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) DeactivateLicense(request LicenseRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err := c.c.Post("deactivate_license", nil, request, nil)
	return err
}

func (c *LicenseClient) CheckLicense(request ActivationRequest) (resp CheckResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("check_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) GenerateOfflineLicenseRequest(request ActivationRequest, requestType OfflineRequestType) (resp OfflineRequest, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	date := time.Now()
	resp = OfflineRequest{
		ActivationRequest: request,
		Date:              date.Format(time.RFC1123),
		Request:           requestType,
		Signature:         c.OfflineLicenseSignature(date, request.LicenseRequest.Key),
		RequestId:         uuid.New().String(),
	}
	return resp, err
}

func (c *LicenseClient) ActivateOffline(request OfflineRequest) (resp LicenseResponse, err error) {
	if request.Request != OfflineActivationRequest {
		return LicenseResponse{}, errors.New("Activate offline request needs to be of activate type")
	}
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("activate_offline", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) DeactivateOffline(request OfflineRequest) error {
	if request.Request != OfflineDeactivationRequest {
		return errors.New("Deactivate offline request needs to be of deactivate type")
	}
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err := c.c.Post("deactivate_offline", nil, request, nil)
	return err
}

func (c *LicenseClient) AddConsumption(request ConsumptionRequest) (resp ConsumptionResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("add_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddFeatureConsumption(request FeatureConsumptionRequest) (resp FeatureConsumptionResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("add_feature_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrialKey(request TrialLicenseRequest) (resp TrialKeyResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("trial_key", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) ProductDetails() (resp ProductDetails, err error) {
	request := ProductDetailsRequest{Product: c.ProductCode}
	err = c.c.Get("product_details", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrackDeviceVariables(request DeviceVariablesRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err := c.c.Post("track_device_variables", nil, request, nil)
	return err
}

func (c *LicenseClient) GetDeviceVariables(request LicenseRequest) (resp []DeviceVariable, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("get_device_variables", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) FloatingBorrow(request FloatingBorrowRequest) (resp FloatingBorrowResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("floating/borrow", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) FloatingRelease(request LicenseRequest) error {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err := c.c.Post("floating/release", nil, request, nil)
	return err
}

func (c *LicenseClient) ChangePassword(request ChangePasswordRequest) error {
	err := c.c.Post("change_password", nil, request, nil)
	return err
}

func (c *LicenseClient) Versions(request LicenseRequest) (resp []Version, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("versions", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) InstallationFile(request LicenseRequest) (resp InstallationFileResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("installation_file", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) CustomerLicenseUsers(request CustomerLicenseUsersRequest) (resp CustomerLicenseUsersResponse, err error) {
	request.Product = c.ProductCode
	err = c.c.Get("customer_license_users", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) SSOUrl(request SSOUrlRequest) (resp SSOUrlResponse, err error) {
	request.Product = c.ProductCode
	err = c.c.Get("sso_url", nil, request, &resp)
	return resp, err
}
