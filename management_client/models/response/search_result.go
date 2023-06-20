package management_response

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type SearchResult[T any] struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []T    `json:"results,omitempty"`
	Error    error
}

func SearchResultFromJson[T any](resp *resty.Response, err error) SearchResult[T] {
	if err != nil {
		return errorSearchResult[T](err)
	}
	res := SearchResult[T]{}
	unmarshalError := json.Unmarshal(resp.Body(), &res)
	if unmarshalError != nil {
		return errorSearchResult[T](err)
	}
	return res
}

func errorSearchResult[T any](err error) SearchResult[T] {
	return SearchResult[T]{
		Error: err,
	}
}
