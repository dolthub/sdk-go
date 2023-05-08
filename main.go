package main

import (
	"gitlab.com/l3178/sdk-go/license_client"
)

func main() {
	cfg := license_client.NewLicenseClientConfiguration("apiKey", "sharedKey", "product")
	client := license_client.NewLicenseClient(cfg)

	license, err := client.AirgapActivation("licensePolicy", "confirmationCode", 1234)

	client.ActivateLicense(license_client.ActivationRequest{
		LicenseRequest: license_client.LicenseRequest{
			Auth: license_client.Auth{}.FromKey("G3KB-PBFB-4BAK-YGTV"),
		},
	})
}
