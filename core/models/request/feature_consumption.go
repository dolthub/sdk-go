package core_request

type FeatureConsumptionRequest struct {
	LicenseRequest

	Feature      string `json:"feature"`
	Consumptions int    `json:"consumptions,omitempty"`
}
