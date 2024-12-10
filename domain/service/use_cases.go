package service

import (
	"refactoring_tech/domain/entities"
)

type UseCases struct {
	Repository Repository
}

func NewUseCases(repository Repository) *UseCases {
	return &UseCases{
		Repository: repository,
	}
}

func (u *UseCases) GetListOfTechs() (*[]entities.Tech, error) {
	return u.Repository.GetListOfTechs()
}