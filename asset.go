package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// AssetsService service
type AssetsService service

// Asset represents a Contentful asset
type Asset struct {
	Locale string
	Sys    *Sys         `json:"sys,omitempty"`
	Fields *AssetFields `json:"fields,omitempty"`
}

// AssetFields model
type AssetFields struct {
	Title       LocaleItem[string] `json:"title,omitempty"`
	Description LocaleItem[string] `json:"description,omitempty"`
	File        LocaleItem[File]   `json:"file,omitempty"`
}

// File represents a Contentful File
type File struct {
	URL         string       `json:"url,omitempty"`
	UploadURL   string       `json:"upload,omitempty"`
	UploadFrom  *UploadFrom  `json:"uploadFrom,omitempty"`
	Details     *FileDetails `json:"details,omitempty"`
	FileName    string       `json:"fileName,omitempty"`
	ContentType string       `json:"contentType,omitempty"`
}

// UploadFrom model
type UploadFrom struct {
	Sys *Sys `json:"sys,omitempty"`
}

// FileDetails model
type FileDetails struct {
	Size  int          `json:"size,omitempty"`
	Image *ImageFields `json:"image,omitempty"`
}

// ImageFields model
type ImageFields struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

// GetVersion returns entity version
func (asset *Asset) GetVersion() int {
	version := 1
	if asset.Sys != nil {
		version = asset.Sys.Version
	}

	return version
}

// List returns asset collection
func (service *AssetsService) List(ctx context.Context, spaceID string, query *Query) (*Collection[Asset], error) {
	path := fmt.Sprintf("/spaces/%s/assets", spaceID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col, err := newCollection[Asset](query, service.c, req)
	if err != nil {
		return nil, err
	}

	return col, nil
}

// ListPublished return a content type collection, with only activated content types
func (service *AssetsService) ListPublished(ctx context.Context, spaceID string, query *Query) (*Collection[Asset], error) {
	path := fmt.Sprintf("/spaces/%s/public/assets", spaceID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col, err := newCollection[Asset](query, service.c, req)
	if err != nil {
		return nil, err
	}

	return col, nil
}

// Get returns a single asset entity
func (service *AssetsService) Get(ctx context.Context, spaceID, assetID string) (*Asset, error) {
	path := fmt.Sprintf("/spaces/%s/assets/%s", spaceID, assetID)
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	var asset Asset
	if err := service.c.do(req, &asset); err != nil {
		return nil, err
	}

	return &asset, nil
}

// Upsert updates or creates a new asset entity
func (service *AssetsService) Upsert(ctx context.Context, spaceID string, asset *Asset) error {
	bytesArray, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	var path string
	var method string

	if asset.Sys != nil && asset.Sys.ID != "" {
		path = fmt.Sprintf("/spaces/%s/assets/%s", spaceID, asset.Sys.ID)
		method = "PUT"
	} else {
		path = fmt.Sprintf("/spaces/%s/assets", spaceID)
		method = "POST"
	}

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("X-Contentful-Version", strconv.Itoa(asset.GetVersion()))

	return service.c.do(req, asset)
}

// Delete sends delete request
func (service *AssetsService) Delete(ctx context.Context, spaceID string, asset *Asset) error {
	path := fmt.Sprintf("/spaces/%s/assets/%s", spaceID, asset.Sys.ID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(asset.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}

// Process the asset
func (service *AssetsService) Process(ctx context.Context, spaceID string, asset *Asset) error {
	path := fmt.Sprintf("/spaces/%s/assets/%s/files/%s/process", spaceID, asset.Sys.ID, asset.Locale)
	method := "PUT"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(asset.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, nil)
}

// Publish published the asset
func (service *AssetsService) Publish(ctx context.Context, spaceID string, asset *Asset) error {
	path := fmt.Sprintf("/spaces/%s/assets/%s/published", spaceID, asset.Sys.ID)
	method := "PUT"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(asset.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, asset)
}

// Unpublish the asset
func (service *AssetsService) Unpublish(ctx context.Context, spaceID string, asset *Asset) error {
	path := fmt.Sprintf("/spaces/%s/assets/%s/published", spaceID, asset.Sys.ID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(asset.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, asset)
}

// Archive archives the asset
func (service *AssetsService) Archive(ctx context.Context, spaceID string, asset *Asset) error {
	path := fmt.Sprintf("/spaces/%s/assets/%s/archived", spaceID, asset.Sys.ID)
	method := "PUT"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(asset.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, asset)
}

// Unarchive unarchives the asset
func (service *AssetsService) Unarchive(ctx context.Context, spaceID string, asset *Asset) error {
	path := fmt.Sprintf("/spaces/%s/assets/%s/archived", spaceID, asset.Sys.ID)
	method := "DELETE"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return err
	}

	version := strconv.Itoa(asset.Sys.Version)
	req.Header.Set("X-Contentful-Version", version)

	return service.c.do(req, asset)
}
