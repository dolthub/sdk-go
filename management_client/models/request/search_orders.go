package management_request

type SearchOrdersRequest struct {
	Limit                       int
	Offset                      int
	OrderBy                     string
	Company                     int
	CustomerId                  int
	CustomerEmail               string
	CustomerEmailStartsWith     string
	CustomerEmailContains       string
	CustomerCompanyNameContains string
	CustomerCompanyName         string
	CustomerReferenceContains   string
	CustomerNameContains        string
	ClientOrderId               string
	ClientOrderIdStartsWith     string
	ClientOrderIdContains       string
	CustomerLabel               string
	SortDescending              bool
}
