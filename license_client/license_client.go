package license_client

import (
	"licensespring/go-sdk/core/client"
	"licensespring/go-sdk/core/configuration"
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
