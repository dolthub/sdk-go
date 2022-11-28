package management_request

type SearchLicenseCustomFieldsRequest struct {
	Limit   int   `json:"limit,omitempty"`
	Offset  int   `json:"offset,omitempty"`
	License int64 `json:"license,omitempty"`
}
