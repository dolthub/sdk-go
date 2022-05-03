package management_client

import (
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type ProductApi ManagementClient

// PRODUCTS

func (api *ProductApi) ListProducts(request management_request.SearchProductsRequest) (resp management_response.SearchResult[management_models.BackOfficeProduct], err error) {
	err = api.c.Get("products", nil, request, &resp)
	return resp, err
}

func (api *ProductApi) ShowProduct(id int64) (resp management_models.BackOfficeProduct, err error) {
	err = api.c.Get("products/{id}", nil, id64ToParams(id), &resp)
	return resp, err
}

// INSTALLATION FILES

func (api *ProductApi) ListInstallationFiles(request management_request.SearchInstallationFilesRequest) (resp management_response.SearchResult[management_models.BackOfficeInstallationFile], err error) {
	err = api.c.Get("installation-files", nil, request, &resp)
	return resp, err
}

func (api *ProductApi) ShowInstallationFile(id int) (resp management_models.BackOfficeInstallationFile, err error) {
	err = api.c.Get("installation-files/{id}", idToParams(id), nil, &resp)
	return resp, err
}

func (api *ProductApi) CreateInstallationFile(request management_request.CreateInstallationFileRequest) (resp management_models.BackOfficeInstallationFile, err error) {
	err = api.c.Post("installation-files", nil, request, &resp)
	return resp, err
}

// CUSTOM FIELDS

func (api *ProductApi) ListProductCustomFields(request management_request.SearchProductCustomFieldsRequest) (resp management_response.SearchResult[management_models.BackOfficeProductCustomField], err error) {
	err = api.c.Get("product-custom-fields", nil, request, &resp)
	return resp, err
}

func (api *ProductApi) ShowProductCustomField(id int) (resp management_models.BackOfficeProductCustomField, err error) {
	err = api.c.Get("product-custom-fields/{id}", nil, idToParams(id), &resp)
	return resp, err
}

func (api *ProductApi) CreateProductCustomField(request management_request.CreateProductCustomFieldRequest) (resp management_models.BackOfficeProductCustomField, err error) {
	err = api.c.Post("product-custom-fields", nil, request, &resp)
	return resp, err
}

func (api *ProductApi) UpdateProductCustomField(id int, request management_request.UpdateProductCustomFieldRequest) (resp management_models.BackOfficeProductCustomField, err error) {
	err = api.c.Patch("product-custom-fields/{id}", idToParams(id), request, &resp)
	return resp, err
}

func (api *ProductApi) DeleteProductCustomField(id int) error {
	return api.c.Delete("product-custom-fields/{id}", idToParams(id), nil)
}
