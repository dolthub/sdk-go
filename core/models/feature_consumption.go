package core_models

type FeatureConsumptionResponse struct {
	ConsumptionResponse

	ResetConsumption  bool   `json:"reset_consumption,omitempty"`
	ConsumptionPeriod string `json:"consumption_period,omitempty"`
}
