package floating_client

type ProductRequest struct {
	Product string `json:"product"`
}

type FloatingRegistrationRequest struct {
	User    string `json:"user"`
	Product string `json:"product"`
}
