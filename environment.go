package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// EnvironmentsService service
type EnvironmentsService service

// Environment model
type Environment struct {
	Sys  *Sys   `json:"sys"`
	Name string `json:"name"`
}

// GetVersion returns entity version
func (e *Environment) GetVersion() int {
	version := 1
	if e.Sys != nil {
		version = e.Sys.Version
	}

	return version
}

// List returns an environments collection
func (service *EnvironmentsService) List(ctx context.Context, spaceID string, query *Query) (*Collection[Environment], error) {
	path := fmt.Sprintf("/spaces/%s/environments", spaceID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col, err := newCollection[Environment](query, service.c, req)
	if err != nil {
		return nil, err
	}

	return col, nil
}

// Get returns a single environment entity
func (service *EnvironmentsService) Get(ctx context.Context, spaceID, environmentID string) (*Environment, error) {
	path := fmt.Sprintf("/spaces/%s/environments/%s", spaceID, environmentID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var environment Environment
	if err := service.c.do(req, &environment); err != nil {
		return nil, err
	}

	return &environment, nil
}

// Upsert updates or creates a new environment
func (service *EnvironmentsService) Upsert(ctx context.Context, spaceID string, e *Environment) error {
	bytesArray, err := json.Marshal(e)
	if err != nil {
		return err
	}

	var path string
	var method string

	if e.Sys != nil && e.Sys.ID != "" {
		path = fmt.Sprintf("/spaces/%s/environments/%s", spaceID, e.Sys.ID)
		method = "PUT"
	} else {
		path = fmt.Sprintf("/spaces/%s/environments/%s", spaceID, e.Name)
		method = "PUT"
	}

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("X-Contentful-Version", strconv.Itoa(e.GetVersion()))

	return service.c.do(req, e)
}

// Delete the environment
func (service *EnvironmentsService) Delete(ctx context.Context, spaceID string, e *Environment) error {
	path := fmt.Sprintf("/spaces/%s/environments/%s", spaceID, e.Sys.ID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(e.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}
