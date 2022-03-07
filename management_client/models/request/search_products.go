package management_request

type SearchProductsRequest struct {
	Limit          int
	Offset         int
	OrderBy        string
	IsArchived     bool
	SortDescending bool
}
