package management_request

type CreateCustomerLabelRequest struct {
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
}
