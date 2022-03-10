package management_models

type OrderItem struct {
	ProductCode string
	Licenses    []OrderLicense
}
