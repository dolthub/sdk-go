package client

import "encoding/json"

type Response[T any] struct {
	Payload []byte
	Value   T
	Error   error
}

func NewResponse[T any](payload []byte, err error) Response[T] {
	if err != nil {
		return ErrorResponse[T](err)
	}

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
