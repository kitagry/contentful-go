package contentful

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotsService_ListEntrySnapshots(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environments/master/entries/hfM9RCJIk0wIm06WkEOQY/snapshots")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("snapshot_entry.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.Snapshots.ListEntrySnapshots(context.Background(), spaceID, "hfM9RCJIk0wIm06WkEOQY", nil)
	assertions.Nil(err)
	entrySnapshot := collection.Items
	assertions.Equal(1, len(entrySnapshot))
	assertions.Equal("Hello, World!", entrySnapshot[0].EntrySnapshotDetail.Fields["title"].(map[string]interface{})["en-US"])
}

func TestSnapshotsService_GetEntrySnapshot(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environments/master/entries/hfM9RCJIk0wIm06WkEOQY/snapshots/4FLrUHftHW3v2BLi9fzfjU")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("snapshot_entry_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	entrySnapshot, err := cma.Snapshots.GetEntrySnapshot(context.Background(), spaceID, "hfM9RCJIk0wIm06WkEOQY", "4FLrUHftHW3v2BLi9fzfjU")
	assertions.Nil(err)
	assertions.Equal("Hello, World!", entrySnapshot.EntrySnapshotDetail.Fields["title"].(map[string]interface{})["en-US"])
}

func TestSnapshotsService_GetEntrySnapshot_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environments/master/entries/hfM9RCJIk0wIm06WkEOQY/snapshots/4FLrUHftHW3v2BLi9fzfjU")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("snapshot_entry_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.Snapshots.GetEntrySnapshot(context.Background(), spaceID, "hfM9RCJIk0wIm06WkEOQY", "4FLrUHftHW3v2BLi9fzfjU")
	assertions.Nil(err)
}

func TestSnapshotsService_ListContentTypeSnapshots(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environments/master/content_types/hfM9RCJIk0wIm06WkEOQY/snapshots")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("snapshot_content_type.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	collection, err := cma.Snapshots.ListContentTypeSnapshots(context.Background(), spaceID, "hfM9RCJIk0wIm06WkEOQY", nil)
	assertions.Nil(err)
	entrySnapshot := collection.Items
	assertions.Equal(1, len(entrySnapshot))
	assertions.Equal("Blog Post", entrySnapshot[0].ContentTypeSnapshotDetail.Name)
}

func TestSnapshotsService_GetContentTypeSnapshots(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environments/master/content_types/hfM9RCJIk0wIm06WkEOQY/snapshots/4FLrUHftHW3v2BLi9fzfjU")

		checkHeaders(r, assertions)

		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, readTestData("snapshot_content_type_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	entrySnapshot, err := cma.Snapshots.GetContentTypeSnapshots(context.Background(), spaceID, "hfM9RCJIk0wIm06WkEOQY", "4FLrUHftHW3v2BLi9fzfjU")
	assertions.Nil(err)
	assertions.Equal("Blog Post", entrySnapshot.ContentTypeSnapshotDetail.Name)
}

func TestSnapshotsService_GetContentTypeSnapshots_2(t *testing.T) {
	var err error
	assertions := assert.New(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertions.Equal(r.Method, "GET")
		assertions.Equal(r.URL.Path, "/spaces/"+spaceID+"/environments/master/content_types/hfM9RCJIk0wIm06WkEOQY/snapshots/4FLrUHftHW3v2BLi9fzfjU")

		checkHeaders(r, assertions)

		w.WriteHeader(400)
		_, _ = fmt.Fprintln(w, readTestData("snapshot_content_type_1.json"))
	})

	// test server
	server := httptest.NewServer(handler)
	defer server.Close()

	// cma client
	cma = NewCMA(CMAToken)
	cma.BaseURL = server.URL

	_, err = cma.Snapshots.GetContentTypeSnapshots(context.Background(), spaceID, "hfM9RCJIk0wIm06WkEOQY", "4FLrUHftHW3v2BLi9fzfjU")
	assertions.Nil(err)
}
