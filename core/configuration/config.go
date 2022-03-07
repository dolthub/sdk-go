package configuration

type CoreConfiguration struct {
	SdkVersion     string
	UrlBase        string
	EndpointPrefix string
	ProductCode    string

	AdditionalHeaders map[string]string

	EnableRetry bool

	Keys
}

type Keys struct {
	ApiKey    string
	SharedKey string
}

func NewClientConfig() *CoreConfiguration {
	return &CoreConfiguration{
		SdkVersion:        "",
		UrlBase:           "https://api.licensespring.com",
		EndpointPrefix:    "/api/v4/",
		ProductCode:       "",
		AdditionalHeaders: make(map[string]string),
		EnableRetry:       true,
		Keys:              Keys{},
	}
}

func (config *CoreConfiguration) SetKeys(apiKey, sharedKey string) {
	config.Keys = Keys{
		ApiKey:    apiKey,
		SharedKey: sharedKey,
	}
}

func (config *CoreConfiguration) AddHeader(key, value string) {
	config.AdditionalHeaders[key] = value
}
