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

func TestAppDefinitionsService_List(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/app_definitions")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("app_definition.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.AppDefinitions.List(context.Background(), "organization_id", nil)
	assertions.Nil(err)

	definitions := collection.Items
	assertions.Equal(1, len(definitions))
	assertions.Equal("app_definition_id", definitions[0].Sys.ID)
	assertions.Equal("https://example.com/app.html", definitions[0].SRC)
}

func TestAppDefinitionsService_Get(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/app_definitions/app_definition_id")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("app_definition_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	definition, err := cma.AppDefinitions.Get(context.Background(), "organization_id", "app_definition_id")
	assertions.Nil(err)
	assertions.Equal("app_definition_id", definition.Sys.ID)
	assertions.Equal("Hello world!", definition.Name)
}

func TestAppDefinitionsService_Get_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/organizations/organization_id/app_definitions/app_definition_id")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("app_definition_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.AppDefinitions.Get(context.Background(), "organization_id", "app_definition_id")
	assertions.Nil(err)
}

func TestAppDefinitionsService_Upsert_Create(t *testing.T) {
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "POST")
		assertions.Equal(r.RequestURI, "/organizations/organization_id/app_definitions")
		checkHeaders(r, assertions)

		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)
		assertions.Equal("Hello world!", payload["name"])
		assertions.Equal("https://example.com/app.html", payload["src"])

		w.WriteHeader(201)
		_, _ = fmt.Fprintln(w, readTestData("app_definition_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	definition := &AppDefinition{
		Name: "Hello world!",
		SRC:  "https://example.com/app.html",
		Locations: []Locations{
			{
				Location: "entry-sidebar",
			},
		},
	}

	err := cma.AppDefinitions.Upsert(context.Background(), "organization_id", definition)
	assertions.Nil(err)
	assertions.Equal("app_definition_id", definition.Sys.ID)
	assertions.Equal("Hello world!", definition.Name)
}

func TestAppDefinitionsService_Upsert_Update(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "PUT")
		assertions.Equal(r.RequestURI, "/organizations/organization_id/app_definitions/app_definition_id")
		checkHeaders(r, assertions)

		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)
		assertions.Equal("Hello Pluto", payload["name"])
		assertions.Equal("https://example.com/hellopluto.html", payload["src"])

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("app_definition_updated.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	definition, err := appDefinitionFromTestFile("app_definition_1.json")
	assertions.Nil(err)

	definition.Name = "Hello Pluto"
	definition.SRC = "https://example.com/hellopluto.html"

	err = cma.AppDefinitions.Upsert(context.Background(), "organization_id", definition)
	assertions.Nil(err)
	assertions.Equal("Hello Pluto", definition.Name)
	assertions.Equal("https://example.com/hellopluto.html", definition.SRC)
}

func TestAppDefinitionsService_Delete(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "DELETE")
		assertions.Equal(r.RequestURI, "/organizations/organization_id/app_definitions/app_definition_id")
		checkHeaders(r, assertions)

		w.WriteHeader(200)
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	definition, err := appDefinitionFromTestFile("app_definition_1.json")
	assertions.Nil(err)

	err = cma.AppDefinitions.Delete(context.Background(), "organization_id", definition.Sys.ID)
	assertions.Nil(err)
}
