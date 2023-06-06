package client

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gitlab.com/l3178/sdk-go/core/configuration"
	"gitlab.com/l3178/sdk-go/core/helpers"
	"log"
	"time"
)

type Client struct {
	c      *resty.Client
	config configuration.Config
}

func (c *Client) Get(endpoint string, pathParams map[string]string, data interface{}) ([]byte, error) {
	values, err := helpers.Values(data)
	if err != nil {
		return nil, err
	}

	req := c.c.R().
		SetPathParams(pathParams).
		SetQueryParamsFromValues(values).
		SetHeaders(c.buildHeaders())

	return c.send(endpoint, resty.MethodGet, req)
}

func (c *Client) Post(endpoint string, pathParams map[string]string, data interface{}) ([]byte, error) {
	req := c.c.R().
		SetPathParams(pathParams).
		SetBody(data).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodPost, req)
}

func (c *Client) Patch(endpoint string, pathParams map[string]string, data interface{}) ([]byte, error) {
	req := c.c.R().
		SetPathParams(pathParams).
		SetBody(data).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodPatch, req)
}

func (c *Client) Delete(endpoint string, pathParams map[string]string) ([]byte, error) {
	req := c.c.R().
		SetPathParams(pathParams).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodDelete, req)
}

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) send(endpoint, method string, req *resty.Request) ([]byte, error) {
	req.URL = c.buildUrl(endpoint)
	req.Method = method
	resp, err := req.Send()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if resp.StatusCode() > 299 {
		var errResp errorResponse
		err = json.Unmarshal(resp.Body(), &errResp)
		if err != nil {
			log.Panic(err)
		}
		return nil, errors.New(errResp.Message)
	}

	if !isJSON(resp.Body()) {
		return nil, nil
	}

	return resp.Body(), nil
}

func isJSON(str []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(str, &js) == nil
}

func (c *Client) buildUrl(endpoint string) string {
	return c.config.UrlPrefix() + endpoint
}

func NewClient(config configuration.Config) *Client {
	requestConfig := config.GetRequestConfig()
	c := resty.New()
	c.Debug = true
	c.SetRetryCount(requestConfig.RetryCount)
	c.SetTimeout(time.Duration(requestConfig.RequestTimeout) * time.Second)
	return &Client{
		c:      c,
		config: config,
	}
}
