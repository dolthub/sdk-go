package management_client

const (
	baseUrlTest   = "https://saas-staging.licensespring.com"
	managementKey = "QdAnoflC.OxnhzlBW5cozbhaTTtOFbyaBrzzJ1Kj3"
)

func Setup() ManagementClient {
	config := NewManagementClientConfiguration(managementKey)
	config.BaseUrl = baseUrlTest
	return NewManagementClient(config)
}
