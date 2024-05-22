package repositories

import (
	"analyze-web/app/config"
	"analyze-web/app/container"
	"net/http"
)

type DataRepository struct {
	HTTPClient *http.Client
}

func NewDataRepository(c *config.Config, a *container.Adapters) *DataRepository {
	return &DataRepository{
		HTTPClient: a.HTTPClient,
	}
}
