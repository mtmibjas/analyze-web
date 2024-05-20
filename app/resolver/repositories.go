package resolver

import (
	"analyze-web/app/container"
	"analyze-web/repositories"
)

func (r *Resolver) resolveRepositories() {
	r.Repositories = container.Repositories{
		DataRepository: repositories.NewDataRepository(r.Config, &r.Adapters),
	}
}
