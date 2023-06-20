package main

import (
	"context"
	"fmt"
	"gitlab.com/l3178/sdk-go/core/auth"
	core_request "gitlab.com/l3178/sdk-go/core/models/request"
	"gitlab.com/l3178/sdk-go/license_client"
)

func main() {
	config := license_client.NewLicenseClientConfiguration("b46bb3c2-ddf7-459a-a198-c088bb5020ed",
		"GxehYAx_-HYvcxxKIBq42WQvCYGxfoy9I3cFxWRMLvs",
		"core")
	config.Verbose = true
	config.VerifySignature = true
	client := license_client.NewLicenseClient(config)
	resp := client.ActivateLicense(context.Background(), license_client.ActivationRequest{
		LicenseRequest: core_request.LicenseRequest{
			Auth: auth.FromKey("GVMX-U4JH-UBAK-6ECO"),
		},
	})
	fmt.Println(resp.Error)
	fmt.Println(string(resp.Payload))
}
