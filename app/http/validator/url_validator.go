package validator

import (
	"errors"
	"net/url"
)

func ValidateURL(inputURL string) error {
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return errors.New("invalid URL format")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL scheme must be http or https")
	}

	return nil
}
