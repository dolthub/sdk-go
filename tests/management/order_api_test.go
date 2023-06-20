package management

import (
	"context"
	"github.com/stretchr/testify/assert"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	"testing"
	"time"
)

const (
	orderID = 1593781720493279
)

func TestOrders(t *testing.T) {
	c := Setup()

	ctx := context.Background()

	t.Run("SearchOrders", func(t *testing.T) {
		resp := c.OrderApi().SearchOrders(ctx, management_request.SearchOrdersRequest{})
		assert.NoError(t, resp.Error)
	})

	t.Run("GetOrder", func(t *testing.T) {
		resp := c.OrderApi().GetOrder(ctx, orderID)
		assert.NoError(t, resp.Error)
	})

	t.Run("CreateOrder", func(t *testing.T) {
		resp := c.OrderApi().CreateOrder(ctx, management_models.WebhookOrder{
			Items: []management_models.OrderItem{
				{
					ProductCode: "code",
					Licenses: []management_models.OrderLicense{
						{
							Key: "key",
						},
					},
				},
			},
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("ExportToCSV", func(t *testing.T) {
		timeStart := time.Now().Add(-5 * time.Hour)
		timeEnd := time.Now()
		resp := c.OrderApi().ExportToCsv(ctx, timeStart, timeEnd)
		assert.NoError(t, resp.Error)
	})

	t.Run("ExportToCSVRange", func(t *testing.T) {
		var r management_models.Range
		r = r.Last30()
		resp := c.OrderApi().ExportToCsvRange(ctx, r)
		assert.NoError(t, resp.Error)
	})

	t.Run("GenerateLicenseKeys", func(t *testing.T) {
		resp := c.OrderApi().GenerateLicenseKeys(ctx, management_request.GenerateLicenseRequest{
			Product:  "product",
			Quantity: 1,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("AddManagerToOrder", func(t *testing.T) {
		resp := c.OrderApi().AddManagerToOrder(ctx, orderID, management_request.AddManagerToOrderRequest{
			Email:    "email",
			Password: "password",
		})
		assert.NoError(t, resp.Error)
	})
}
