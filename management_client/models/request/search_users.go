package management_request

type SearchUsersRequest struct {
	Limit             int    `json:"limit,omitempty"`
	Offset            int    `json:"offset,omitempty"`
	EmailContains     string `json:"true_email__icontains,omitempty"`
	EmailExact        string `json:"true_email__iexact,omitempty"`
	ManagingOrders    string `json:"managing_orders__isnull,omitempty"`
	FirstNameContains string `json:"first_name__icontains,omitempty"`
	LastNameContains  string `json:"last_name__icontains,omitempty"`
	PhoneContains     string `json:"phone__icontains,omitempty"`
}
