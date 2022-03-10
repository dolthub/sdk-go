package core_request

type CustomerLicenseUsersRequest struct {
	Product  string `json:"product"`
	Customer string `json:"customer"`
}
