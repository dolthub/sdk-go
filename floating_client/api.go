package floating_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/core/models"
)

func (c *FloatingClient) Register(request FloatingRegistrationRequest) (resp core_models.LicenseResponse, err error) {
	err = c.c.Post("register", nil, request, &resp)
	return resp, err
}

func (c *FloatingClient) Unregister(request FloatingRegistrationRequest) error {
	return c.c.Post("unregister", nil, request, nil)
}

func (c *FloatingClient) UnregisterAll() error {
	return c.c.Post("unregister/all", nil, nil, nil)
}

func (c *FloatingClient) AddConsumption(request ProductRequest) error {
	return c.c.Post("add_consumption", nil, request, nil)
}

func (c *FloatingClient) GetLicense(request ProductRequest) (resp core_models.LicenseResponse, err error) {
	err = c.c.Get("license", nil, request, &resp)
	return resp, err
}

func (c *FloatingClient) Ping() error {
	return c.c.Get("health", nil, nil, nil)
}
