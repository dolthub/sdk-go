package core_models

type ConsumptionResponse struct {
	MaxConsumptions   int  `json:"max_consumptions,omitempty"`
	TotalConsumptions int  `json:"total_consumptions,omitempty"`
	AllowOverages     bool `json:"allow_overages,omitempty"`
	MaxOverages       int  `json:"max_overages,omitempty"`
}
