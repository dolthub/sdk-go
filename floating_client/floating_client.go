package floating_client

import (
	"licensespring/go-sdk/core/client"
	"licensespring/go-sdk/core/configuration"
)

type FloatingClient struct {
	c *client.Client
}

func NewFloatingClient(url string) *FloatingClient {
	config := configuration.NewClientConfig()
	config.UrlBase = url
	config.EndpointPrefix = "/api/v4/floating"

	return &FloatingClient{
		c: client.NewClient(config),
	}
}
