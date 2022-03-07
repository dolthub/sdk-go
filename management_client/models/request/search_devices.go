package management_request

type SearchDevicesRequest struct {
	Limit          int
	Offset         int
	OrderBy        string
	License        int64
	Blacklisted    bool
	Hostname       string
	HardwareId     string
	SortDescending bool
}
