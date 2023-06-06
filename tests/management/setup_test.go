package management

import (
	"gitlab.com/l3178/sdk-go/management_client"
)

const (
	baseUrlTest   = "https://saas-staging.licensespring.com"
	managementKey = "QdAnoflC.OxnhzlBW5cozbhaTTtOFbyaBrzzJ1Kj3"
)

func Setup() management_client.ManagementClient {
	config := management_client.NewManagementClientConfiguration(managementKey)
	config.BaseUrl = baseUrlTest
	return management_client.NewManagementClient(config)
}
