package management_response

import management_models "gitlab.com/l3178/sdk-go/management_client/models"

type CreateOrderResponse struct {
	OrderId      int64
	NewPasswords []management_models.LicenseUserPassword
}
