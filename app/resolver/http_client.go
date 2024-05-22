package resolver

import (
	"analyze-web/app/config"
	"net/http"
	"time"
)

func resolveHTTPClientAdapter(cfg *config.Config) *http.Client {
	return &http.Client{
		Timeout: time.Duration(cfg.Service.Timeout) * time.Second,
	}
}
