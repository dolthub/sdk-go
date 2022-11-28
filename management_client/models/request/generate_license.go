package management_request

type GenerateLicenseRequest struct {
	Product  string `json:"product,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
