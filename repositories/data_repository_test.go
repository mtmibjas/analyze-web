package repositories

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dr *DataRepository = &DataRepository{
	HTTPClient: &http.Client{},
}

func TestGetURLData_Success(t *testing.T) {
	// Setup a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<html><head><title>Test</title></head><body><p>Hello, World!</p></body></html>`)
	}))
	defer ts.Close()

	doc, err := dr.GetURLData(ts.URL)

	assert.NoError(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "Test", doc.Find("title").Text())
	assert.Equal(t, "Hello, World!", doc.Find("p").Text())
}

func TestGetURLData_HTTPError(t *testing.T) {
	// Setup a test server that returns 500 status
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	doc, err := dr.GetURLData(ts.URL)

	assert.Error(t, err)
	assert.Nil(t, doc)
	assert.Contains(t, err.Error(), "failed to fetch URL")
}

func TestGetURLData_InvalidURL(t *testing.T) {
	doc, err := dr.GetURLData("http://invalid-url")

	assert.Error(t, err)
	assert.Nil(t, doc)
}

func TestGetURLData_InvalidHTML(t *testing.T) {
	// Setup a test server with invalid HTML
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `Invalid HTML`)
	}))
	defer ts.Close()

	doc, err := dr.GetURLData(ts.URL)

	assert.Error(t, err)
	assert.Nil(t, doc)
}

func TestGetURLData_NotFound(t *testing.T) {
	// Setup a test server that returns 404 status
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer ts.Close()

	doc, err := dr.GetURLData(ts.URL)

	assert.Error(t, err)
	assert.Nil(t, doc)
	assert.Contains(t, err.Error(), "failed to fetch URL")
}
