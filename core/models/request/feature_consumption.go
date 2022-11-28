package core_request

type FeatureConsumptionRequest struct {
	LicenseRequest

	Feature      string `json:"feature,omitempty"`
	Consumptions int    `json:"consumptions,omitempty"`
}
