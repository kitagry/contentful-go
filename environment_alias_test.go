package contentful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentAliasesServices_List(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environment_aliases")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("environment_alias.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.EnvironmentAliases.List(context.Background(), spaceID, nil)
	assertions.Nil(err)
	environmentAlias := collection.Items
	assertions.Equal(1, len(environmentAlias))
	assertions.Equal("master-18-3-2020", environmentAlias[0].Alias.Sys.ID)
}

func TestEnvironmentAliasesServices_Get(t *testing.T) {
	var err error
	assertions := assert.New(t)

	// Only tests master environment, as this is the only environment that always exists.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environment_aliases/master")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("environment_alias_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	environmentAlias, err := cma.EnvironmentAliases.Get(context.Background(), spaceID, "master")
	assertions.Nil(err)
	assertions.Equal("master-18-3-2020", environmentAlias.Alias.Sys.ID)
}

func TestEnvironmentAliasesServices_Get_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	// Only tests master environment, as this is the only environment that always exists.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environment_aliases/master")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("environment_alias_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.EnvironmentAliases.Get(context.Background(), spaceID, "master")
	assertions.NotNil(err)
}

func TestEnvironmentAliasesService_Update(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "PUT")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/environment_aliases/master")

		checkHeaders(r, assertions)

		var payload EnvironmentAlias
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)
		assertions.Equal("staging", payload.Alias.Sys.ID)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("environment_alias_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	environmentAlias, err := environmentAliasFromTestData("environment_alias_1.json")
	assertions.Nil(err)

	environmentAlias.Alias.Sys.ID = "staging"

	err = cma.EnvironmentAliases.Update(context.Background(), spaceID, environmentAlias)
	assertions.Nil(err)
}
