package management_models

type BackOfficeCustomField struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	DataType     string `json:"data_type,omitempty"`
	DefaultValue string `json:"default_value,omitempty"`
	Description  string `json:"description,omitempty"`
}
