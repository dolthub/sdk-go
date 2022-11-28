package management_request

type SearchProductsRequest struct {
	Limit          int    `json:"limit,omitempty"`
	Offset         int    `json:"offset,omitempty"`
	OrderBy        string `json:"order_by,omitempty"`
	IsArchived     bool   `json:"is_archived,omitempty"`
	SortDescending bool   `json:"sort_descending,omitempty"`
}
