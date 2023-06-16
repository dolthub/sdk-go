package license

import (
	"gitlab.com/l3178/sdk-go/license_client"
)

const (
	apiKey     = "19cd1ea0-61bf-456e-9837-51330d803c75"
	sharedKey  = "melShB9ByVAbvbeW3Tk4AWcxPmws42YL6g0FiJH7U_o"
	product    = "test"
	baseUrl    = "https://api-staging.licensespring.com"
	hardwareID = "SDK-TEST"
)

func Setup() license_client.LicenseClient {
	config := license_client.NewLicenseClientConfiguration(apiKey, sharedKey, product)
	config.BaseUrl = baseUrl
	config.HardwareId = hardwareID
	return license_client.NewLicenseClient(config)
}
