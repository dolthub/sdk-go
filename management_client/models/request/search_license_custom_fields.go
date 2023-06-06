package management_request

type SearchLicenseCustomFieldsRequest struct {
	Limit   int   `json:"limit,omitempty" url:"limit"`
	Offset  int   `json:"offset,omitempty" url:"offset"`
	License int64 `json:"license,omitempty" url:"license"`
}
