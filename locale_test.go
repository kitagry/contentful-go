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

func TestLocalesService_List(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/locales")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("locale.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.Locales.List(context.Background(), spaceID, nil)
	assertions.Nil(err)
	locale := collection.Items
	assertions.Equal("34N35DoyUQAtaKwWTgZs34", locale[0].Sys.ID)
}

func TestLocalesService_Get(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/locales/4aGeQYgByqQFJtToAOh2JJ")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("locale_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	locale, err := cma.Locales.Get(context.Background(), spaceID, "4aGeQYgByqQFJtToAOh2JJ")
	assertions.Nil(err)
	assertions.Equal("U.S. English", locale.Name)
	assertions.Equal("en-US", locale.Code)
}

func TestLocalesService_Get_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/locales/4aGeQYgByqQFJtToAOh2JJ")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("locale_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.Locales.Get(context.Background(), spaceID, "4aGeQYgByqQFJtToAOh2JJ")
	assertions.NotNil(err)
}

func TestLocalesService_Upsert_Create(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "POST")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/locales")

		checkHeaders(r, assertions)

		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)
		assertions.Equal("German (Austria)", payload["name"])
		assertions.Equal("de-AT", payload["code"])

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("locale_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	locale := &Locale{
		Name: "German (Austria)",
		Code: "de-AT",
	}

	err = cma.Locales.Upsert(context.Background(), spaceID, locale)
	assertions.Nil(err)
}

func TestLocalesService_Upsert_Update(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "PUT")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/locales/4aGeQYgByqQFJtToAOh2JJ")

		checkHeaders(r, assertions)

		var payload map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		assertions.Nil(err)
		assertions.Equal("modified-name", payload["name"])
		assertions.Equal("modified-code", payload["code"])

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("locale_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	locale, err := localeFromTestData("locale_1.json")
	assertions.Nil(err)

	locale.Name = "modified-name"
	locale.Code = "modified-code"

	err = cma.Locales.Upsert(context.Background(), spaceID, locale)
	assertions.Nil(err)
}

func TestLocalesService_Delete(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "DELETE")
		assertions.Equal(r.RequestURI, "/spaces/"+spaceID+"/locales/4aGeQYgByqQFJtToAOh2JJ")
		checkHeaders(r, assertions)

		w.WriteHeader(200)
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	// test locale
	locale, err := localeFromTestData("locale_1.json")
	assertions.Nil(err)

	// delete locale
	err = cma.Locales.Delete(context.Background(), spaceID, locale)
	assertions.Nil(err)
}
