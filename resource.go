package contentful

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
)

// ResourcesService service
type ResourcesService service

// Resource model
type Resource struct {
	Sys *Sys `json:"sys"`
}

// Get returns a single resource/upload
func (service *ResourcesService) Get(ctx context.Context, spaceID, resourceID string) (*Resource, error) {
	path := fmt.Sprintf("/spaces/%s/uploads/%s", spaceID, resourceID)
	query := url.Values{}
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, query, nil)
	if err != nil {
		return &Resource{}, err
	}

	var resource Resource
	if ok := service.c.do(req, &resource); ok != nil {
		return nil, err
	}

	return &resource, err
}

// Create creates an upload resource
func (service *ResourcesService) Create(ctx context.Context, spaceID, filePath string) error {
	bytesArray, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("/spaces/%s/uploads", spaceID)
	method := "POST"

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")

	return service.c.do(req, bytesArray)
}

// Delete the resource
func (service *ResourcesService) Delete(ctx context.Context, spaceID, resourceID string) error {
	path := fmt.Sprintf("/spaces/%s/uploads/%s", spaceID, resourceID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	return service.c.do(req, nil)
}
