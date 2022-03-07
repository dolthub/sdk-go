package client

import (
	"time"
)

const dateFormat = time.RFC1123

const (
	DateHeader = "Date"
	AuthHeader = "Authorization"
)

func (c *Client) buildHeaders() map[string]string {
	date := time.Now()

	headers := map[string]string{
		DateHeader:     time.Now().Format(dateFormat),
		AuthHeader:     c.createAuthHeader(date),
		"Content-Type": "application/json",
	}
	for key, value := range c.AdditionalHeaders {
		headers[key] = value
	}
	return headers
}
