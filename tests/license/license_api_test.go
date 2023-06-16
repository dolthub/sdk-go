package license

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gitlab.com/l3178/sdk-go/core/auth"
	core_request "gitlab.com/l3178/sdk-go/core/models/request"
	"gitlab.com/l3178/sdk-go/license_client"
	"testing"
)

func TestLicenseApi(t *testing.T) {
	c := Setup()

	ctx := context.Background()

	licenseKey := "GZGL-N7BY-NPJK-2KTE"

	t.Run("Activate", func(t *testing.T) {
		resp := c.ActivateLicense(ctx, license_client.ActivationRequest{
			LicenseRequest: core_request.LicenseRequest{
				Auth: auth.FromKey(licenseKey),
			},
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Check", func(t *testing.T) {
		resp := c.CheckLicense(ctx, license_client.ActivationRequest{
			LicenseRequest: core_request.LicenseRequest{
				Auth: auth.FromKey(licenseKey),
			},
		})
		assert.NoError(t, resp.Error)
		assert.True(t, resp.Value.LicenseActive)
	})

	t.Run("Add consumption", func(t *testing.T) {
		resp := c.AddConsumption(ctx, license_client.ConsumptionRequest{
			LicenseRequest: core_request.LicenseRequest{
				Auth: auth.FromKey(licenseKey),
			},
			Consumptions: 1,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Add feature consumption", func(t *testing.T) {
		resp := c.AddFeatureConsumption(ctx, license_client.FeatureConsumptionRequest{
			LicenseRequest: core_request.LicenseRequest{
				Auth: auth.FromKey(licenseKey),
			},
			Feature:      "consumption-feature",
			Consumptions: 1,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Product details", func(t *testing.T) {
		resp := c.ProductDetails(ctx)
		assert.NoError(t, resp.Error)
	})

	t.Run("Deactivate", func(t *testing.T) {
		err := c.DeactivateLicense(ctx, license_client.LicenseRequest{
			Auth: auth.FromKey(licenseKey),
		})
		assert.NoError(t, err)
	})
}
