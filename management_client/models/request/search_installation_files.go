package management_request

type SearchInstallationFilesRequest struct {
	Limit          int
	Offset         int
	OrderBy        string
	Product        int64
	SortDescending bool
}
