package configuration

type CoreConfiguration struct {
	BaseUrl   string
	ApiPrefix string

	Verbose bool

	RequestConfig
}

type RequestConfig struct {
	RetryCount int

	RequestTimeout int
}

func (config *CoreConfiguration) UrlPrefix() string {
	return config.BaseUrl + config.ApiPrefix
}

func (config *CoreConfiguration) GetRequestConfig() RequestConfig {
	return config.RequestConfig
}

func NewClientConfig(baseUrl, apiPrefix string) *CoreConfiguration {
	return &CoreConfiguration{
		BaseUrl:   baseUrl,
		ApiPrefix: apiPrefix,
		RequestConfig: RequestConfig{
			RetryCount:     3,
			RequestTimeout: 10,
		},
	}
}
