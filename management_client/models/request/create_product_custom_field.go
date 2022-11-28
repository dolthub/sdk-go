package management_request

type CreateProductCustomFieldRequest struct {
	Name         string `json:"name,omitempty"`
	DefaultValue string `json:"default_value,omitempty"`
	Product      int64  `json:"product,omitempty"`
}
