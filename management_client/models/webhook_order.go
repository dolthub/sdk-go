package management_models

import (
	core_models "gitlab.com/l3178/sdk-go/core/models"
	"time"
)

type WebhookOrder struct {
	Id             string               `json:"id,omitempty"`
	Append         bool                 `json:"append,omitempty"`
	IsTest         bool                 `json:"is_test,omitempty"`
	Created        time.Time            `json:"created"`
	OrderType      OrderType            `json:"order_type,omitempty"`
	Language       string               `json:"language,omitempty"`
	CampaignParams string               `json:"campaign_params,omitempty"`
	DownloadId     string               `json:"download_id,omitempty"`
	Customer       core_models.Customer `json:"customer"`
	Items          []OrderItem          `json:"items,omitempty"`
	PreventVm      bool                 `json:"prevent_vm,omitempty"`
}
