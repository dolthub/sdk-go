package core_request

type FeatureConsumptionRequest struct {
	LicenseRequest

	Feature      string
	Consumptions int `json:"consumptions,omitempty"`
}
