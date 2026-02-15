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

//
// type Paste struct {
// 	ID        int
// 	Link      string
// 	CreatedAt time.Time
// 	ExpiresAt time.Time
// 	OwnerID   int // foreign key к userID
// }
