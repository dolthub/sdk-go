package management_response

type CreateOrderResponse struct {
	OrderId      int64
	NewPasswords []management_models.LicenseUserPassword
}
