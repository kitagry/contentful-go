package contentful

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookCallsService_List(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/webhooks/0KzM2HxYr5O1pZ4SaUzK8h/calls")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("webhook_call.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	it, err := cma.WebhookCalls.List(context.Background(), spaceID, "0KzM2HxYr5O1pZ4SaUzK8h")
	assertions.Nil(err)
	collection, err := it.Next()
	assertions.Nil(err)

	spaces := collection.To()
	assertions.Equal(1, len(spaces))
	assertions.Equal("bar", spaces[0].Sys.ID)
}

func TestWebhookCallsService_Get(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/webhooks/0KzM2HxYr5O1pZ4SaUzK8h/calls/bar")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("webhook_call_detail.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	callDetails, err := cma.WebhookCalls.Get(context.Background(), spaceID, "0KzM2HxYr5O1pZ4SaUzK8h", "bar")
	assertions.Nil(err)
	assertions.Equal("bar", callDetails.Sys.ID)
	assertions.Equal("https://webhooks.example.com/endpoint", callDetails.Request.URL)
}

func TestWebhookCallsService_Get_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/webhooks/0KzM2HxYr5O1pZ4SaUzK8h/calls/bar")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("webhook_call_detail.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.WebhookCalls.Get(context.Background(), spaceID, "0KzM2HxYr5O1pZ4SaUzK8h", "bar")
	assertions.Empty(err)
}

func TestWebhookCallsService_Health(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/webhooks/0KzM2HxYr5O1pZ4SaUzK8h/health")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("webhook_health.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	health, err := cma.WebhookCalls.Health(context.Background(), spaceID, "0KzM2HxYr5O1pZ4SaUzK8h")
	assertions.Nil(err)
	assertions.Equal("bar", health.Sys.ID)
	assertions.Equal(233, health.Calls.Total)
}

func TestWebhookCallsService_Health_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/webhooks/0KzM2HxYr5O1pZ4SaUzK8h/health")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("webhook_health.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.WebhookCalls.Health(context.Background(), spaceID, "0KzM2HxYr5O1pZ4SaUzK8h")
	assertions.Nil(err)
}
