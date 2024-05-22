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

func (r *Resolver) Resolve() *container.Container {
	r.resolveDBAdapters()
	r.resolveRepositories()

	return &container.Container{
		Adapters:     r.Adapters,
		Repositories: r.Repositories,
	}
}

func (r *Resolver) resolveDBAdapters() {
	httpCient := resolveHTTPClientAdapter(r.Config)
	r.Adapters = container.Adapters{
		HTTPClient: httpCient,
	}
}
