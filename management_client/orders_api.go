package management_client

import (
	"encoding/json"
	"gitlab.com/l3178/sdk/sdk-go/v1/core/client"
	management_models "gitlab.com/l3178/sdk/sdk-go/v1/management_client/models"
	management_request "gitlab.com/l3178/sdk/sdk-go/v1/management_client/models/request"
	management_response "gitlab.com/l3178/sdk/sdk-go/v1/management_client/models/response"
	"time"
)

type OrdersApi struct {
	c *client.Client
}

func (api *OrdersApi) SearchOrders(request management_request.SearchOrdersRequest) (resp management_response.SearchResult[management_models.Order], err error) {
	err = api.c.Get("orders", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) GetOrder(id int64) (resp management_models.Order, err error) {
	err = api.c.Get("orders/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *OrdersApi) CreateOrder(request management_models.WebhookOrder) (resp management_response.CreateOrderResponse, err error) {
	err = api.c.Post("orders/create_order", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) ExportToCsv(from, to time.Time) (resp management_response.Response, err error) {
	req, _ := json.Marshal(map[string]interface{}{
		"from": from,
		"to":   to,
	})
	err = api.c.Get("orders/export", nil, req, &resp)
	return resp, err
}

func (api *OrdersApi) ExportToCsvRange(r management_models.Range) (resp management_response.Response, err error) {
	req, _ := json.Marshal(map[string]interface{}{
		"range": r,
	})
	err = api.c.Get("orders/export", nil, req, &resp)
	return resp, err
}

func (api *OrdersApi) GenerateLicenseKeys(request management_request.GenerateLicenseRequest) (resp []string, err error) {
	err = api.c.Get("orders/generate_license", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) AddManagerToOrder(id int64, request management_request.AddManagerToOrderRequest) (resp management_models.Manager, err error) {
	err = api.c.Post("orders/{id}/add_manager/", id64ToParams(id), request, &resp)
	return resp, err
}

func (api *OrdersApi) SearchProducts(request management_request.SearchProductsRequest) (resp management_response.SearchResult[management_models.BackOfficeProduct], err error) {
	err = api.c.Get("products", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) GetProduct(id int64) (resp management_models.BackOfficeProduct, err error) {
	err = api.c.Get("products/{id}", nil, id64ToParams(id), &resp)
	return resp, err
}

func (api *OrdersApi) SearchInstallationFiles(request management_request.SearchInstallationFilesRequest) (resp management_response.SearchResult[management_models.BackOfficeInstallationFile], err error) {
	err = api.c.Get("installation-files", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) GetInstallationFile(id int) (resp management_models.BackOfficeInstallationFile, err error) {
	err = api.c.Get("installation-files/{id}", idToParams(id), nil, &resp)
	return resp, err
}

func (api *OrdersApi) CreateInstallationFile(request management_request.CreateInstallationFileRequest) (resp management_models.BackOfficeInstallationFile, err error) {
	err = api.c.Post("installation-files", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) SearchProductCustomFields(request management_request.SearchProductCustomFieldsRequest) (resp management_response.SearchResult[management_models.BackOfficeProductCustomField], err error) {
	err = api.c.Get("product-custom-fields", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) GetProductCustomField(id int) (resp management_models.BackOfficeProductCustomField, err error) {
	err = api.c.Get("product-custom-fields/{id}", nil, idToParams(id), &resp)
	return resp, err
}

func (api *OrdersApi) CreateProductCustomField(request management_request.CreateProductCustomFieldRequest) (resp management_models.BackOfficeProductCustomField, err error) {
	err = api.c.Post("product-custom-fields", nil, request, &resp)
	return resp, err
}

func (api *OrdersApi) UpdateProductCustomField(id int, request management_request.UpdateProductCustomFieldRequest) (resp management_models.BackOfficeProductCustomField, err error) {
	err = api.c.Patch("product-custom-fields/{id}", idToParams(id), request, &resp)
	return resp, err
}

func (api *OrdersApi) DeleteProductCustomField(id int) error {
	return api.c.Delete("product-custom-fields/{id}", idToParams(id), nil)
}
