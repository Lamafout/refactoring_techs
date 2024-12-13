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

func (u *UseCases) GetListOfTechs() (*[]entities.TechShort, error) {
	return u.Repository.GetListOfTechs()
}

func (u *UseCases) GetTechById(id int) (*entities.Tech, error) {
	return u.Repository.GetConcreteTech(id)
}

func (u *UseCases) InsertProducer(producer entities.Producer) error {
	return u.Repository.InsertProducer(producer)
}