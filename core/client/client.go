package client

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"gitlab.com/l3178/sdk-go/core/configuration"
	"log"
	"net/http"
	"time"
)

type Client struct {
	c      *resty.Client
	config configuration.Config
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

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) send(endpoint, method string, req *resty.Request, response interface{}) error {
	req.URL = c.buildUrl(endpoint)
	req.Method = method
	resp, err := req.Send()
	if err != nil {
		return errors.New(err.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		var errResp errorResponse
		err = json.Unmarshal(resp.Body(), &errResp)
		if err != nil {
			log.Panic(err)
		}
		return errors.New(errResp.Message)
	}

	if !isJSON(resp.Body()) {
		return nil
	}

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
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
	c.SetRetryCount(requestConfig.RetryCount)
	c.SetTimeout(time.Duration(requestConfig.RequestTimeout) * time.Second)
	return &Client{
		c:      c,
		config: config,
	}
}
