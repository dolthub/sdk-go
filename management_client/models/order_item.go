package management_models

type OrderItem struct {
	ProductCode string         `json:"product_code,omitempty"`
	Licenses    []OrderLicense `json:"licenses,omitempty"`
}
