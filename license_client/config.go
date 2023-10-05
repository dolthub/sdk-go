package license_client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"gitlab.com/l3178/sdk-go/core/configuration"
	"time"
)

const (
	baseUrl   = "https://api.licensespring.com"
	apiPrefix = "/api/v4/"

	signaturePrefix = "licenseSpring"
	headers         = "date"
	algorithm       = "hmac-sha256"
	dateFormat      = time.RFC1123
)

type LicenseClientConfiguration struct {
	*configuration.CoreConfiguration

	ApiKey    string
	SharedKey string

	ProductCode string
	HardwareId  string

	AppName    string
	AppVersion string

	VerifySignature bool
}

func (config LicenseClientConfiguration) AuthHeader(date time.Time) string {
	dateStr := date.Format(dateFormat)
	signatureString := fmt.Sprintf("%s\ndate: %s", signaturePrefix, dateStr)

	encoder := hmac.New(sha256.New, []byte(config.SharedKey))
	encoder.Write([]byte(signatureString))
	encodedSignature := base64.StdEncoding.EncodeToString(encoder.Sum(nil))

	return fmt.Sprintf("algorithm=\"%s\",headers=\"%s\",signature=\"%s\",apikey=\"%s\"",
		algorithm, headers, encodedSignature, config.ApiKey)
}

func (config LicenseClientConfiguration) OfflineLicenseSignature(date time.Time, licenseKey string) string {
	dateStr := date.Format(dateFormat)
	signatureString := fmt.Sprintf("%s\ndate: %s\n%s\n%s\n%s",
		signaturePrefix, dateStr, licenseKey, config.HardwareId, config.ApiKey)

	encoder := hmac.New(sha256.New, []byte(config.SharedKey))
	encoder.Write([]byte(signatureString))
	return base64.StdEncoding.EncodeToString(encoder.Sum(nil))
}

func (config LicenseClientConfiguration) UrlPrefix() string {
	return config.CoreConfiguration.UrlPrefix()
}

func NewLicenseClientConfiguration(apiKey, sharedKey, productCode string) LicenseClientConfiguration {
	config := configuration.NewClientConfig(baseUrl, apiPrefix)

	hardwareId, err := machineid.ProtectedID(configuration.SdkName)
	if err != nil {
		panic(err)
	}

	return LicenseClientConfiguration{
		CoreConfiguration: config,
		ApiKey:            apiKey,
		SharedKey:         sharedKey,
		ProductCode:       productCode,
		HardwareId:        hardwareId,
		VerifySignature:   true,
	}
}
