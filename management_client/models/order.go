package management_models

import (
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
	Type                 OrderType
	IsTest               bool
	IsTrial              bool
	DownloadId           string
	Language             string
	CampaignParams       string
	Company              int
}
