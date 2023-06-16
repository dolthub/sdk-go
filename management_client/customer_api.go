package management_client

import (
	"context"
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type CustomerApi ManagementClient

// CUSTOMERS

func (api *CustomerApi) ListCustomers(ctx context.Context, request management_request.SearchCustomersRequest) management_response.SearchResult[management_models.BackOfficeCustomer] {
	body, err := api.c.Get(ctx, "customers", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) ShowCustomer(ctx context.Context, id int64) client.Response[management_models.BackOfficeCustomer] {
	body, err := api.c.Get(ctx, "customers/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) CreateCustomer(ctx context.Context, request management_request.CreateCustomerRequest) client.Response[management_models.BackOfficeCustomer] {
	body, err := api.c.Post(ctx, "customers", nil, request)
	return client.NewResponse[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) EditCustomer(ctx context.Context, id int64, request management_request.CreateCustomerRequest) client.Response[management_models.BackOfficeCustomer] {
	body, err := api.c.Patch(ctx, "customers/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) DeleteCustomer(ctx context.Context, id int64) error {
	_, err := api.c.Delete(ctx, "customers/{id}", id64ToParams(id))
	return err
}

// CUSTOMER LABELS

func (api *CustomerApi) ListCustomerLabels(ctx context.Context, request management_request.SearchRequest) management_response.SearchResult[management_models.CustomerLabel] {
	body, err := api.c.Get(ctx, "clabels", nil, request)
	return management_response.SearchResultFromJson[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) ShowCustomerLabel(ctx context.Context, id int64) client.Response[management_models.CustomerLabel] {
	body, err := api.c.Get(ctx, "clabels/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) CreateCustomerLabel(ctx context.Context, request management_request.CreateCustomerLabelRequest) client.Response[management_models.CustomerLabel] {
	body, err := api.c.Post(ctx, "clabels", nil, request)
	return client.NewResponse[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) EditCustomerLabel(ctx context.Context, id int64, request management_request.CreateCustomerLabelRequest) client.Response[management_models.CustomerLabel] {
	body, err := api.c.Patch(ctx, "clabels/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) DeleteCustomerLabel(ctx context.Context, id int64) error {
	_, err := api.c.Delete(ctx, "clabels/{id}", id64ToParams(id))
	return err
}

// CUSTOMER LABEL MAPPING

func (api *CustomerApi) AddLabelToCustomer(ctx context.Context, request management_models.CustomerLabelMapping) client.Response[management_models.CustomerLabelMapping] {
	body, err := api.c.Post(ctx, "customerclabels", nil, request)
	return client.NewResponse[management_models.CustomerLabelMapping](body, err)
}

func (api *CustomerApi) RemoveLabelFromCustomer(ctx context.Context, id int64) error {
	_, err := api.c.Delete(ctx, "customerclabels/{id}", id64ToParams(id))
	return err
}

// CUSTOMER ACCOUNTS

func (api *CustomerApi) ListCustomerAccounts(ctx context.Context, request management_request.SearchRequest) management_response.SearchResult[management_models.CustomerAccount] {
	body, err := api.c.Get(ctx, "customer-accounts", nil, request)
	return management_response.SearchResultFromJson[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) ShowCustomerAccount(ctx context.Context, id int64) client.Response[management_models.CustomerAccount] {
	body, err := api.c.Get(ctx, "customer-accounts/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) CreateCustomerAccount(ctx context.Context, request management_request.CreateCustomerAccountRequest) client.Response[management_models.CustomerAccount] {
	body, err := api.c.Post(ctx, "customer-accounts", nil, request)
	return client.NewResponse[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) EditCustomerAccount(ctx context.Context, id int64, request management_request.CreateCustomerAccountRequest) client.Response[management_models.CustomerAccount] {
	body, err := api.c.Patch(ctx, "customer-accounts/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) DeleteCustomerAccount(ctx context.Context, id int64) error {
	_, err := api.c.Delete(ctx, "customer-accounts/{id}", id64ToParams(id))
	return err
}
