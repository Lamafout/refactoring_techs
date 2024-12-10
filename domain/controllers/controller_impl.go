package controllers

import (
	service "refactoring_tech/domain"
	"refactoring_tech/domain/entities"
)

type ControllerImpl struct {
	UseCases *service.UseCases
}

func NewControllerImpl(useCases *service.UseCases) *ControllerImpl {
	return &ControllerImpl{
		UseCases: useCases,
	}
}

func (c *ControllerImpl) GetTechs() (*[]entities.Tech, error) {
	return c.UseCases.GetListOfTechs()
}