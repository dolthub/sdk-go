package management_models

import (
	"time"
)

type ProductFeature struct {
	Id                int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Name              string
	Code              string
	IsDeleted         bool
	FeatureType       string
	MaxConsumption    int
	AllowOverages     bool
	MaxOverages       int
	ResetConsumption  bool
	ConsumptionPeriod string
}
