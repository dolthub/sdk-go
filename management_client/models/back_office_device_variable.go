package management_models

import (
	"time"
)

type BackOfficeDeviceVariable struct {
	Id        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Variable  string    `json:"variable,omitempty"`
	Value     string    `json:"value,omitempty"`
	Device    int64     `json:"device,omitempty"`
}
