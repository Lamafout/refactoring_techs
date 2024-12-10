package controllers

import (
	service "refactoring_tech/domain"
	"refactoring_tech/domain/entities"
)

type ControllerImpl struct {
	UseCases *service.UseCases
}

func (c *ControllerImpl) GetTech() (*[]entities.Tech, error) {
	return c.UseCases.GetListOfTechs()
}