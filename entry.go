package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// EntriesService service
type EntriesService service

// Entry model
type Entry struct {
	Locale string         `json:"locale"`
	Sys    *Sys           `json:"sys"`
	Fields map[string]any `json:"fields"`
}

// GetVersion returns entity version
func (entry *Entry) GetVersion() int {
	version := 1
	if entry.Sys != nil {
		version = entry.Sys.Version
	}

	return version
}

// List returns entries collection
func (service *EntriesService) List(ctx context.Context, env *Environment) (*Collection[Entry], error) {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries", env.Sys.Space.Sys.ID, env.Sys.ID)

	req, err := service.c.newRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col := NewCollection[Entry](&CollectionOptions{})
	col.c = service.c
	col.req = req

	return col, nil
}

// Get returns a single entry
func (service *EntriesService) Get(ctx context.Context, env *Environment, entryID string) (*Entry, error) {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s", env.Sys.Space.Sys.ID, env.Sys.ID, entryID)
	query := url.Values{}
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, query, nil)
	if err != nil {
		return &Entry{}, err
	}

	var entry Entry
	if ok := service.c.do(req, &entry); ok != nil {
		return nil, err
	}

	return &entry, err
}

// Delete the entry
func (service *EntriesService) Delete(ctx context.Context, env *Environment, entryID string) error {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s", env.Sys.Space.Sys.ID, env.Sys.ID, entryID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.c.do(req, nil)
}

// Publish the entry
func (service *EntriesService) Publish(ctx context.Context, env *Environment, entry *Entry) error {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s/published", env.Sys.Space.Sys.ID, env.Sys.ID, entry.Sys.ID)
	method := "PUT"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(entry.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}

// Unpublish the entry
func (service *EntriesService) Unpublish(ctx context.Context, env *Environment, entry *Entry) error {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s/published", env.Sys.Space.Sys.ID, env.Sys.ID, entry.Sys.ID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(entry.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}

// Upsert updates or creates a new entry
func (service *EntriesService) Upsert(ctx context.Context, env *Environment, contentTypeID string, e *Entry) error {
	bytesArray, err := json.Marshal(e)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/spaces/%s/environments/%s", env.Sys.Space.Sys.ID, env.Sys.ID)
	var method string

	if e.Sys != nil && e.Sys.ID != "" {
		path += fmt.Sprintf("/entries/%s", e.Sys.ID)
		method = "PUT"
	} else {
		path += "/entries"
		method = "POST"
	}

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("X-Contentful-Content-Type", contentTypeID)
	req.Header.Set("X-Contentful-Version", strconv.Itoa(e.GetVersion()))

	return service.c.do(req, e)
}

// Archive the entry
func (service *EntriesService) Archive(ctx context.Context, env *Environment, entry *Entry) error {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s/archived", env.Sys.Space.Sys.ID, env.Sys.ID, entry.Sys.ID)
	method := "PUT"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(entry.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}

// Unarchive the entry
func (service *EntriesService) Unarchive(ctx context.Context, env *Environment, entry *Entry) error {
	path := fmt.Sprintf("/spaces/%s/environments/%s/entries/%s/archived", env.Sys.Space.Sys.ID, env.Sys.ID, entry.Sys.ID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(entry.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}
