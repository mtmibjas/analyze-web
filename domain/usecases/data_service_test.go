package usecases

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

var service = Service{}

func TestHTMLVersion(t *testing.T) {
	tests := []struct {
		html     string
		expected string
	}{
		{"<!DOCTYPE html><html></html>", "HTML5"},
		{"<!DOCTYPE something><html></html>", "Unknown"},
	}

	for _, test := range tests {
		res, err := goquery.NewDocumentFromReader(strings.NewReader(test.html))
		assert.NoError(t, err)
		htmlversion := <-service.getHTMLVersion(res)
		assert.Equal(t, test.expected, htmlversion)
	}
}

func TestHeadings(t *testing.T) {
	html := `
		<!DOCTYPE html>
		<html>
			<head><title>Test</title></head>
			<body>
				<h1>Heading 1</h1>
				<h2>Heading 2</h2>
				<h2>Another Heading 2</h2>
				<h3>Heading 3</h3>
			</body>
		</html>
	`
	res, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	assert.NoError(t, err)
	headings := <-service.getHeadings(res)

	expected := map[string]int{
		"h1": 1,
		"h2": 2,
		"h3": 1,
		"h4": 0,
		"h5": 0,
		"h6": 0,
	}

	assert.Equal(t, expected, headings)
}

func TestLinkAnalysis(t *testing.T) {
	html := `
		<!DOCTYPE html>
		<html>
			<head><title>Test</title></head>
			<body>
				<a href="http://external.com">External</a>
				<a href="/internal">Internal</a>
				<a href="invalid-url">Invalid</a>
			</body>
		</html>
	`
	res, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	assert.NoError(t, err)
	// server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(html))
	// }))
	//defer server.Close()

	result := <-service.processLinks(res)

	assert.Equal(t, 1, result.Internal)
	assert.Equal(t, 1, result.External)
	assert.Equal(t, 1, result.Inaccessible)
}

func TestContainsLoginForm(t *testing.T) {
	tests := []struct {
		html     string
		expected bool
	}{
		{"<form><input type='password'></form>", true},
		{"<form><input type='text'></form>", false},
		{"<div>No form</div>", false},
	}

	for _, test := range tests {
		res, err := goquery.NewDocumentFromReader(strings.NewReader(test.html))
		assert.NoError(t, err)
		containsLogin := <-service.containsLoginForm(res)
		assert.Equal(t, test.expected, containsLogin)
	}
}
