package license_client

import (
	"gitlab.com/l3178/sdk-go/license_client"
)

func ProductDetails() {
	client := InitializeClient()

	// activate license
	client.ProductDetails()
}

func TrackDeviceVariables() {
	client := InitializeClient()

	// create requests
	licenseAuth := license_client.Auth{}.FromKey("<your_license>")
	licenseRequest := license_client.LicenseRequest{Auth: licenseAuth}

	deviceVariablesRequest := license_client.DeviceVariablesRequest{
		LicenseRequest: licenseRequest,
		Variables: map[string]string{
			"key": "value",
		},
	}

	// activate license
	client.TrackDeviceVariables(deviceVariablesRequest)
}

func InstallationFile() {
	client := InitializeClient()

	// create requests
	licenseAuth := license_client.Auth{}.FromKey("<your_license>")
	licenseRequest := license_client.LicenseRequest{Auth: licenseAuth}

	// activate license
	client.InstallationFile(licenseRequest)
}
