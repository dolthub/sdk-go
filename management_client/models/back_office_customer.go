package management_models

type BackOfficeCustomer struct {
	Id            int
	TotalOrders   int
	TotalLicenses int
	Labels        []CustomerLabel
	Email         string
	FirstName     string
	LastName      string
	CompanyName   string
	Phone         string
	Reference     string
	Address       string
	Postcode      string
	City          string
	Country       string
	State         string
	AllLabelNames string
	Company       int
	LicenseUser   string
	LabelIds      []int
}
