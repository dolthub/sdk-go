package management_request

type CreateLicenseCustomFieldRequest struct {
	Value              string `json:"value,omitempty"`
	License            int64  `json:"license,omitempty"`
	ProductCustomField int    `json:"product_custom_field,omitempty"`
}
