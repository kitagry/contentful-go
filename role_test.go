package contentful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRolesService_List(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/roles")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("role.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.Roles.List(context.Background(), spaceID, nil)
	require.NoError(t, err)
	role := collection.Items
	assertions.Equal(2, len(role))
	assertions.Equal("Author", role[0].Name)
}

func TestRolesService_Get(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/roles/0xvkNW6WdQ8JkWlWZ8BC4x")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("role_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	role, err := cma.Roles.Get(context.Background(), spaceID, "0xvkNW6WdQ8JkWlWZ8BC4x")
	assertions.Nil(err)
	assertions.Equal("Some role", role.Name)
}

func TestRolesService_Get_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/roles/0xvkNW6WdQ8JkWlWZ8BC4x")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("role_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.Roles.Get(context.Background(), spaceID, "0xvkNW6WdQ8JkWlWZ8BC4x")
	assertions.Nil(err)
}

func TestRolesService_Upsert_Create(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "POST")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/roles")

		checkHeaders(r, assertions)

		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)

		name := payload["name"]
		description := payload["description"]
		assertions.Equal("Author", name)
		assertions.Equal("Describes the author", description)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("role_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	role := &Role{
		Name:        "Author",
		Description: "Describes the author",
		Policies: []Policy{
			{
				Effect: "allow",
				Actions: []string{
					"create",
				},
				Constraint: Constraint{
					And: []map[string]any{
						{
							"equals": map[string]any{
								"doc": map[string]any{
									"doc": "sys.type",
								},
								"item_type": "Entry",
							},
						},
					},
				},
			},
		},
		Permissions: Permissions{
			ContentModel: []string{
				"read",
			},
			Settings:           "all",
			ContentDelivery:    "all",
			Environments:       "all",
			EnvironmentAliases: "all",
		},
	}

	err = cma.Roles.Upsert(context.Background(), spaceID, role)
	assertions.Nil(err)
}

func TestRolesService_Upsert_Update(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "PUT")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/roles/0xvkNW6WdQ8JkWlWZ8BC4x")

		checkHeaders(r, assertions)

		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)

		description := payload["description"]
		assertions.Equal("Edited text", description)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("role_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	role, err := roleFromTestData("role_1.json")
	assertions.Nil(err)

	role.Description = "Edited text"

	err = cma.Roles.Upsert(context.Background(), spaceID, role)
	assertions.Nil(err)
}

func TestRolesServiceDelete(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "DELETE")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/roles/0xvkNW6WdQ8JkWlWZ8BC4x")
		checkHeaders(r, assertions)

		w.WriteHeader(200)
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	// test role
	role, err := roleFromTestData("role_1.json")
	assertions.Nil(err)

	// delete role
	err = cma.Roles.Delete(context.Background(), spaceID, role.Sys.ID)
	assertions.Nil(err)
}
