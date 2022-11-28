package management_models

import (
	"time"
)

type Order struct {
	Id                   int64                 `json:"id,omitempty"`
	BackOfficeOrderItems []BackOfficeOrderItem `json:"back_office_order_items,omitempty"`
	Customer             BackOfficeCustomer    `json:"customer"`
	HasActiveLicenses    bool                  `json:"has_active_licenses,omitempty"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
	ClientOrderId        string                `json:"client_order_id,omitempty"`
	Type                 OrderType             `json:"type,omitempty"`
	IsTest               bool                  `json:"is_test,omitempty"`
	IsTrial              bool                  `json:"is_trial,omitempty"`
	DownloadId           string                `json:"download_id,omitempty"`
	Language             string                `json:"language,omitempty"`
	CampaignParams       string                `json:"campaign_params,omitempty"`
	Company              int                   `json:"company,omitempty"`
}
