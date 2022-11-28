package management_request

type UpdateProductCustomFieldRequest struct {
	Name         string `json:"name,omitempty"`
	DefaultValue string `json:"default_value,omitempty"`
}
