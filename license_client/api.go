package license_client

import (
	"gitlab.com/l3178/sdk-go/core/models"
)

func (c *LicenseClient) ActivateLicense(request ActivationRequest) (resp core_models.LicenseResponse, err error) {
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

func (c *LicenseClient) CheckLicense(request ActivationRequest) (resp core_models.CheckResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("check_license", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddConsumption(request ConsumptionRequest) (resp core_models.ConsumptionResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("add_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) AddFeatureConsumption(request FeatureConsumptionRequest) (resp core_models.FeatureConsumptionResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Post("add_feature_consumption", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) TrialKey(request TrialLicenseRequest) (resp core_models.TrialKeyResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("trial_key", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) ProductDetails() (resp core_models.ProductDetails, err error) {
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

func (c *LicenseClient) Versions(request LicenseRequest) (resp core_models.VersionsResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("versions", nil, request, &resp)
	return resp, err
}

func (c *LicenseClient) InstallationFile(request LicenseRequest) (resp core_models.InstallationFileResponse, err error) {
	request.Product = c.ProductCode
	request.HardwareId = c.HardwareId
	err = c.c.Get("installation_file", nil, request, &resp)
	return resp, err
}
