package management_client

import (
	"fmt"
	"gitlab.com/l3178/sdk-go/core/configuration"
	"time"
)

const (
	baseUrl   = "https://saas.licensespring.com"
	apiPrefix = "/api/v1"
)

type ManagementClientConfiguration struct {
	*configuration.CoreConfiguration
	ManagementKey string
}

func (config ManagementClientConfiguration) AuthHeader(_ time.Time) string {
	return fmt.Sprintf("Api-Key %s", config.ManagementKey)
}

func (config ManagementClientConfiguration) UrlPrefix() string {
	return config.CoreConfiguration.UrlPrefix()
}

func NewManagementClientConfiguration(managementKey string) ManagementClientConfiguration {
	config := configuration.NewClientConfig(baseUrl, apiPrefix)

	return ManagementClientConfiguration{
		CoreConfiguration: config,
		ManagementKey:     managementKey,
	}
}
