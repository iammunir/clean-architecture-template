package usecase

import "github.com/iammunir/clean-architecture-template/repository"

type UseCase interface {
	Ping() error
}

type useCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (usecase *useCase) Ping() error {
	return usecase.repo.Ping()
}
