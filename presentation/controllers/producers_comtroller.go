package controllers

import "refactoring_tech/domain/entities"

type ProducerController interface {
	InsertProducer(producer entities.Producer) error
}
