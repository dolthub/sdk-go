package management_client

import (
	"context"
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type DeviceApi ManagementClient

// DEVICES

func (api *DeviceApi) ListDevices(ctx context.Context, request management_request.SearchDevicesRequest) management_response.SearchResult[management_models.Device] {
	body, err := api.c.Get(ctx, "devices", nil, request)
	return management_response.SearchResultFromJson[management_models.Device](body, err)
}

func (api *DeviceApi) ShowDevice(ctx context.Context, id int64) client.Response[management_models.Device] {
	body, err := api.c.Get(ctx, "devices/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.Device](body, err)
}

func (api *DeviceApi) ResetDevice(ctx context.Context, id int64) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "devices/{id}/reset", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *DeviceApi) BlacklistDevice(ctx context.Context, id int64) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "devices/{id}/blacklist", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

// DEVICE VARIABLES

func (api *DeviceApi) ListDeviceVariables(ctx context.Context, request management_request.SearchDeviceVariablesRequest) management_response.SearchResult[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Get(ctx, "device-variables", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) ShowDeviceVariable(ctx context.Context, id int64) client.Response[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Get(ctx, "device-variables/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) CreateDeviceVariable(ctx context.Context, request management_request.CreateDeviceVariableRequest) client.Response[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Post(ctx, "device-variables", nil, request)
	return client.NewResponse[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) UpdateDeviceVariable(ctx context.Context, id int64, request management_request.UpdateDeviceVariableRequest) client.Response[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Patch(ctx, "device-variables/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) DeleteDeviceVariable(ctx context.Context, id int64) error {
	_, err := api.c.Delete(ctx, "device-variables/{id}", id64ToParams(id))
	return err
}
