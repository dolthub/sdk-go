package request

type ConsumptionRequest struct {
	LicenseRequest

	Consumptions  int  `json:"consumptions,omitempty"`
	MaxOverages   int  `json:"max_overages,omitempty"`
	AllowOverages bool `json:"allow_overages,omitempty"`
}
