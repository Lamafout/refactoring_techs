package controllers

import "refactoring_tech/domain/entities"

type TechsController interface {
	GetTechs() (*[] entities.Tech, error)
}