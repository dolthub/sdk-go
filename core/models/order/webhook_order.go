package order

import (
	"gitlab.com/l216/sdk/sdk-go/v1/core/models/response"
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
