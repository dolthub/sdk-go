package management_request

type SearchProductCustomFieldsRequest struct {
	Limit   int   `json:"limit,omitempty"`
	Offset  int   `json:"offset,omitempty"`
	Product int64 `json:"product,omitempty"`
}
