package management_response

type SearchResult[T any] struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []T    `json:"results,omitempty"`
}
