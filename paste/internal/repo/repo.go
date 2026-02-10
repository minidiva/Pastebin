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
		INSERT INTO pastes (link, expires_at, owner_id)
		VALUES ($1, $2, $3);
	`

	_, err := r.db.ExecContext(ctx, query,
		paste.Link,
		paste.ExpiresAt,
		paste.OwnerID,
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
