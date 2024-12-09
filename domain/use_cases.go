package service

import "refactoring_tech/domain/entities"

type UseCases struct {
	Repository Repository
	Controller *Controller
}

func (u *UseCases) GetListOfTechs() (*[]entities.Tech, error) {
	return u.Repository.GetListOfTechs()
}