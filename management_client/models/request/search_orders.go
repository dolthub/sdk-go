package management_request

type SearchOrdersRequest struct {
	Limit                       int    `json:"limit,omitempty"`
	Offset                      int    `json:"offset,omitempty"`
	OrderBy                     string `json:"order_by,omitempty"`
	Company                     int    `json:"company,omitempty"`
	CustomerId                  int    `json:"customer_id,omitempty"`
	CustomerEmail               string `json:"customer_email,omitempty"`
	CustomerEmailStartsWith     string `json:"customer_email_starts_with,omitempty"`
	CustomerEmailContains       string `json:"customer_email_contains,omitempty"`
	CustomerCompanyNameContains string `json:"customer_company_name_contains,omitempty"`
	CustomerCompanyName         string `json:"customer_company_name,omitempty"`
	CustomerReferenceContains   string `json:"customer_reference_contains,omitempty"`
	CustomerNameContains        string `json:"customer_name_contains,omitempty"`
	ClientOrderId               string `json:"client_order_id,omitempty"`
	ClientOrderIdStartsWith     string `json:"client_order_id_starts_with,omitempty"`
	ClientOrderIdContains       string `json:"client_order_id_contains,omitempty"`
	CustomerLabel               string `json:"customer_label,omitempty"`
	SortDescending              bool   `json:"sort_descending,omitempty"`
}
