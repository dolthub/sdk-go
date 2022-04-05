package license_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
	"gitlab.com/l3178/sdk-go/core/configuration"
)

type LicenseClient struct {
	c *client.Client
}

func NewLicenseClient(apiKey, sharedKey string) *LicenseClient {
	config := configuration.NewClientConfig()
	config.SetKeys(apiKey, sharedKey)

	return &LicenseClient{
		c: client.NewClient(config),
	}
}

func NewLicenseClient2(apiKey, sharedKey string) *LicenseClient {
	config := configuration.NewClientConfig()
	config.SetKeys(apiKey, sharedKey)

	return &LicenseClient{
		c: client.NewClient(config),
	}
}
