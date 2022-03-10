package core_models

type ConsumptionResponse struct {
	MaxConsumptions   int  `json:"max_consumptions"`
	TotalConsumptions int  `json:"total_consumptions"`
	AllowOverages     bool `json:"allow_overages"`
	MaxOverages       int  `json:"max_overages"`
}
