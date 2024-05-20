package usecases

import (
	"analyze-web/app/container"
	"analyze-web/domain/repositories"
)

type Service struct {
	DataRepository repositories.DataRepositoriesInterface
}

func NewDataService(container *container.Container) *Service {
	return &Service{
		DataRepository: container.Repositories.DataRepository,
	}
}
