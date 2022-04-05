package floating_client

import (
	"gitlab.com/l3178/sdk-go/v1/core/client"
	"gitlab.com/l3178/sdk-go/v1/core/configuration"
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
