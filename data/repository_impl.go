package db

import (
	"database/sql"
)

type RepositoryImpl struct {
	db *sql.DB
}