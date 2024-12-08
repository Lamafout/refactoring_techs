package service

import (
	"refactoring_tech/domain/entities"
)

type Repository interface {
	GetListOfTechs () (*entities.Tech, error)
}