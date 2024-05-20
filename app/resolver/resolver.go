package resolver

import (
	"analyze-web/app/config"
	"analyze-web/app/container"
)

type Resolver struct {
	Config       *config.Config
	Adapters     container.Adapters
	Repositories container.Repositories
}

func NewAdapter(cfg *config.Config) *Resolver {
	return &Resolver{
		Config: cfg,
	}
}

func (a *Resolver) Resolve() *container.Container {

	a.resolveDBAdapters()
	a.resolveRepositories()

	return &container.Container{
		Adapters:     a.Adapters,
		Repositories: a.Repositories,
	}
}

func (a *Resolver) resolveDBAdapters() {

	httpCient := resolveHttpClientAdapter(a.Config.Service)
	a.Adapters = container.Adapters{
		HttpClient: httpCient,
	}
}
