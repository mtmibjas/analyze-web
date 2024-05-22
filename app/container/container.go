package container

import (
	"analyze-web/domain/repositories"
	"net/http"
)

type Container struct {
	Adapters     Adapters
	Repositories Repositories
}

type Repositories struct {
	DataRepository repositories.DataRepositoriesInterface
}

type Adapters struct {
	HTTPClient *http.Client
}
