package management_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
	"time"
)

type OrderApi ManagementClient

func (api *OrderApi) SearchOrders(request management_request.SearchOrdersRequest) management_response.SearchResult[management_models.Order] {
	body, err := api.c.Get("orders", nil, request)
	return management_response.SearchResultFromJson[management_models.Order](body, err)
}

func (api *OrderApi) GetOrder(id int64) client.Response[management_models.Order] {
	body, err := api.c.Get("orders/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.Order](body, err)
}

func (api *OrderApi) CreateOrder(request management_models.WebhookOrder) client.Response[management_response.CreateOrderResponse] {
	body, err := api.c.Post("orders/create_order", nil, request)
	return client.NewResponse[management_response.CreateOrderResponse](body, err)
}

func (api *OrderApi) ExportToCsv(from, to time.Time) client.Response[management_response.Response] {
	req := struct {
		From string
		To   string
	}{
		From: from.Format("2006-01-02"),
		To:   to.Format("2006-01-02"),
	}
	body, err := api.c.Get("orders/export", nil, req)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *OrderApi) ExportToCsvRange(r management_models.Range) client.Response[management_response.Response] {
	req := struct {
		Range string
	}{
		Range: string(r),
	}
	body, err := api.c.Get("orders/export", nil, req)
	return client.NewResponse[management_response.Response](body, err)
}

func (api *OrderApi) GenerateLicenseKeys(request management_request.GenerateLicenseRequest) client.Response[[]string] {
	body, err := api.c.Get("orders/generate_license", nil, request)
	return client.NewResponse[[]string](body, err)
}

func (api *OrderApi) AddManagerToOrder(id int64, request management_request.AddManagerToOrderRequest) client.Response[management_models.Manager] {
	body, err := api.c.Post("orders/{id}/add_manager/", id64ToParams(id), request)
	return client.NewResponse[management_models.Manager](body, err)
}
