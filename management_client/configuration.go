package management_client

import (
	"fmt"
	"gitlab.com/l3178/sdk-go/core/configuration"
)

type ManagementConfiguration struct {
	*configuration.CoreConfiguration
	ManagementKey string
}

func NewManagementConfiguration(managementKey string) ManagementConfiguration {
	coreConfig := configuration.NewClientConfig()
	coreConfig.AdditionalHeaders["Authorization"] = fmt.Sprintf("Api-Key %s", managementKey)
	return ManagementConfiguration{
		CoreConfiguration: coreConfig,
		ManagementKey:     managementKey,
	}
}
