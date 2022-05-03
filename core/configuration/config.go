package configuration

import (
	"time"
)

type Config interface {
	AuthHeader(date time.Time) string
	UrlPrefix() string
	GetRequestConfig() RequestConfig
}
