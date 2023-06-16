package management_client

import (
	"context"
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type LicenseApi ManagementClient

// LICENSE

func (api *LicenseApi) SearchLicenses(ctx context.Context, request management_request.SearchLicensesRequest) management_response.SearchResult[management_models.BackOfficeLicense] {
	body, err := api.c.Get(ctx, "licenses", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeLicense](body, err)
}

func (api *LicenseApi) GetLicense(ctx context.Context, id int64) client.Response[management_models.BackOfficeLicense] {
	body, err := api.c.Get(ctx, "licenses/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeLicense](body, err)
}

func (api *LicenseApi) UpdateLicense(ctx context.Context, id int64, request management_request.UpdateLicenseRequest) client.Response[management_models.BackOfficeLicense] {
	body, err := api.c.Patch(ctx, "licenses/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.BackOfficeLicense](body, err)
}

func (api *LicenseApi) DisableLicense(ctx context.Context, id int64) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "licenses/{id}/disable", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *LicenseApi) DisableAllLicenses(ctx context.Context, ids []int64) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "licenses/disable_bulk", nil, ids)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *LicenseApi) EnableLicense(ctx context.Context, id int64) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "licenses/{id}/enable", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *LicenseApi) ResetLicense(ctx context.Context, id int64) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "licenses/{id}/reset", id64ToParams(id), nil)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *LicenseApi) PatchLicenseFeatures(ctx context.Context, id int64, request management_request.PatchLicenseFeaturesRequest) client.Response[management_models.BackOfficeLicense] {
	body, err := api.c.Patch(ctx, "licenses/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.BackOfficeLicense](body, err)
}

// USER

func (api *LicenseApi) ListLicenseUsers(ctx context.Context, request management_request.SearchUsersRequest) management_response.SearchResult[management_models.LicenseUser] {
	body, err := api.c.Get(ctx, "license-users", nil, request)
	return management_response.SearchResultFromJson[management_models.LicenseUser](body, err)
}

func (api *LicenseApi) AssignUser(ctx context.Context, id int64, request management_request.AssignUserToLicenseRequest) client.Response[management_response.LicenseAssignmentResponse] {
	body, err := api.c.Post(ctx, "licenses/{id}/assign_user", id64ToParams(id), request)
	return client.NewResponse[management_response.LicenseAssignmentResponse](body, err)
}

func (api *LicenseApi) AssignMultipleUsers(ctx context.Context, id int64, request management_request.AssignMultipleUsersRequest) client.Response[[]management_response.MultipleAssignmentResponse] {
	body, err := api.c.Post(ctx, "licenses/{id}/assign_users", id64ToParams(id), request)
	return client.NewResponse[[]management_response.MultipleAssignmentResponse](body, err)
}

func (api *LicenseApi) UnassignUser(ctx context.Context, id int64, request management_request.UnassignRequest) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "licenses/{id}/unassign_user", id64ToParams(id), request)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *LicenseApi) SetUserPassword(ctx context.Context, id int, request management_request.SetPasswordRequest) client.Response[management_response.Response] {
	body, err := api.c.Post(ctx, "license-users/{id}/set_password", idToParams(id), request)
	return client.NewResponse[management_response.Response](body, err)
}

// CUSTOM FIELDS

func (api *LicenseApi) SearchLicenseCustomFields(ctx context.Context, request management_request.SearchLicenseCustomFieldsRequest) management_response.SearchResult[management_models.BackOfficeLicenseCustomField] {
	body, err := api.c.Get(ctx, "license-custom-fields", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeLicenseCustomField](body, err)
}

func (api *LicenseApi) GetLicenseCustomField(ctx context.Context, id int) client.Response[management_models.BackOfficeLicenseCustomField] {
	body, err := api.c.Get(ctx, "license-custom-fields/{id}", idToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeLicenseCustomField](body, err)
}

func (api *LicenseApi) CreateLicenseCustomField(ctx context.Context, request management_request.CreateLicenseCustomFieldRequest) client.Response[management_models.BackOfficeLicenseCustomField] {
	body, err := api.c.Post(ctx, "license-custom-fields", nil, request)
	return client.NewResponse[management_models.BackOfficeLicenseCustomField](body, err)
}

func (api *LicenseApi) UpdateLicenseCustomField(ctx context.Context, id int, request management_request.UpdateCustomLicenseFieldRequest) client.Response[management_models.BackOfficeLicenseCustomField] {
	body, err := api.c.Patch(ctx, "license-custom-fields/{id}", idToParams(id), request)
	return client.NewResponse[management_models.BackOfficeLicenseCustomField](body, err)
}

func (api *LicenseApi) DeleteLicenseCustomField(ctx context.Context, id int) error {
	_, err := api.c.Delete(ctx, "product-custom-fields/{id}", idToParams(id))
	return err
}
