package management_models

type BackOfficeCustomer struct {
	Id            int64           `json:"id,omitempty"`
	TotalOrders   int             `json:"total_orders,omitempty"`
	TotalLicenses int             `json:"total_licenses,omitempty"`
	Labels        []CustomerLabel `json:"labels,omitempty"`
	Email         string          `json:"email,omitempty"`
	FirstName     string          `json:"first_name,omitempty"`
	LastName      string          `json:"last_name,omitempty"`
	CompanyName   string          `json:"company_name,omitempty"`
	Phone         string          `json:"phone,omitempty"`
	Reference     string          `json:"reference,omitempty"`
	Address       string          `json:"address,omitempty"`
	Postcode      string          `json:"postcode,omitempty"`
	City          string          `json:"city,omitempty"`
	Country       string          `json:"country,omitempty"`
	State         string          `json:"state,omitempty"`
	AllLabelNames string          `json:"all_label_names,omitempty"`
	Company       int             `json:"company,omitempty"`
	LicenseUser   int             `json:"license_user,omitempty"`
	LabelIds      []int           `json:"label_ids,omitempty"`
}
