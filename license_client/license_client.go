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

type AirgapClient struct {
	LicenseClient
	PublicKey string
}

func NewAirgapClient(config LicenseClientConfiguration, publicKey string) AirgapClient {
	lc := NewLicenseClient(config)
	return AirgapClient{
		LicenseClient: lc,
		PublicKey:     publicKey,
	}
}
