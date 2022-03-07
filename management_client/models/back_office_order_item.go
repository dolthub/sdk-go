package management_models

import (
	"time"
)

type BackOfficeOrderItem struct {
	Id        int
	ProductId string
	Licenses  []BackOfficeLicense
	CreatedAt time.Time
	UpdatedAt time.Time
	Order     int64
	Product   int64
}
