package management_request

type AddManagerToOrderRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
