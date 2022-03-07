package management_models

type BackOfficeLicenseFeature struct {
	Id                int64
	ProductFeature    ProductFeature
	MaxConsumption    int
	TotalConsumptions int
}
