package management_request

type SearchDevicesRequest struct {
	Limit          int    `json:"limit,omitempty"`
	Offset         int    `json:"offset,omitempty"`
	OrderBy        string `json:"order_by,omitempty"`
	License        int64  `json:"license,omitempty"`
	Blacklisted    bool   `json:"blacklisted,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	HardwareId     string `json:"hardware_id,omitempty"`
	SortDescending bool   `json:"sort_descending,omitempty"`
}
