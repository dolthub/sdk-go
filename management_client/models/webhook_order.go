package management_models

import (
	"time"
)

type WebhookOrder struct {
	Id             string
	Append         bool
	IsTest         bool
	Created        time.Time
	OrderType      OrderType
	Language       string
	CampaignParams string
	DownloadId     string
	Customer       core_models.Customer
	Items          []OrderItem
	PreventVm      bool
}
