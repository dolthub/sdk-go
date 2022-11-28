package management_response

import management_models "gitlab.com/l3178/sdk-go/management_client/models"

type CreateOrderResponse struct {
	OrderId      int64                                   `json:"order_id,omitempty"`
	NewPasswords []management_models.LicenseUserPassword `json:"new_passwords,omitempty"`
}
