package management_models

import (
	"licensespring/go-sdk/core/models/order"
	"time"
)

type Order struct {
	Id                   int64
	BackOfficeOrderItems []BackOfficeOrderItem
	Customer             BackOfficeCustomer
	HasActiveLicenses    bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
	ClientOrderId        string
	Type                 order.OrderType
	IsTest               bool
	IsTrial              bool
	DownloadId           string
	Language             string
	CampaignParams       string
	Company              int
}
