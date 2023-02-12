package contentful

import (
	"context"
	"fmt"
)

// OrganizationsService service
type OrganizationsService service

// Organization model
type Organization struct {
	Sys  *Sys   `json:"sys"`
	Name string `json:"name"`
}

// List returns an organizations collection
func (service *OrganizationsService) List(ctx context.Context, query *Query) (*Collection[Organization], error) {
	path := fmt.Sprintf("/organizations")
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col, err := newCollection[Organization](query, service.c, req)
	if err != nil {
		return nil, err
	}

	return col, nil
}
