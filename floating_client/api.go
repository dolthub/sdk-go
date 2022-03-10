package floating_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/core/models"
	floating_models "gitlab.com/l3178/sdk/sdk-go/v1/floating_client/models"
)

func (c *FloatingClient) Register(request floating_models.FloatingRegistrationRequest) (resp core_models.LicenseResponse, err error) {
	err = c.c.Post("register", nil, request, &resp)
	return resp, err
}

func (c *FloatingClient) Unregister(request floating_models.FloatingRegistrationRequest) error {
	return c.c.Post("unregister", nil, request, nil)
}

func (c *FloatingClient) UnregisterAll() error {
	return c.c.Post("unregister/all", nil, nil, nil)
}

func (c *FloatingClient) AddConsumption(request floating_models.ProductRequest) error {
	return c.c.Post("add_consumption", nil, request, nil)
}

func (c *FloatingClient) GetLicense(request floating_models.ProductRequest) (resp core_models.LicenseResponse, err error) {
	err = c.c.Get("license", nil, request, &resp)
	return resp, err
}

func (c *FloatingClient) Ping() error {
	return c.c.Get("health", nil, nil, nil)
}
