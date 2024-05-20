package resolver

import (
	"analyze-web/app/config"
	"net/http"
	"time"
)

func resolveHttpClientAdapter(cfg config.ServiceConfig) *http.Client {
	return &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}
}
