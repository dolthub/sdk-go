package client

import (
	"fmt"
	"time"
)

const dateFormat = time.RFC1123

const (
	DateHeader = "Date"
	AuthHeader = "Authorization"
)

func (c *Client) buildHeaders() map[string]string {
	date := time.Now()
	fmt.Println("date", date.String())
	headers := map[string]string{
		DateHeader:     date.Format(dateFormat),
		AuthHeader:     c.config.AuthHeader(date),
		"Content-Type": "application/json",
	}
	return headers
}
