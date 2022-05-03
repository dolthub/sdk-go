package management_client

import (
	"encoding/json"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type LicenseApi ManagementClient

func (api *LicenseApi) SearchLicenses(request management_request.SearchLicensesRequest) (resp management_response.SearchResult[management_models.BackOfficeLicense], err error) {
	err = api.c.Get("licenses", nil, request, &resp)
	return resp, err
}
func (api *LicenseApi) GetLicense(id int64) (resp management_models.BackOfficeLicense, err error) {
	err = api.c.Get("licenses/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicenseApi) UpdateLicense(id int64, request management_request.UpdateLicenseRequest) (resp management_models.BackOfficeLicense, err error) {
	err = api.c.Patch("licenses/{id}", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicenseApi) DisableLicense(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("licenses/{id}/disable", id64ToParams(id), nil, &resp)
	return resp, err
}
func (api *LicenseApi) DisableAllLicenses(ids []int64) (resp management_response.Response, err error) {
	req, _ := json.Marshal(ids)
	err = api.c.Post("licenses/disable_bulk", nil, req, &resp)
	return resp, err
}
func (api *LicenseApi) ResetLicense(id int64) (resp management_response.Response, err error) {
	err = api.c.Post("licenses/{id}/reset", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *LicenseApi) PatchLicenseFeatures(id int64, request management_request.PatchLicenseFeaturesRequest) (resp management_models.BackOfficeLicense, err error) {
	err = api.c.Patch("licenses/{id}", id64ToParams(id), request, &resp)
	return resp, err
}

func (api *LicenseApi) AssignUser(id int64, request management_request.AssignUserToLicenseRequest) (resp management_response.LicenseAssignmentResponse, err error) {
	err = api.c.Post("licenses/{id}/assign_user", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicenseApi) AssignMultipleUsers(id int64, request management_request.AssignMultipleUsersRequest) (resp []management_response.MultipleAssignmentResponse, err error) {
	err = api.c.Post("licenses/{id}/assign_users", id64ToParams(id), request, &resp)
	return resp, err
}
func (api *LicenseApi) UnassignUser(id int64, request management_request.UnassignRequest) (resp management_response.Response, err error) {
	err = api.c.Post("licenses/{id}/unassign_use", id64ToParams(id), request, &resp)
	return resp, err
}

func (api *LicenseApi) SetUserPassword(id int, request management_request.SetPasswordRequest) (resp management_response.Response, err error) {
	err = api.c.Post("license-users/{id}/set_password", idToParams(id), request, &resp)
	return resp, err
}

func (api *LicenseApi) SearchLicenseCustomFields(request management_request.SearchLicenseCustomFieldsRequest) (resp management_response.SearchResult[management_models.BackOfficeLicenseCustomField], err error) {
	err = api.c.Get("license-custom-fields", nil, request, &resp)
	return resp, err
}
func (api *LicenseApi) GetLicenseCustomField(id int) (resp management_models.BackOfficeLicenseCustomField, err error) {
	err = api.c.Get("license-custom-fields/{id}", idToParams(id), nil, &resp)
	return resp, err
}
func (api *LicenseApi) CreateLicenseCustomField(request management_request.CreateLicenseCustomFieldRequest) (resp management_models.BackOfficeLicenseCustomField, err error) {
	err = api.c.Post("license-custom-fields", nil, request, &resp)
	return resp, err
}
func (api *LicenseApi) UpdateLicenseCustomField(id int, request management_request.UpdateCustomLicenseFieldRequest) (resp management_models.BackOfficeLicenseCustomField, err error) {
	err = api.c.Patch("license-custom-fields/{id}", idToParams(id), request, &resp)
	return resp, err
}
func (api *LicenseApi) DeleteLicenseCustomField(id int) error {
	return api.c.Delete("product-custom-fields/{id}", idToParams(id), nil)
}
