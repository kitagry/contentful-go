package contentful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LocaleItem[T any] struct {
	Item *T
	Map  map[string]T
}

func (l LocaleItem[T]) MarshalJSON() ([]byte, error) {
	if l.Item != nil {
		return json.Marshal(*l.Item)
	}
	return json.Marshal(l.Map)
}

func (l *LocaleItem[T]) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return fmt.Errorf("data is empty")
	}
	err := json.Unmarshal(b, &l.Map)
	if err == nil {
		return nil
	}

	var item T
	err = json.Unmarshal(b, &item)
	if err != nil {
		return err
	}
	l.Item = &item
	return nil
}

// CollectionOptions holds init options
type CollectionOptions struct {
	Limit uint16
}

// Collection model
type Collection[T any] struct {
	Query
	c        *Client
	req      *http.Request
	page     uint16
	Sys      *Sys        `json:"sys"`
	Total    int         `json:"total"`
	Skip     int         `json:"skip"`
	Limit    int         `json:"limit"`
	Items    []T         `json:"items"`
	Includes interface{} `json:"includes"`
}

// NewCollection initializes a new collection
func NewCollection[T any](options *CollectionOptions) *Collection[T] {
	query := NewQuery()
	query.Order("sys.createdAt", true)

	if options.Limit > 0 {
		query.Limit(options.Limit)
	}

	return &Collection[T]{
		Query: *query,
		page:  1,
	}
}

// Next makes the col.req
func (col *Collection[T]) Next() (*Collection[T], error) {
	// setup query params
	skip := uint16(col.Limit) * (col.page - 1)
	col.Query.Skip(skip)

	// override request query
	col.req.URL.RawQuery = col.Query.String()

	// makes api call
	err := col.c.do(col.req, col)
	if err != nil {
		return nil, err
	}

	col.page++

	return col, nil
}

func (col *Collection[T]) To() []T {
	return col.Items
}
