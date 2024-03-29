package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// AccessTokensService service
type AccessTokensService service

// AccessToken model
type AccessToken struct {
	Sys       *Sys     `json:"sys,omitempty"`
	Name      string   `json:"name,omitempty"`
	RevokedAt string   `json:"revokedAt,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`
}

// GetVersion returns entity version
func (accessToken *AccessToken) GetVersion() int {
	version := 1
	if accessToken.Sys != nil {
		version = accessToken.Sys.Version
	}

	return version
}

// List returns an access tokens collection
func (service *AccessTokensService) List(ctx context.Context, query *Query) (*Collection[AccessToken], error) {
	path := fmt.Sprint("/users/me/access_tokens")
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, nil, nil)
	if err != nil {
		return nil, err
	}

	col, err := newCollection[AccessToken](query, service.c, req)
	if err != nil {
		return nil, err
	}

	return col, nil
}

// Get returns a single access token
func (service *AccessTokensService) Get(ctx context.Context, accessTokenID string) (*AccessToken, error) {
	path := fmt.Sprintf("/users/me/access_tokens/%s", accessTokenID)
	query := url.Values{}
	method := "GET"

	req, err := service.c.newRequest(ctx, method, path, query, nil)
	if err != nil {
		return &AccessToken{}, err
	}

	var accessToken AccessToken
	req.Header.Set("X-Contentful-Version", strconv.Itoa(accessToken.GetVersion()))
	if ok := service.c.do(req, &accessToken); ok != nil {
		return nil, err
	}

	return &accessToken, err
}

// Create creates a new access token
func (service *AccessTokensService) Create(ctx context.Context, accessToken *AccessToken) error {
	bytesArray, err := json.Marshal(accessToken)
	if err != nil {
		return err
	}

	path := fmt.Sprint("/users/me/access_tokens")
	method := "POST"
	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	req.Header.Set("X-Contentful-Version", strconv.Itoa(accessToken.GetVersion()))
	return service.c.do(req, accessToken)
}

// Revoke revokes a personal access token
func (service *AccessTokensService) Revoke(ctx context.Context, accessToken *AccessToken) error {
	bytesArray, err := json.Marshal(accessToken)
	if err != nil {
		return err
	}

	var path string
	var method string

	if accessToken.Sys != nil && accessToken.Sys.CreatedAt != "" {
		path = fmt.Sprintf("/users/me/access_tokens/%s/revoked", accessToken.Sys.ID)
		method = "PUT"
	}

	req, err := service.c.newRequest(ctx, method, path, nil, bytes.NewReader(bytesArray))
	if err != nil {
		return err
	}

	return service.c.do(req, accessToken)
}
