package management

import (
	"context"
	"github.com/stretchr/testify/assert"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	"testing"
)

const (
	deviceID = 1590981949868340
)

func TestDevices(t *testing.T) {
	c := Setup()

	ctx := context.Background()

	t.Run("SearchDevices", func(t *testing.T) {
		resp := c.DeviceApi().ListDevices(ctx, management_request.SearchDevicesRequest{
			License: testLicenseID,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("ShowDevice", func(t *testing.T) {
		c.DeviceApi().ShowDevice(ctx, deviceID)
	})

	t.Run("ResetDevice", func(t *testing.T) {
		c.DeviceApi().ResetDevice(ctx, deviceID)
	})

	t.Run("BlacklistDevice", func(t *testing.T) {
		c.DeviceApi().BlacklistDevice(ctx, deviceID)
	})
}

func TestDeviceVariables(t *testing.T) {
	c := Setup()

	ctx := context.Background()

	t.Run("CreateDeviceVariable", func(t *testing.T) {
		resp := c.DeviceApi().CreateDeviceVariable(ctx, management_request.CreateDeviceVariableRequest{
			Variable: "test",
			Value:    "value",
			Device:   deviceID,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("UpdateDeviceVariable", func(t *testing.T) {
		resp := c.DeviceApi().UpdateDeviceVariable(ctx, deviceID, management_request.UpdateDeviceVariableRequest{
			Variable: "test",
			Value:    "value",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("DeleteDeviceVariable", func(t *testing.T) {
		err := c.DeviceApi().DeleteDeviceVariable(ctx, deviceID)
		assert.NoError(t, err)
	})

	t.Run("SearchDeviceVariable", func(t *testing.T) {
		resp := c.DeviceApi().ListDeviceVariables(ctx, management_request.SearchDeviceVariablesRequest{
			Device: deviceID,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("ShowDeviceVariable", func(t *testing.T) {
		resp := c.DeviceApi().ShowDeviceVariable(ctx, deviceID)
		assert.NoError(t, resp.Error)
	})
}
