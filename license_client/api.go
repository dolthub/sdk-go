package license_client

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models/request"
	"gitlab.com/l216/sdk/sdk-go/v1/core/models/response"
)

func (c *LicenseClient) ActivateLicense(request request.ActivationRequest) (resp response.LicenseResponse, err error) {
	err = c.c.Post("activate_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) DeactivateLicense(request request.LicenseRequest) error {
	err := c.c.Post("deactivate_license", nil, request, nil)
	return err
}

func (c *LicenseClient) CheckLicense(request request.ActivationRequest) (resp response.CheckResponse, err error) {
	err = c.c.Get("check_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddConsumption(request request.ConsumptionRequest) (resp response.ConsumptionResponse, err error) {
	err = c.c.Post("add_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddFeatureConsumption(request request.FeatureConsumptionRequest) (resp response.FeatureConsumptionResponse, err error) {
	err = c.c.Post("add_feature_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrialKey(request request.TrialLicenseRequest) (resp response.TrialKeyResponse, err error) {
	err = c.c.Get("trial_key", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) ProductDetails(request request.ProductDetailsRequest) (resp response.ProductDetails, err error) {
	err = c.c.Get("product_details", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrackDeviceVariables(request request.DeviceVariablesRequest) error {
	err := c.c.Post("track_device_variables", nil, request, nil)
	return err
}

func (c *LicenseClient) Versions(request request.LicenseRequest) (resp response.VersionsResponse, err error) {
	err = c.c.Get("versions", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) InstallationFile(request request.LicenseRequest) (resp response.InstallationFileResponse, err error) {
	err = c.c.Get("installation_file", nil, request, &resp)
	return resp, err
}
