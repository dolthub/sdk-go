package management_client

import (
	management_models "gitlab.com/l3178/sdk/sdk-go/v1/management_client/models"
	management_request "gitlab.com/l3178/sdk/sdk-go/v1/management_client/models/request"
	management_response "gitlab.com/l3178/sdk/sdk-go/v1/management_client/models/response"
)

type DeviceApi ManagementClient

// DEVICES

func (api *DeviceApi) ListDevices(request management_request.SearchDevicesRequest) (resp management_response.SearchResult[management_models.Device], err error) {
	err = api.c.Get("devices", nil, request, &resp)
	return resp, err
}

func (api *DeviceApi) ShowDevice(id int64) (resp management_models.Device, err error) {
	err = api.c.Get("devices/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *DeviceApi) ResetDevice(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("devices/{id}/reset", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *DeviceApi) BlacklistDevice(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("devices/{id}/blacklist", id64ToParams(id), nil, &resp)
	return resp, err
}

// DEVICE VARIABLES

func (api *DeviceApi) ListDeviceVariables(request management_request.SearchDeviceVariablesRequest) (resp management_response.SearchResult[management_models.BackOfficeDeviceVariable], err error) {
	err = api.c.Get("device-variables", nil, request, &resp)
	return resp, err
}

func (api *DeviceApi) ShowDeviceVariable(id int64) (resp management_models.BackOfficeDeviceVariable, err error) {
	err = api.c.Get("device-variables/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *DeviceApi) CreateDeviceVariable(request management_request.CreateDeviceVariableRequest) (resp management_models.BackOfficeDeviceVariable, err error) {
	err = api.c.Post("device-variables", nil, request, &resp)
	return resp, err
}

func (api *DeviceApi) UpdateDeviceVariable(id int64, request management_request.UpdateDeviceVariableRequest) (resp management_models.BackOfficeDeviceVariable, err error) {
	err = api.c.Patch("device-variables/{id}", id64ToParams(id), request, &resp)
	return resp, err
}

func (api *DeviceApi) DeleteDeviceVariable(id int64) error {
	return api.c.Delete("device-variables/{id}", id64ToParams(id), nil)
}
