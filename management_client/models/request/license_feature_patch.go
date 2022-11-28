package management_request

type LicenseFeaturePatch struct {
	ProductFeature string `json:"product_feature,omitempty"`
	MaxConsumption int64  `json:"max_consumption,omitempty"`
}
