package management_response

type SearchResult[T any] struct {
	Count    int
	Next     string
	Previous string
	Results  T
}
