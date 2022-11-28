package management_models

type BackOfficeLicenseFeature struct {
	Id                int64          `json:"id,omitempty"`
	ProductFeature    ProductFeature `json:"product_feature"`
	MaxConsumption    int            `json:"max_consumption,omitempty"`
	TotalConsumptions int            `json:"total_consumptions,omitempty"`
}
