package license_client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	. "gitlab.com/l3178/sdk-go/core/models"
	"gitlab.com/l3178/sdk-go/license_client/airgap"
)

func (c *LicenseClient) AirgapInitialization(licenseKey, signingKey string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(signingKey)
	if err != nil {
		return "", fmt.Errorf("invalid initialization code: %w", err)
	}

	privateKey := airgap.KeyFromBytes(decoded)

	signingString := fmt.Sprintf("%s%s", c.HardwareId, licenseKey)

	return airgap.GenerateSignature(privateKey, signingString)
}

func (c *LicenseClient) AirgapActivation(licensePolicy, confirmationCode string, policyId int64) (LicenseResponse, error) {
	lcsJson, err := base64.StdEncoding.DecodeString(licensePolicy)
	if err != nil {
		return LicenseResponse{}, fmt.Errorf("invalid license policy: %w", err)
	}

	if confirmationCode == "" {
		return LicenseResponse{}, fmt.Errorf("invalid confirmation code")
	}

	var license LicenseResponse

	err = json.Unmarshal(lcsJson, &license)
	if err != nil {
		return LicenseResponse{}, fmt.Errorf("invalid license policy: %w", err)
	}

	if license.Id != policyId {
		return LicenseResponse{}, fmt.Errorf("invalid policy ID: %w", err)
	}

	return license, nil
}
