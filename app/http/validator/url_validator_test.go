package validator

import (
	"analyze-web/app/config"
	"errors"
	"os"
	"testing"

	"analyze-web/pkg/logger/zap"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	cfg := config.Parse("/config")
	loggger := zap.NewLogger(cfg)
	loggger.Init()
	code := m.Run()
	os.Exit(code)
}

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
