package db

import(
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	DbName string
}

func ConnectToPostgres(config Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %w", err)
	}
	return db, nil
}