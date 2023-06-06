package management_models

type CustomerAccount struct {
	Id                int64  `json:"id,omitempty"`
	CustomersCount    int    `json:"customers_count,omitempty"`
	Code              string `json:"code,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Address           string `json:"address,omitempty"`
	Email             string `json:"email,omitempty"`
	CognitoUserPoolId string `json:"cognito_user_pool_id,omitempty"`
	Company           int    `json:"company,omitempty"`
}
