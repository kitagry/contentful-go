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
func (service *OrganizationsService) List(ctx context.Context) (*Collection[Organization], error) {
	path := fmt.Sprintf("/organizations")
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col := NewCollection[Organization](&CollectionOptions{})
	col.c = service.c
	col.req = req

	return col, nil
}
