package configuration

import (
	"github.com/sirupsen/logrus"
)

type CoreConfiguration struct {
	BaseUrl   string
	ApiPrefix string

	Logger *logrus.Logger

	RequestConfig
}

type RequestConfig struct {
	RetryCount int

	RequestTimeout int
}

func (config *CoreConfiguration) UrlPrefix() string {
	return config.BaseUrl + config.ApiPrefix
}

func (config *CoreConfiguration) LoggingLevel(level logrus.Level) {
	config.Logger.SetLevel(level)
}

func (config *CoreConfiguration) GetRequestConfig() RequestConfig {
	return config.RequestConfig
}

func NewClientConfig(baseUrl, apiPrefix string) *CoreConfiguration {
	return &CoreConfiguration{
		BaseUrl:   baseUrl,
		ApiPrefix: apiPrefix,
		Logger:    logrus.New(),
		RequestConfig: RequestConfig{
			RetryCount:     3,
			RequestTimeout: 10,
		},
	}
}
