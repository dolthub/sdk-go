package license_client

import (
	"gitlab.com/l3178/sdk-go/core/models"
	"gitlab.com/l3178/sdk-go/core/models/request"
)

func (c *LicenseClient) ActivateLicense(request core_request.ActivationRequest) (resp core_models.LicenseResponse, err error) {
	err = c.c.Post("activate_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) DeactivateLicense(request core_request.LicenseRequest) error {
	err := c.c.Post("deactivate_license", nil, request, nil)
	return err
}

func (c *LicenseClient) CheckLicense(request core_request.ActivationRequest) (resp core_models.CheckResponse, err error) {
	err = c.c.Get("check_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddConsumption(request core_request.ConsumptionRequest) (resp core_models.ConsumptionResponse, err error) {
	err = c.c.Post("add_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddFeatureConsumption(request core_request.FeatureConsumptionRequest) (resp core_models.FeatureConsumptionResponse, err error) {
	err = c.c.Post("add_feature_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrialKey(request core_request.TrialLicenseRequest) (resp core_models.TrialKeyResponse, err error) {
	err = c.c.Get("trial_key", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) ProductDetails(request core_request.ProductDetailsRequest) (resp core_models.ProductDetails, err error) {
	err = c.c.Get("product_details", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrackDeviceVariables(request core_request.DeviceVariablesRequest) error {
	err := c.c.Post("track_device_variables", nil, request, nil)
	return err
}

func (c *LicenseClient) Versions(request core_request.LicenseRequest) (resp core_models.VersionsResponse, err error) {
	err = c.c.Get("versions", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) InstallationFile(request core_request.LicenseRequest) (resp core_models.InstallationFileResponse, err error) {
	err = c.c.Get("installation_file", nil, request, &resp)
	return resp, err
}
