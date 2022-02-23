package usecase

import (
	"fmt"

	"github.com/guil95/go-wire-di/internal/repository"
)

type DomainUseCase struct {
	repo repository.Repository
}

type UseCase interface {
	GetName(name string) string
}

func NewUseCase(repository repository.Repository) UseCase {
	return &DomainUseCase{
		repo: repository,
	}
}

func (uc *DomainUseCase) GetName(name string) string{
	return fmt.Sprintf("My name is %s", name)
}
