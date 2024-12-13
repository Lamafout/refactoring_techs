package service

import (
	"refactoring_tech/domain/entities"
)

type Repository interface {
	GetListOfTechs () (*[]entities.TechShort, error)
	GetConcreteTech (id int) (*entities.Tech, error)
	InsertProducer (producer entities.Producer) error
}