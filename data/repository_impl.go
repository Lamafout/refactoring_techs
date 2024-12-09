package db

import (
	"database/sql"
	"refactoring_tech/data/models"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) GetTechs() (*[]models.TechModel, error) {
	
}