package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := "user=postgre dbname=postgre password=postgre host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("не удалось подключиться к БД: %v", err)
	}

	DB = db
	return nil
}
