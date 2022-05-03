package floating_client

import (
	"gitlab.com/l3178/sdk-go/core/configuration"
	"time"
)

const (
	apiPrefix = "/api/v4/floating"
)

type FloatingClientConfiguration struct {
	*configuration.CoreConfiguration
}

func (config FloatingClientConfiguration) AuthHeader(date time.Time) string {
	return ""
}

func (config FloatingClientConfiguration) UrlPrefix() string {
	return config.CoreConfiguration.UrlPrefix()
}

func NewFloatingClientConfiguration(url string) FloatingClientConfiguration {
	config := configuration.NewClientConfig(url, apiPrefix)

	return FloatingClientConfiguration{
		CoreConfiguration: config,
	}
}
