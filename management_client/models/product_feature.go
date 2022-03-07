package management_models

import (
	"time"
)

type ProductFeature struct {
	Id                int64     `json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	IsDeleted         bool      `json:"isDeleted"`
	FeatureType       string    `json:"featureType"`
	MaxConsumption    int       `json:"maxConsumption"`
	AllowOverages     bool      `json:"allowOverages"`
	MaxOverages       int       `json:"maxOverages"`
	ResetConsumption  bool      `json:"resetConsumption"`
	ConsumptionPeriod string    `json:"consumptionPeriod"`
}
