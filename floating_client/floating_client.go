package floating_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
)

type FloatingClient struct {
	c *client.Client
}

func NewFloatingClient(config FloatingClientConfiguration) *FloatingClient {
	return &FloatingClient{
		c: client.NewClient(config, config.Verbose),
	}
}
