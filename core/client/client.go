package client

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gitlab.com/l3178/sdk-go/core/configuration"
	"gitlab.com/l3178/sdk-go/core/helpers"
	"time"
)

type Client struct {
	c      *resty.Client
	config configuration.Config
}

func (c *Client) Get(ctx context.Context, endpoint string, pathParams map[string]string, data interface{}) ([]byte, error) {
	values, err := helpers.Values(data)
	if err != nil {
		return nil, err
	}

	req := c.c.R().
		SetContext(ctx).
		SetPathParams(pathParams).
		SetQueryParamsFromValues(values).
		SetHeaders(c.buildHeaders())

	return c.send(endpoint, resty.MethodGet, req)
}

func (c *Client) Post(ctx context.Context, endpoint string, pathParams map[string]string, data interface{}) ([]byte, error) {
	req := c.c.R().
		SetContext(ctx).
		SetPathParams(pathParams).
		SetBody(data).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodPost, req)
}

func (c *Client) Patch(ctx context.Context, endpoint string, pathParams map[string]string, data interface{}) ([]byte, error) {
	req := c.c.R().
		SetContext(ctx).
		SetPathParams(pathParams).
		SetBody(data).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodPatch, req)
}

func (c *Client) Delete(ctx context.Context, endpoint string, pathParams map[string]string) ([]byte, error) {
	req := c.c.R().
		SetContext(ctx).
		SetPathParams(pathParams).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodDelete, req)
}

func (c *Client) send(endpoint, method string, req *resty.Request) ([]byte, error) {
	req.URL = c.buildUrl(endpoint)
	req.Method = method
	resp, err := req.Send()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if resp.StatusCode() > 299 {
		m := make(map[string]interface{})
		jsonErr := json.Unmarshal(resp.Body(), &m)
		if jsonErr != nil {
			return nil, errors.New("unable to unmarshal JSON")
		}
		return nil, errors.New(m["message"].(string))
	}

	return resp.Body(), nil
}

func (c *Client) buildUrl(endpoint string) string {
	return c.config.UrlPrefix() + endpoint
}

func NewClient(config configuration.Config, verbose bool) *Client {
	requestConfig := config.GetRequestConfig()
	c := resty.New()
	c.Debug = verbose
	c.SetRetryCount(requestConfig.RetryCount)
	c.SetTimeout(time.Duration(requestConfig.RequestTimeout) * time.Second)
	return &Client{
		c:      c,
		config: config,
	}
}
