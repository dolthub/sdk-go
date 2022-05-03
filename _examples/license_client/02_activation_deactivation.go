package license_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/license_client"
)

func ActivateAndDeactivateLicense() {
	client := InitializeClient()

	// create requests
	licenseAuth := license_client.Auth{}.FromKey("<your_license>")
	licenseRequest := license_client.LicenseRequest{Auth: licenseAuth}
	activationRequest := licenseRequest.ToActivationRequest()

	// activate license
	client.ActivateLicense(activationRequest)

	// deactivate license
	client.DeactivateLicense(licenseRequest)
}

func ActivateWithAdditionalData() {
	client := InitializeClient()

	// create requests
	licenseAuth := license_client.Auth{}.FromKey("<your_license>")
	licenseRequest := license_client.LicenseRequest{Auth: licenseAuth}

	// additional data for activation request
	activationRequest := license_client.ActivationRequest{
		LicenseRequest: licenseRequest,
		Hostname:       "my_pc",
		OsVersion:      "Windows 10",
		IP:             "123.456.789.000",
	}

	// activate license
	client.ActivateLicense(activationRequest)
}
