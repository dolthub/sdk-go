package license_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
)

type LicenseClient struct {
	c *client.Client
	LicenseClientConfiguration
}

func NewLicenseClient(config LicenseClientConfiguration) LicenseClient {
	return LicenseClient{
		c:                          client.NewClient(config, config.Verbose),
		LicenseClientConfiguration: config,
	}
}
