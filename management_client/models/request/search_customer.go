package management_request

type SearchCustomersRequest struct {
	SearchRequest
	LabelIn              string `json:"label__in"`
	Email                string `json:"email"`
	EmailIcontains       string `json:"email__icontains"`
	EmailStartswith      string `json:"email__startswith"`
	CompanyNameIcontains string `json:"company_name__icontains"`
	FirstNameIcontains   string `json:"first_name__icontains"`
	LastNameIcontains    string `json:"last_name__icontains"`
	PhoneIcontains       string `json:"phone__icontains"`
	ReferenceExact       string `json:"reference__exact"`
	ReferenceIcontains   string `json:"reference__icontains"`
	ReferenceStartswith  string `json:"reference__startswith"`
	Account              string `json:"account"`
	AccountName          string `json:"account__name"`
	AccountNameIcontains string `json:"account__name__icontains"`
	AccountCode          string `json:"account__code"`
	AccountCodeIcontains string `json:"account__code__icontains"`
}
