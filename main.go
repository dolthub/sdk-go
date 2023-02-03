package main

import (
	"gitlab.com/l3178/sdk-go/license_client"
)

func main() {
	cfg := license_client.NewLicenseClientConfiguration("15676b6d-f241-42d4-aadd-29405329835c", "UANRVYyD1P2Gi8BiCgIs9pXZlD3uZgHUtKNTFuu5D20", "tvm4")
	cfg.BaseUrl = "https://api-dev.licensespring.com"
	c := license_client.NewLicenseClient(cfg)

	c.ActivateLicense(license_client.ActivationRequest{
		LicenseRequest: license_client.LicenseRequest{
			Auth: license_client.Auth{}.FromKey("G3KB-PBFB-4BAK-YGTV"),
		},
	})
}
