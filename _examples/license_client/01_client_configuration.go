package license_client

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/l3178/sdk/sdk-go/v1/license_client"
)

func Configuration() {
	// set API keys from LicenseSpring
	apiKey := "<your_api_key>"
	sharedKey := "<your_shared_key>"
	// set the product code for your license
	productCode := "<your_product_code>"

	// initialize config
	config := license_client.NewLicenseClientConfiguration(apiKey, sharedKey, productCode)

	// we can further setup our config, or we can use the default values

	// setup logging
	config.LoggingLevel(logrus.DebugLevel)
	// disable multiple retries on failed requests
	config.RetryCount = 0
	// set request timeout
	config.RequestTimeout = 10
}

func InitializeClient() *license_client.LicenseClient {
	config := license_client.NewLicenseClientConfiguration("<your_api_key>", "<your_shared_key>", "<your_product_code>")

	// initialize client
	return license_client.NewLicenseClient(config)
}
