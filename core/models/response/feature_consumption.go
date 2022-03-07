package response

type FeatureConsumptionResponse struct {
	ConsumptionResponse

	ResetConsumption  bool   `json:"reset_consumption"`
	ConsumptionPeriod string `json:"consumption_period"`
}
