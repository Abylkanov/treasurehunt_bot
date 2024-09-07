package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Создание таблицы, если она не существует
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id BIGINT PRIMARY KEY,
        username TEXT,
        first_name TEXT,
        last_name TEXT
    )`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
