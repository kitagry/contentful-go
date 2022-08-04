package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// EnvironmentAliasesService service
type EnvironmentAliasesService service

// EnvironmentAlias model
type EnvironmentAlias struct {
	Sys   *Sys         `json:"sys"`
	Alias *AliasDetail `json:"environment"`
}

// AliasDetail model
type AliasDetail struct {
	Sys *Sys `json:"sys"`
}

// GetVersion returns entity version
func (environmentAlias *EnvironmentAlias) GetVersion() int {
	version := 1
	if environmentAlias.Sys != nil {
		version = environmentAlias.Sys.Version
	}

	return version
}

// List returns an environment aliases collection
func (service *EnvironmentAliasesService) List(ctx context.Context, spaceID string) *Collection {
	path := fmt.Sprintf("/spaces/%s/environment_aliases", spaceID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil
	}

	col := NewCollection(&CollectionOptions{})
	col.c = service.c
	col.req = req

	return col
}

// Get returns a single environment alias entity
func (service *EnvironmentAliasesService) Get(ctx context.Context, spaceID, environmentAliasID string) (*EnvironmentAlias, error) {
	path := fmt.Sprintf("/spaces/%s/environment_aliases/%s", spaceID, environmentAliasID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var environmentAlias EnvironmentAlias
	if err := service.c.do(req, &environmentAlias); err != nil {
		return nil, err
	}

	return &environmentAlias, nil
}

// Update updates an environment alias
func (service *EnvironmentAliasesService) Update(ctx context.Context, spaceID string, ea *EnvironmentAlias) error {
	bytesArray, err := json.Marshal(ea)
	if err != nil {
		return err
	}

	var path string
	var method string

	if ea.Sys != nil && ea.Sys.ID != "" {
		path = fmt.Sprintf("/spaces/%s/environment_aliases/%s", spaceID, ea.Sys.ID)
		method = "PUT"
	}

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("X-Contentful-Version", strconv.Itoa(ea.GetVersion()))

	return service.c.do(req, ea)
}
