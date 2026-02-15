package entity

import "time"

type Paste struct {
	Key       string
	ExpiresAt time.Time
	// OwnerID   int // foreign key к userID
}
