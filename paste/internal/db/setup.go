package db

import (
	"database/sql"
	"fmt"
	"pastebin/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupDB(cfg *config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDBName,
	)
	db, err := sql.Open("pgx", dsn)

	// Open DB
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// Ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB %w", err)
	}

	fmt.Println("DB Connected successfully!")
	return db, nil

}

// psql -U postgres -d pastes_db

// CREATE TABLE pastes (id SERIAL PRIMARY KEY, key TEXT NOT NULL UNIQUE, expires_at TIMESTAMP WITH TIME ZONE NOT NULL, created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW());

// docker run -d --name pastes-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1 -e POSTGRES_DB=pastes_db -p 5432:5432 postgres:16
