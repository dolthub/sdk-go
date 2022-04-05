package management_client

import (
	"encoding/json"
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type LicensesApi struct {
	c *client.Client
}

func (api *LicensesApi) SearchLicenses(request management_request.SearchLicensesRequest) (resp management_response.SearchResult[management_models.BackOfficeLicense], err error) {
	err = api.c.Get("licenses", nil, request, &resp)
	return resp, err
}
func (api *LicensesApi) GetLicense(id int64) (resp management_models.BackOfficeLicense, err error) {
	err = api.c.Get("licenses/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicensesApi) UpdateLicense(id int64, request management_request.UpdateLicenseRequest) (resp management_models.BackOfficeLicense, err error) {
	err = api.c.Patch("licenses/{id}", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicensesApi) DisableLicense(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("licenses/{id}/disable", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicensesApi) DisableAllLicenses(ids []int64) (resp management_response.Response, err error) {
	req, _ := json.Marshal(ids)
	err = api.c.Post("licenses/disable_bulk", nil, req, &resp)
	return resp, err
}
func (api *LicensesApi) ResetLicense(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("licenses/{id}/reset", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *LicensesApi) PatchLicenseFeatures(id int64, request management_request.PatchLicenseFeaturesRequest) (resp management_models.BackOfficeLicense, err error) {
	err = api.c.Patch("licenses/{id}", id64ToParams(id), request, &resp)
	return resp, err
}

func (api *LicensesApi) AssignUser(id int64, request management_request.AssignUserToLicenseRequest) (resp management_response.LicenseAssignmentResponse, err error) {
	err = api.c.Post("licenses/{id}/assign_user", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicensesApi) AssignMultipleUsers(id int64, request management_request.AssignMultipleUsersRequest) (resp []management_response.MultipleAssignmentResponse, err error) {
	err = api.c.Post("licenses/{id}/assign_users", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicensesApi) UnassignUser(id int64, request management_request.UnassignRequest) (resp management_response.Response, err error) {
	err = api.c.Post("licenses/{id}/unassign_use", id64ToParams(id), request, &resp)
	return resp, err
}

func (api *LicensesApi) SetUserPassword(id int, request management_request.SetPasswordRequest) (resp management_response.Response, err error) {
	err = api.c.Post("license-users/{id}/set_password", idToParams(id), request, &resp)
	return resp, err
}

func (api *LicensesApi) SearchDevices(request management_request.SearchDevicesRequest) (resp management_response.SearchResult[management_models.Device], err error) {
	err = api.c.Get("devices", nil, request, &resp)
	return resp, err
}
func (api *LicensesApi) GetDevice(id int64) (resp management_models.Device, err error) {
	err = api.c.Get("devices/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicensesApi) ResetDevice(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("devices/{id}/reset", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicensesApi) BlacklistDevice(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("devices/{id}/blacklist", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *LicensesApi) SearchLicenseCustomFields(request management_request.SearchLicenseCustomFieldsRequest) (resp management_response.SearchResult[management_models.BackOfficeLicenseCustomField], err error) {
	err = api.c.Get("license-custom-fields", nil, request, &resp)
	return resp, err
}
func (api *LicensesApi) GetLicenseCustomField(id int) (resp management_models.BackOfficeLicenseCustomField, err error) {
	err = api.c.Get("license-custom-fields/{id}", idToParams(id), nil, &resp)
	return resp, err
}
func (api *LicensesApi) CreateLicenseCustomField(request management_request.CreateLicenseCustomFieldRequest) (resp management_models.BackOfficeLicenseCustomField, err error) {
	err = api.c.Post("license-custom-fields", nil, request, &resp)
	return resp, err
}
func (api *LicensesApi) UpdateLicenseCustomField(id int, request management_request.UpdateCustomLicenseFieldRequest) (resp management_models.BackOfficeLicenseCustomField, err error) {
	err = api.c.Patch("license-custom-fields/{id}", idToParams(id), request, &resp)
	return resp, err
}
func (api *LicensesApi) DeleteLicenseCustomField(id int) error {
	return api.c.Delete("product-custom-fields/{id}", idToParams(id), nil)
}

func (api *LicensesApi) SearchDeviceVariables(request management_request.SearchDeviceVariablesRequest) (resp management_response.SearchResult[management_models.BackOfficeDeviceVariable], err error) {
	err = api.c.Get("device-variables", nil, request, &resp)
	return resp, err
}
func (api *LicensesApi) GetDeviceVariable(id int64) (resp management_models.BackOfficeDeviceVariable, err error) {
	err = api.c.Get("device-variables/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicensesApi) CreateDeviceVariable(request management_request.CreateDeviceVariableRequest) (resp management_models.BackOfficeDeviceVariable, err error) {
	err = api.c.Post("device-variables", nil, request, &resp)
	return resp, err
}
func (api *LicensesApi) UpdateDeviceVariable(id int64, request management_request.UpdateDeviceVariableRequest) (resp management_models.BackOfficeDeviceVariable, err error) {
	err = api.c.Patch("device-variables/{id}", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicensesApi) DeleteDeviceVariable(id int64) error {
	return api.c.Delete("device-variables/{id}", id64ToParams(id), nil)
}
