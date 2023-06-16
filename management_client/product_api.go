package management_client

import (
	"context"
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type ProductApi ManagementClient

// PRODUCTS

func (api *ProductApi) ListProducts(ctx context.Context, request management_request.SearchProductsRequest) management_response.SearchResult[management_models.BackOfficeProduct] {
	body, err := api.c.Get(ctx, "products", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeProduct](body, err)
}

func (api *ProductApi) ShowProduct(ctx context.Context, id int64) client.Response[management_models.BackOfficeProduct] {
	body, err := api.c.Get(ctx, "products/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeProduct](body, err)
}

// INSTALLATION FILES

func (api *ProductApi) ListInstallationFiles(ctx context.Context, request management_request.SearchInstallationFilesRequest) management_response.SearchResult[management_models.BackOfficeInstallationFile] {
	body, err := api.c.Get(ctx, "installation-files", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeInstallationFile](body, err)
}

func (api *ProductApi) ShowInstallationFile(ctx context.Context, id int) client.Response[management_models.BackOfficeInstallationFile] {
	body, err := api.c.Get(ctx, "installation-files/{id}", idToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeInstallationFile](body, err)
}

func (api *ProductApi) CreateInstallationFile(ctx context.Context, request management_request.CreateInstallationFileRequest) client.Response[management_models.BackOfficeInstallationFile] {
	body, err := api.c.Post(ctx, "installation-files", nil, request)
	return client.NewResponse[management_models.BackOfficeInstallationFile](body, err)
}

// CUSTOM FIELDS

func (api *ProductApi) ListProductCustomFields(ctx context.Context, request management_request.SearchProductCustomFieldsRequest) management_response.SearchResult[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Get(ctx, "product-custom-fields", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) ShowProductCustomField(ctx context.Context, id int) client.Response[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Get(ctx, "product-custom-fields/{id}", idToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) CreateProductCustomField(ctx context.Context, request management_request.CreateProductCustomFieldRequest) client.Response[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Post(ctx, "product-custom-fields", nil, request)
	return client.NewResponse[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) UpdateProductCustomField(ctx context.Context, id int, request management_request.UpdateProductCustomFieldRequest) client.Response[management_models.BackOfficeProductCustomField] {
	body, err := api.c.Patch(ctx, "product-custom-fields/{id}", idToParams(id), request)
	return client.NewResponse[management_models.BackOfficeProductCustomField](body, err)
}

func (api *ProductApi) DeleteProductCustomField(ctx context.Context, id int) error {
	_, err := api.c.Delete(ctx, "product-custom-fields/{id}", idToParams(id))
	return err
}
