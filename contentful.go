package contentful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"moul.io/http2curl"
)

// Client model
type Client struct {
	client        *http.Client
	api           string
	token         string
	Debug         bool
	QueryParams   map[string]string
	Headers       map[string]string
	BaseURL       string
	Environment   string
	commonService service

	Spaces             *SpacesService
	Users              *UsersService
	Environments       *EnvironmentsService
	EnvironmentAliases *EnvironmentAliasesService
	Organizations      *OrganizationsService
	Roles              *RolesService
	Memberships        *MembershipsService
	Snapshots          *SnapshotsService
	APIKeys            *APIKeyService
	AccessTokens       *AccessTokensService
	Assets             *AssetsService
	ContentTypes       *ContentTypesService
	Entries            *EntriesService
	EntryTasks         *EntryTasksService
	ScheduledActions   *ScheduledActionsService
	Locales            *LocalesService
	Webhooks           *WebhooksService
	WebhookCalls       *WebhookCallsService
	EditorInterfaces   *EditorInterfacesService
	Extensions         *ExtensionsService
	AppDefinitions     *AppDefinitionsService
	AppInstallations   *AppInstallationsService
	Usages             *UsagesService
	Resources          *ResourcesService
}

type service struct {
	c *Client
}

// NewCMA returns a CMA client
func NewCMA(token string) *Client {
	c := &Client{
		client: http.DefaultClient,
		api:    "CMA",
		token:  token,
		Debug:  false,
		Headers: map[string]string{
			"Authorization":           fmt.Sprintf("Bearer %s", token),
			"Content-Type":            "application/vnd.contentful.management.v1+json",
			"X-Contentful-User-Agent": fmt.Sprintf("sdk contentful.go/%s", Version),
		},
		BaseURL:     "https://api.contentful.com",
		Environment: "master",
	}
	c.commonService.c = c

	c.Spaces = (*SpacesService)(&c.commonService)
	c.Users = (*UsersService)(&c.commonService)
	c.Environments = (*EnvironmentsService)(&c.commonService)
	c.EnvironmentAliases = (*EnvironmentAliasesService)(&c.commonService)
	c.Organizations = (*OrganizationsService)(&c.commonService)
	c.Roles = (*RolesService)(&c.commonService)
	c.Memberships = (*MembershipsService)(&c.commonService)
	c.Snapshots = (*SnapshotsService)(&c.commonService)
	c.APIKeys = (*APIKeyService)(&c.commonService)
	c.AccessTokens = (*AccessTokensService)(&c.commonService)
	c.Assets = (*AssetsService)(&c.commonService)
	c.ContentTypes = (*ContentTypesService)(&c.commonService)
	c.Entries = (*EntriesService)(&c.commonService)
	c.EntryTasks = (*EntryTasksService)(&c.commonService)
	c.ScheduledActions = (*ScheduledActionsService)(&c.commonService)
	c.Locales = (*LocalesService)(&c.commonService)
	c.Webhooks = (*WebhooksService)(&c.commonService)
	c.WebhookCalls = (*WebhookCallsService)(&c.commonService)
	c.EditorInterfaces = (*EditorInterfacesService)(&c.commonService)
	c.Extensions = (*ExtensionsService)(&c.commonService)
	c.AppDefinitions = (*AppDefinitionsService)(&c.commonService)
	c.AppInstallations = (*AppInstallationsService)(&c.commonService)
	c.Usages = (*UsagesService)(&c.commonService)
	return c
}

// NewCDA returns a CDA client
func NewCDA(token string) *Client {
	c := &Client{
		client: http.DefaultClient,
		api:    "CDA",
		token:  token,
		Debug:  false,
		Headers: map[string]string{
			"Authorization":           "Bearer " + token,
			"Content-Type":            "application/vnd.contentful.delivery.v1+json",
			"X-Contentful-User-Agent": fmt.Sprintf("contentful-go/%s", Version),
		},
		BaseURL:     "https://cdn.contentful.com",
		Environment: "master",
	}
	c.commonService.c = c

	c.Spaces = (*SpacesService)(&c.commonService)
	c.APIKeys = (*APIKeyService)(&c.commonService)
	c.Assets = (*AssetsService)(&c.commonService)
	c.ContentTypes = (*ContentTypesService)(&c.commonService)
	c.Entries = (*EntriesService)(&c.commonService)
	c.Locales = (*LocalesService)(&c.commonService)
	c.Webhooks = (*WebhooksService)(&c.commonService)

	return c
}

// NewCPA returns a CPA client
func NewCPA(token string) *Client {
	c := &Client{
		client: http.DefaultClient,
		Debug:  false,
		api:    "CPA",
		token:  token,
		Headers: map[string]string{
			"Authorization": "Bearer " + token,
		},
		BaseURL: "https://preview.contentful.com",
	}

	c.Spaces = &SpacesService{c: c}
	c.APIKeys = &APIKeyService{c: c}
	c.Assets = &AssetsService{c: c}
	c.ContentTypes = &ContentTypesService{c: c}
	c.Entries = &EntriesService{c: c}
	c.Locales = &LocalesService{c: c}
	c.Webhooks = &WebhooksService{c: c}

	return c
}

// NewResourceClient returns a client for the resource/uploads endpoints
func NewResourceClient(token string) *Client {
	c := &Client{
		client: http.DefaultClient,
		api:    "URC",
		Debug:  false,
		token:  token,
		Headers: map[string]string{
			"Authorization": "Bearer " + token,
		},
		BaseURL: "https://upload.contentful.com",
	}
	c.commonService.c = c

	c.Resources = (*ResourcesService)(&c.commonService)

	return c
}

// SetOrganization sets the given organization id
func (c *Client) SetOrganization(organizationID string) *Client {
	c.Headers["X-Contentful-Organization"] = organizationID

	return c
}

// SetEnvironment sets the given environment.
// https://www.contentful.com/developers/docs/references/content-management-api/#/reference/environments
func (c *Client) SetEnvironment(environment string) *Client {
	c.Environment = environment
	return c
}

// SetHTTPClient sets the underlying http.Client used to make requests.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.client = client
}

func (c *Client) newRequest(ctx context.Context, method, path string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	// set query params
	for key, value := range c.QueryParams {
		query.Set(key, value)
	}

	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	if c.Debug {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 400 {
		// Upload/Create Resource response cannot be decoded
		if !(c.api == "URC" && req.Method == "POST") {
			if v != nil {
				b, err := io.ReadAll(res.Body)
				if err != nil {
					return err
				}
				err = json.NewDecoder(bytes.NewReader(b)).Decode(v)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	// parse api response
	apiError := c.handleError(req, res)

	// return apiError if it is not rate limit error
	if _, ok := apiError.(RateLimitExceededError); !ok {
		return apiError
	}

	resetHeader := res.Header.Get("x-contentful-ratelimit-reset")

	// return apiError if Ratelimit-Reset header is not presented
	if resetHeader == "" {
		return apiError
	}

	// wait X-Contentful-Ratelimit-Reset amount of seconds
	waitSeconds, err := strconv.Atoi(resetHeader)
	if err != nil {
		return apiError
	}

	time.Sleep(time.Second * time.Duration(waitSeconds))

	return c.do(req, v)
}

func (c *Client) handleError(req *http.Request, res *http.Response) error {
	if c.Debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%q", dump)
	}

	var e ErrorResponse
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&e)
	if err != nil {
		return err
	}

	apiError := APIError{
		req: req,
		res: res,
		err: &e,
	}

	switch errType := e.Sys.ID; errType {
	case "NotFound":
		return NotFoundError{apiError}
	case "RateLimitExceeded":
		return RateLimitExceededError{apiError}
	case "AccessTokenInvalid":
		return AccessTokenInvalidError{apiError}
	case "ValidationFailed":
		return ValidationFailedError{apiError}
	case "VersionMismatch":
		return VersionMismatchError{apiError}
	case "Conflict":
		return VersionMismatchError{apiError}
	default:
		return e
	}
}
