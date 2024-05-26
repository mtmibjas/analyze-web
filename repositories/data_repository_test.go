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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<html><head><title>Test</title></head><body><p>Hello, World!</p></body></html>`)
	}))
	defer ts.Close()

	res, err := dr.GetURLData(ts.URL)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Test", res.Find("title").Text())
	assert.Equal(t, "Hello, World!", res.Find("p").Text())
}

func TestGetURLData_HTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	res, err := dr.GetURLData(ts.URL)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "failed to fetch URL")
}

func TestGetURLData_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer ts.Close()

	res, err := dr.GetURLData(ts.URL)
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "failed to fetch URL")
}
