package core_request

type CustomerLicenseUsersRequest struct {
	Product  string `json:"product,omitempty"`
	Customer string `json:"customer,omitempty"`
}
