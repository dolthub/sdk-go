package management_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type DeviceApi ManagementClient

// DEVICES

func (api *DeviceApi) ListDevices(request management_request.SearchDevicesRequest) management_response.SearchResult[management_models.Device] {
	body, err := api.c.Get("devices", nil, request)
	return management_response.SearchResultFromJson[management_models.Device](body, err)
}

func (api *DeviceApi) ShowDevice(id int64) client.Response[management_models.Device] {
	body, err := api.c.Get("devices/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.Device](body, err)
}

func (api *DeviceApi) ResetDevice(id int64) client.Response[management_response.Response] {
	body, err := api.c.Post("devices/{id}/reset", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *DeviceApi) BlacklistDevice(id int64) client.Response[management_response.Response] {
	body, err := api.c.Post("devices/{id}/blacklist", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

// DEVICE VARIABLES

func (api *DeviceApi) ListDeviceVariables(request management_request.SearchDeviceVariablesRequest) management_response.SearchResult[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Get("device-variables", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) ShowDeviceVariable(id int64) client.Response[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Get("device-variables/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) CreateDeviceVariable(request management_request.CreateDeviceVariableRequest) client.Response[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Post("device-variables", nil, request)
	return client.NewResponse[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) UpdateDeviceVariable(id int64, request management_request.UpdateDeviceVariableRequest) client.Response[management_models.BackOfficeDeviceVariable] {
	body, err := api.c.Patch("device-variables/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.BackOfficeDeviceVariable](body, err)
}

func (api *DeviceApi) DeleteDeviceVariable(id int64) error {
	_, err := api.c.Delete("device-variables/{id}", id64ToParams(id))
	return err
}
