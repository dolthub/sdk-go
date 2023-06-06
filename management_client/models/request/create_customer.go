package management_request

type CreateCustomerRequest struct {
	Email           string `json:"email,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	CompanyName     string `json:"company_name,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Reference       string `json:"reference,omitempty"`
	Address         string `json:"address,omitempty"`
	Postcode        string `json:"postcode,omitempty"`
	City            string `json:"city,omitempty"`
	Country         string `json:"country,omitempty"`
	State           string `json:"state,omitempty"`
	CustomerAccount int    `json:"customer_account,omitempty"`
}
