package controllers

import "refactoring_tech/domain/entities"

type TechsController interface {
	GetTechs() (*[] entities.TechShort, error)
	GetTechById(id int) (*entities.Tech, error)
}