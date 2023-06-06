package management_client

import (
	"gitlab.com/l3178/sdk-go/core/client"
	management_models "gitlab.com/l3178/sdk-go/management_client/models"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	management_response "gitlab.com/l3178/sdk-go/management_client/models/response"
)

type CustomerApi ManagementClient

// CUSTOMERS

func (api *CustomerApi) ListCustomers(request management_request.SearchCustomersRequest) management_response.SearchResult[management_models.BackOfficeCustomer] {
	body, err := api.c.Get("customers", nil, request)
	return management_response.SearchResultFromJson[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) ShowCustomer(id int64) client.Response[management_models.BackOfficeCustomer] {
	body, err := api.c.Get("customers/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) CreateCustomer(request management_request.CreateCustomerRequest) client.Response[management_models.BackOfficeCustomer] {
	body, err := api.c.Post("customers", nil, request)
	return client.NewResponse[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) EditCustomer(id int64, request management_request.CreateCustomerRequest) client.Response[management_models.BackOfficeCustomer] {
	body, err := api.c.Patch("customers/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.BackOfficeCustomer](body, err)
}

func (api *CustomerApi) DeleteCustomer(id int64) error {
	_, err := api.c.Delete("customers/{id}", id64ToParams(id))
	return err
}

// CUSTOMER LABELS

func (api *CustomerApi) ListCustomerLabels(request management_request.SearchRequest) management_response.SearchResult[management_models.CustomerLabel] {
	body, err := api.c.Get("clabels", nil, request)
	return management_response.SearchResultFromJson[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) ShowCustomerLabel(id int64) client.Response[management_models.CustomerLabel] {
	body, err := api.c.Get("clabels/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) CreateCustomerLabel(request management_request.CreateCustomerLabelRequest) client.Response[management_models.CustomerLabel] {
	body, err := api.c.Post("clabels", nil, request)
	return client.NewResponse[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) EditCustomerLabel(id int64, request management_request.CreateCustomerLabelRequest) client.Response[management_models.CustomerLabel] {
	body, err := api.c.Patch("clabels/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.CustomerLabel](body, err)
}

func (api *CustomerApi) DeleteCustomerLabel(id int64) error {
	_, err := api.c.Delete("clabels/{id}", id64ToParams(id))
	return err
}

// CUSTOMER LABEL MAPPING

func (api *CustomerApi) AddLabelToCustomer(request management_models.CustomerLabelMapping) client.Response[management_models.CustomerLabelMapping] {
	body, err := api.c.Post("customerclabels", nil, request)
	return client.NewResponse[management_models.CustomerLabelMapping](body, err)
}

func (api *CustomerApi) RemoveLabelFromCustomer(id int64) error {
	_, err := api.c.Delete("customerclabels/{id}", id64ToParams(id))
	return err
}

// CUSTOMER ACCOUNTS

func (api *CustomerApi) ListCustomerAccounts(request management_request.SearchRequest) management_response.SearchResult[management_models.CustomerAccount] {
	body, err := api.c.Get("customer-accounts", nil, request)
	return management_response.SearchResultFromJson[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) ShowCustomerAccount(id int64) client.Response[management_models.CustomerAccount] {
	body, err := api.c.Get("customer-accounts/{id}", id64ToParams(id), nil)
	return client.NewResponse[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) CreateCustomerAccount(request management_request.CreateCustomerAccountRequest) client.Response[management_models.CustomerAccount] {
	body, err := api.c.Post("customer-accounts", nil, request)
	return client.NewResponse[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) EditCustomerAccount(id int64, request management_request.CreateCustomerAccountRequest) client.Response[management_models.CustomerAccount] {
	body, err := api.c.Patch("customer-accounts/{id}", id64ToParams(id), request)
	return client.NewResponse[management_models.CustomerAccount](body, err)
}

func (api *CustomerApi) DeleteCustomerAccount(id int64) error {
	_, err := api.c.Delete("customer-accounts/{id}", id64ToParams(id))
	return err
}
