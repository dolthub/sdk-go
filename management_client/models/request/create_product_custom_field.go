package management_request

type CreateProductCustomFieldRequest struct {
	Name         string
	DefaultValue string
	Product      int64
}
