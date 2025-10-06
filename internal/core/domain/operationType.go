package domain

import "time"

type OperationType struct {
	ID          int
	Description string
	CreatedAt   time.Time
	DeletedAt   *time.Time
}
