package management_request

type SearchLicensesRequest struct {
	Limit                       int
	Offset                      int
	OrderBy                     string
	IdIn                        string
	Company                     int
	CustomerId                  int
	CustomerEmail               string
	CustomerCompanyNameContains string
	CustomerReferenceContains   string
	CustomerNameContains        string
	LicenseKeyStartsWith        string
	LicenseKeyContains          string
	LicenseKey                  string
	LicenseKeyIsNull            bool
	LicenseUserEmailStartsWith  string
	LicenseUserEmailContains    string
	LicenseUserEmail            string
	OrderStoreIdStartsWith      string
	OrderStoreIdContains        string
	OrderStoreId                string
	NoteContains                string
	IsAnyFloating               bool
	SortDescending              bool
}
