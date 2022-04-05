package client

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"gitlab.com/l3178/sdk-go/v1/core/configuration"
	"net/url"
)

type Client struct {
	c *resty.Client
	*configuration.CoreConfiguration
}

func (c *Client) Get(endpoint string, pathParams map[string]string, data interface{}, response interface{}) error {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var encoded map[string]string
	err = json.Unmarshal(jsonValue, &encoded)
	if err != nil {
		return err
	}

	req := c.c.R().
		SetPathParams(pathParams).
		SetQueryParams(encoded).
		SetHeaders(c.buildHeaders())

	return c.send(endpoint, resty.MethodGet, req, response)
}

func (c *Client) Post(endpoint string, pathParams map[string]string, data interface{}, response interface{}) error {
	req := c.c.R().
		SetPathParams(pathParams).
		SetBody(data).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodPost, req, response)
}

func (c *Client) Patch(endpoint string, pathParams map[string]string, data interface{}, response interface{}) error {
	req := c.c.R().
		SetPathParams(pathParams).
		SetBody(data).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodPatch, req, response)
}

func (c *Client) Delete(endpoint string, pathParams map[string]string, response interface{}) error {
	req := c.c.R().
		SetPathParams(pathParams).
		SetHeaders(c.buildHeaders())
	return c.send(endpoint, resty.MethodDelete, req, response)
}

func (c *Client) send(url, method string, req *resty.Request, response interface{}) error {
	req.URL = url
	req.Method = method
	resp, err := req.Send()
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), response)
	if err != nil {
		return err
	}

	return nil
}

func NewClient(coreConfig *configuration.CoreConfiguration) *Client {
	return &Client{
		c:                 newResty(*coreConfig),
		CoreConfiguration: coreConfig,
	}
}

func newResty(config configuration.CoreConfiguration) *resty.Client {
	c := resty.New()
	u, _ := url.Parse(config.UrlBase)
	u, _ = u.Parse(config.EndpointPrefix)
	c.BaseURL = u.String()

	return c
}
