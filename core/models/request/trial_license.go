package core_request

type TrialLicenseRequest struct {
	LicenseRequest

	Email         string `json:"email,omitempty"`
	LicensePolicy string `json:"license_policy,omitempty"`

	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
	PostCode  string `json:"post_code,omitempty"`
	State     string `json:"state,omitempty"`
	Country   string `json:"country,omitempty"`
	City      string `json:"city,omitempty"`
	Reference string `json:"reference,omitempty"`
}
