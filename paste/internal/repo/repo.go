package repo

import (
	"context"
	"database/sql"
	"fmt"
	"pastebin/internal/entity"
)

type PasteRepo struct {
	db *sql.DB
}

func NewPasteRepo(db *sql.DB) *PasteRepo {
	return &PasteRepo{
		db: db,
	}
}

func (r *PasteRepo) CreatePaste(ctx context.Context, paste entity.Paste) error {
	query := `
		INSERT INTO pastes (key, expires_at)
		VALUES ($1, $2);
	`

	_, err := r.db.ExecContext(ctx, query,
		paste.Key,
		paste.ExpiresAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert paste into DB: %w", err)
	}

	return nil
}

func (r *PasteRepo) GetPaste(ctx context.Context, key string) (*entity.Paste, error) {
	query := `
		SELECT key, expires_at
		FROM pastes 
		WHERE key = $1;
	`

	row := r.db.QueryRowContext(ctx, query, key)

	var paste entity.Paste

	err := row.Scan(
		&paste.Key,
		&paste.ExpiresAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &paste, nil
}
