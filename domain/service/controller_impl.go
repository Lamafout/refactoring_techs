package service

import (
	"refactoring_tech/domain/entities"
)

type ControllerImpl struct {
	UseCases *UseCases
}

func NewControllerImpl(useCases *UseCases) *ControllerImpl {
	return &ControllerImpl{
		UseCases: useCases,
	}
}

func (c *ControllerImpl) GetTechs() (*[]entities.TechShort, error) {
	return c.UseCases.GetListOfTechs()
}

func (c *ControllerImpl) GetTechById(id int) (*entities.Tech, error) {
	return c.UseCases.GetTechById(id)
}