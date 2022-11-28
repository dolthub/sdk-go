package management_models

type BackOfficeLicenseCustomField struct {
	Id                 int    `json:"id,omitempty"`
	Value              string `json:"value,omitempty"`
	License            int64  `json:"license,omitempty"`
	ProductCustomField int    `json:"product_custom_field,omitempty"`
}
