package management_models

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models/order"
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
