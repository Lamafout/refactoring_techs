package controllers

import "refactoring_tech/domain/entities"

type ProducersController interface {
	InsertProducer(producer entities.Producer) error
}