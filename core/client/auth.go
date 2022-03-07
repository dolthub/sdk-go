package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

const (
	signaturePrefix = "licenseSpring"
	headers         = "date"
	algorithm       = "hmac-sha256"
)

func (c *Client) createAuthHeader(date time.Time) string {
	signature := c.createSignature(date)

	return fmt.Sprintf("algorithm=\"%s\",headers=\"%s\",signature=\"%s\",apikey=\"%s\"",
		algorithm, headers, signature, c.ApiKey)
}

func (c *Client) createSignature(date time.Time) string {
	dateStr := date.Format(dateFormat)
	signatureString := fmt.Sprintf("%s\ndate: %s", signaturePrefix, dateStr)
	return c.encode(signatureString)
}

func (c *Client) encode(signature string) string {
	encoder := hmac.New(sha256.New, []byte(c.SharedKey))
	encoder.Write([]byte(signature))
	return base64.StdEncoding.EncodeToString(encoder.Sum(nil))
}
