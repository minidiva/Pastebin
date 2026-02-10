package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupDB() (*sql.DB, error) {

	// тут нужен .env
	dsn := "postgres://postgres:1@localhost:5432/pastes_db?sslmode=disable"

	db, err := sql.Open("pgx", dsn)

	// Open DB
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// Ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB %w", err)
	}

	return db, nil
}

// sql.DB - pool conn

// docker run -d --name pastes-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1 -e POSTGRES_DB=pastes_db -p 5432:5432 postgres:16
