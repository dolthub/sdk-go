package floating_client

import (
	"context"
	"gitlab.com/l3178/sdk-go/core/client"
	"gitlab.com/l3178/sdk-go/core/models"
)

func (c *FloatingClient) Register(ctx context.Context, request FloatingRegistrationRequest) client.Response[core_models.LicenseResponse] {
	body, err := c.c.Post(ctx, "register", nil, request)
	return client.NewResponse[core_models.LicenseResponse](body, err)
}

func (c *FloatingClient) Unregister(ctx context.Context, request FloatingRegistrationRequest) error {
	_, err := c.c.Post(ctx, "unregister", nil, request)
	return err
}

func (c *FloatingClient) UnregisterAll(ctx context.Context) error {
	_, err := c.c.Post(ctx, "unregister/all", nil, nil)
	return err
}

func (c *FloatingClient) AddConsumption(ctx context.Context, request ProductRequest) error {
	_, err := c.c.Post(ctx, "add_consumption", nil, request)
	return err
}

func (c *FloatingClient) GetLicense(ctx context.Context, request ProductRequest) client.Response[core_models.LicenseResponse] {
	body, err := c.c.Get(ctx, "license", nil, request)
	return client.NewResponse[core_models.LicenseResponse](body, err)
}

func (c *FloatingClient) Ping(ctx context.Context) error {
	_, err := c.c.Get(ctx, "health", nil, nil)
	return err
}
