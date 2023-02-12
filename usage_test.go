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

func TestUsagesService_GetOrganizationUsage(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/organization_periodic_usages?order=-usage&metric[in]=cma,cpa,gql&dateRange.startAt=2020-01-01&dateRange.endAt=2020-01-03")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("usage_organization.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL
	it, err := cma.Usages.GetOrganizationUsage(context.Background(), "organization_id", "-usage", "cma,cpa,gql", "2020-01-01", "2020-01-03")
	require.NoError(t, err)
	res, err := it.Next()
	require.NoError(t, err)

	usage := res.Items
	assertions.Equal(1, len(usage))
	assertions.Equal("<usage_metric_id>", usage[0].Sys.ID)
	assertions.Equal("OrganizationPeriodicUsage", usage[0].Sys.Type)
}

func TestUsagesService_GetOrganizationUsage_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/organization_periodic_usages?order=-usage&metric[in]=cma,cpa,gql&dateRange.startAt=2020-01-01&dateRange.endAt=2020-01-03")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("usage_organization.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL
	it, err := cma.Usages.GetOrganizationUsage(context.Background(), "organization_id", "-usage", "cma,cpa,gql", "2020-01-01", "2020-01-03")
	assertions.Nil(err)
	_, err = it.Next()
	assertions.NotNil(err)
}

func TestUsagesService_GetSpaceUsage(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/space_periodic_usages?order=-usage&metric[in]=cma,cpa,gql&dateRange.startAt=2020-01-01&dateRange.endAt=2020-01-03")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("usage_space.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL
	it, err := cma.Usages.GetSpaceUsage(context.Background(), "organization_id", "-usage", "cma,cpa,gql", "2020-01-01", "2020-01-03")
	require.NoError(t, err)
	res, err := it.Next()
	require.NoError(t, err)

	usage := res.Items
	assertions.Equal(1, len(usage))
	assertions.Equal("<usage_metric_id>", usage[0].Sys.ID)
	assertions.Equal("SpacePeriodicUsage", usage[0].Sys.Type)
}

func TestUsagesService_GetSpaceUsage_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/space_periodic_usages?order=-usage&metric[in]=cma,cpa,gql&dateRange.startAt=2020-01-01&dateRange.endAt=2020-01-03")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("usage_space.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL
	it, err := cma.Usages.GetSpaceUsage(context.Background(), "organization_id", "-usage", "cma,cpa,gql", "2020-01-01", "2020-01-03")
	assertions.Nil(err)
	_, err = it.Next()
	assertions.NotNil(err)
}
