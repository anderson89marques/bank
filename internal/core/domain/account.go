package domain

import "time"

type Account struct {
	ID        int
	Document  string
	CreatedAt time.Time
	DeletedAt *time.Time
}
