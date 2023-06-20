package client

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Response[T any] struct {
	Payload []byte
	Value   T
	Error   error
}

func NewResponse[T any](res *resty.Response, err error) Response[T] {
	if err != nil {
		return ErrorResponse[T](err)
	}

	payload := res.Body()

	var resp Response[T]
	resp.Payload = payload

	err = json.Unmarshal(payload, &resp.Value)
	if err != nil {
		resp.Error = err
	}

	return resp
}

func ErrorResponse[T any](err error) Response[T] {
	return Response[T]{
		Error: err,
	}
}
