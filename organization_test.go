package contentful

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganizationsService_List(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("organization.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.Organizations.List(context.Background(), nil)
	require.NoError(t, err)
	organization := collection.Items
	assertions.Equal(1, len(organization))
	assertions.Equal("My organization", organization[0].Name)
}
