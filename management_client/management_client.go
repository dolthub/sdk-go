package management_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/core/client"
)

type ManagementClient struct {
	c *client.Client
}

func (api ManagementClient) CustomerApi() *CustomerApi {
	return &CustomerApi{api.c}
}

func (api ManagementClient) DeviceApi() *DeviceApi {
	return &DeviceApi{api.c}
}

func (api ManagementClient) LicenseApi() *LicenseApi {
	return &LicenseApi{api.c}
}

func (api ManagementClient) OrderApi() *OrderApi {
	return &OrderApi{api.c}
}

func (api ManagementClient) ProductApi() *ProductApi {
	return &ProductApi{api.c}
}

func NewManagementClient(config ManagementClientConfiguration) ManagementClient {
	return ManagementClient{
		c: client.NewClient(config),
	}
}
