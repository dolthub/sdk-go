package management_response

import (
	"licensespring/go-sdk/management_client/models"
)

type CreateOrderResponse struct {
	OrderId      int64
	NewPasswords []management_models.LicenseUserPassword
}
