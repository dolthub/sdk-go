package order

import (
	"licensespring/go-sdk/core/models/response"
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
	Customer       response.Customer
	Items          []OrderItem
	PreventVm      bool
}
