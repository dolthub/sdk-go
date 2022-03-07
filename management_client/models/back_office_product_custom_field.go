package management_models

type BackOfficeProductCustomField struct {
	Id           int64
	Name         string
	DataType     string
	DefaultValue string
	Description  string
	Product      int64
}
