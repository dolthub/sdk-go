package management_request

type SearchInstallationFilesRequest struct {
	Limit          int    `json:"limit,omitempty"`
	Offset         int    `json:"offset,omitempty"`
	OrderBy        string `json:"order_by,omitempty"`
	Product        int64  `json:"product,omitempty"`
	SortDescending bool   `json:"sort_descending,omitempty"`
}
