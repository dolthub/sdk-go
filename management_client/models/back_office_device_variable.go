package management_models

import (
	"time"
)

type BackOfficeDeviceVariable struct {
	Id        int
	CreatedAt time.Time
	Variable  string
	Value     string
	Device    int64
}
