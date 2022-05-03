package management_client

import (
	"encoding/json"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
	"time"
)

type OrderApi ManagementClient

func (api *OrderApi) SearchOrders(request management_request.SearchOrdersRequest) (resp management_response.SearchResult[management_models.Order], err error) {
	err = api.c.Get("orders", nil, request, &resp)
	return resp, err
}

func (api *OrderApi) GetOrder(id int64) (resp management_models.Order, err error) {
	err = api.c.Get("orders/{id}", id64ToParams(id), nil, &resp)
	return resp, err
}

func (api *OrderApi) CreateOrder(request management_models.WebhookOrder) (resp management_response.CreateOrderResponse, err error) {
	err = api.c.Post("orders/create_order", nil, request, &resp)
	return resp, err
}

func (api *OrderApi) ExportToCsv(from, to time.Time) (resp management_response.Response, err error) {
	req, _ := json.Marshal(map[string]interface{}{
		"from": from,
		"to":   to,
	})
	err = api.c.Get("orders/export", nil, req, &resp)
	return resp, err
}

func (api *OrderApi) ExportToCsvRange(r management_models.Range) (resp management_response.Response, err error) {
	req, _ := json.Marshal(map[string]interface{}{
		"range": r,
	})
	err = api.c.Get("orders/export", nil, req, &resp)
	return resp, err
}

func (api *OrderApi) GenerateLicenseKeys(request management_request.GenerateLicenseRequest) (resp []string, err error) {
	err = api.c.Get("orders/generate_license", nil, request, &resp)
	return resp, err
}

func (api *OrderApi) AddManagerToOrder(id int64, request management_request.AddManagerToOrderRequest) (resp management_models.Manager, err error) {
	err = api.c.Post("orders/{id}/add_manager/", id64ToParams(id), request, &resp)
	return resp, err
}
