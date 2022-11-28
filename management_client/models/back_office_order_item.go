package management_models

import (
	"time"
)

type BackOfficeOrderItem struct {
	Id        int                 `json:"id,omitempty"`
	ProductId string              `json:"product_id,omitempty"`
	Licenses  []BackOfficeLicense `json:"licenses,omitempty"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Order     int64               `json:"order,omitempty"`
	Product   int64               `json:"product,omitempty"`
}
