package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
	tests := []struct {
		input    string
		expected error
	}{
		{"http://testweb.com", nil},
		{"https://testweb.com", nil},
		{"ftp://testweb.com", errors.New("URL scheme must be http or https")},
		{"testweb.com", errors.New("invalid URL format")},
		{"", errors.New("invalid URL format")},
	}

	for _, test := range tests {
		err := ValidateURL(test.input)
		assert.Equal(t, test.expected, err)
	}
}
