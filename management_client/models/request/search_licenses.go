package management_request

type SearchLicensesRequest struct {
	Limit                       int    `json:"limit,omitempty"`
	Offset                      int    `json:"offset,omitempty"`
	OrderBy                     string `json:"order_by,omitempty"`
	IdIn                        string `json:"id_in,omitempty"`
	Company                     int    `json:"company,omitempty"`
	CustomerId                  int    `json:"customer_id,omitempty"`
	CustomerEmail               string `json:"customer_email,omitempty"`
	CustomerCompanyNameContains string `json:"customer_company_name_contains,omitempty"`
	CustomerReferenceContains   string `json:"customer_reference_contains,omitempty"`
	CustomerNameContains        string `json:"customer_name_contains,omitempty"`
	LicenseKeyStartsWith        string `json:"license_key_starts_with,omitempty"`
	LicenseKeyContains          string `json:"license_key_contains,omitempty"`
	LicenseKey                  string `json:"license_key,omitempty"`
	LicenseKeyIsNull            bool   `json:"license_key_is_null,omitempty"`
	LicenseUserEmailStartsWith  string `json:"license_user_email_starts_with,omitempty"`
	LicenseUserEmailContains    string `json:"license_user_email_contains,omitempty"`
	LicenseUserEmail            string `json:"license_user_email,omitempty"`
	OrderStoreIdStartsWith      string `json:"order_store_id_starts_with,omitempty"`
	OrderStoreIdContains        string `json:"order_store_id_contains,omitempty"`
	OrderStoreId                string `json:"order_store_id,omitempty"`
	NoteContains                string `json:"note_contains,omitempty"`
	IsAnyFloating               bool   `json:"is_any_floating,omitempty"`
	SortDescending              bool   `json:"sort_descending,omitempty"`
}
