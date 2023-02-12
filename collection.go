package contentful

import (
	"bytes"
	"context"
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
	query    *Query
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

// newCollection initializes a new collection
// if query is nil, order sys.createdAt
func newCollection[T any](query *Query, client *Client, req *http.Request) (*Collection[T], error) {
	if query == nil {
		query = NewQuery()
		query.Order("sys.createdAt", true)
	}
	col := &Collection[T]{
		query: query,
		c:     client,
		req:   req,
		page:  1,
	}

	return col.Next(req.Context())
}

// Next makes the col.req
func (col *Collection[T]) Next(ctx context.Context) (*Collection[T], error) {
	// setup query params
	skip := uint16(col.Limit) * (col.page - 1)
	col.query.Skip(skip)

	// override request query
	col.req.URL.RawQuery = col.query.String()

	// makes api call
	err := col.c.do(col.req, col)
	if err != nil {
		return nil, err
	}

	col.page++

	return col, nil
}
