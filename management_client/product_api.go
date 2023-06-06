package management_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type ProductApi ManagementClient

// PRODUCTS

func (api *ProductApi) ListProducts(request management_request.SearchProductsRequest) management_response.SearchResult[management_models.BackOfficeProduct] {
	body, err := api.c.Get("products", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeProduct](body, err)
}

func (api *ProductApi) ShowProduct(id int64) client.Response[management_models.BackOfficeProduct] {
	body, err := api.c.Get("products/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeProduct](body, err)
}

// INSTALLATION FILES

func (api *ProductApi) ListInstallationFiles(request management_request.SearchInstallationFilesRequest) management_response.SearchResult[management_models.BackOfficeInstallationFile] {
	body, err := api.c.Get("installation-files", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeInstallationFile](body, err)
}

func (api *ProductApi) ShowInstallationFile(id int) client.Response[management_models.BackOfficeInstallationFile] {
	body, err := api.c.Get("installation-files/{id}", idToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeInstallationFile](body, err)
}

func (api *ProductApi) CreateInstallationFile(request management_request.CreateInstallationFileRequest) client.Response[management_models.BackOfficeInstallationFile] {
	body, err := api.c.Post("installation-files", nil, request)
	return client.NewResponse[management_models.BackOfficeInstallationFile](body, err)
}

// CUSTOM FIELDS

func (api *ProductApi) ListProductCustomFields(request management_request.SearchProductCustomFieldsRequest) management_response.SearchResult[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Get("product-custom-fields", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) ShowProductCustomField(id int) client.Response[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Get("product-custom-fields/{id}", idToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) CreateProductCustomField(request management_request.CreateProductCustomFieldRequest) client.Response[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Post("product-custom-fields", nil, request)
	return client.NewResponse[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) UpdateProductCustomField(id int, request management_request.UpdateProductCustomFieldRequest) client.Response[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Patch("product-custom-fields/{id}", idToParams(id), request)
	return client.NewResponse[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) DeleteProductCustomField(id int) error {
	_, err := api.c.Delete("product-custom-fields/{id}", idToParams(id))
	return err
}
