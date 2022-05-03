package license_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/license_client"
)

func Consumption() {
	client := InitializeClient()

	// create requests
	licenseAuth := license_client.Auth{}.FromKey("<your_license>")
	licenseRequest := license_client.LicenseRequest{Auth: licenseAuth}

	consumptionRequest := license_client.ConsumptionRequest{
		LicenseRequest: licenseRequest,
		Consumptions:   5,
		MaxOverages:    3,
		AllowOverages:  true,
	}

	// activate license
	client.AddConsumption(consumptionRequest)
}

func FeatureConsumption() {
	client := InitializeClient()

	// create requests
	licenseAuth := license_client.Auth{}.FromKey("<your_license>")
	licenseRequest := license_client.LicenseRequest{Auth: licenseAuth}

	featureConsumptionRequest := license_client.FeatureConsumptionRequest{
		LicenseRequest: licenseRequest,
		Feature:        "feature",
		Consumptions:   5,
	}

	// activate license
	client.AddFeatureConsumption(featureConsumptionRequest)
}
