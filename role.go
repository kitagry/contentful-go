package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// RolesService service
type RolesService service

// Role model
type Role struct {
	Sys         *Sys        `json:"sys"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Policies    []Policy    `json:"policies"`
	Permissions Permissions `json:"permissions"`
}

// Policy model
type Policy struct {
	Effect     string     `json:"effect"`
	Actions    any        `json:"actions"`
	Constraint Constraint `json:"constraint"`
}

// Permissions model
type Permissions struct {
	ContentModel       []string `json:"ContentModel"`
	Settings           any      `json:"Settings"`
	ContentDelivery    any      `json:"ContentDelivery"`
	Environments       any      `json:"Environments"`
	EnvironmentAliases any      `json:"EnvironmentAliases"`
}

// Constraint model
type Constraint struct {
	And   any `json:"and"`
	Equal any `json:"equal"`
	Or    any `json:"or"`
	Not   any `rson:"not"`
}

// GetVersion returns entity version
func (r *Role) GetVersion() int {
	version := 1
	if r.Sys != nil {
		version = r.Sys.Version
	}

	return version
}

// List returns an environments collection
func (service *RolesService) List(ctx context.Context, spaceID string) (*Collection[Role], error) {
	path := fmt.Sprintf("/spaces/%s/roles", spaceID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col := NewCollection[Role](&CollectionOptions{})
	col.c = service.c
	col.req = req

	return col, nil
}

// Get returns a single role
func (service *RolesService) Get(ctx context.Context, spaceID, roleID string) (*Role, error) {
	path := fmt.Sprintf("/spaces/%s/roles/%s", spaceID, roleID)
	query := url.Values{}
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, query, nil)
	if err != nil {
		return &Role{}, err
	}

	var role Role
	if ok := service.c.do(req, &role); ok != nil {
		return nil, err
	}

	return &role, err
}

// Upsert updates or creates a new role
func (service *RolesService) Upsert(ctx context.Context, spaceID string, r *Role) error {
	bytesArray, err := json.Marshal(r)
	if err != nil {
		return err
	}

	var path string
	var method string

	if r.Sys != nil && r.Sys.ID != "" {
		path = fmt.Sprintf("/spaces/%s/roles/%s", spaceID, r.Sys.ID)
		method = "PUT"
	} else {
		path = fmt.Sprintf("/spaces/%s/roles", spaceID)
		method = "POST"
	}

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("X-Contentful-Version", strconv.Itoa(r.GetVersion()))

	return service.c.do(req, r)
}

// Delete the role
func (service *RolesService) Delete(ctx context.Context, spaceID string, roleID string) error {
	path := fmt.Sprintf("/spaces/%s/roles/%s", spaceID, roleID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.c.do(req, nil)
}
