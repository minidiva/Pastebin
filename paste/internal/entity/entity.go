package entity

import "time"

type Paste struct {
	ID        int
	Link      string
	CreatedAt time.Time
	ExpiresAt time.Time
	OwnerID   int // foreign key к userID
}
