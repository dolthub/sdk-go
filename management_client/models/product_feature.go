package management_models

import (
	"time"
)

type ProductFeature struct {
	Id                int64     `json:"id,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Name              string    `json:"name,omitempty"`
	Code              string    `json:"code,omitempty"`
	IsDeleted         bool      `json:"is_deleted,omitempty"`
	FeatureType       string    `json:"feature_type,omitempty"`
	MaxConsumption    int       `json:"max_consumption,omitempty"`
	AllowOverages     bool      `json:"allow_overages,omitempty"`
	MaxOverages       int       `json:"max_overages,omitempty"`
	ResetConsumption  bool      `json:"reset_consumption,omitempty"`
	ConsumptionPeriod string    `json:"consumption_period,omitempty"`
}
