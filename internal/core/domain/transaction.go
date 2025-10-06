package domain

import "time"

type Transaction struct {
	ID              int
	AccountID       int
	OperationTypeID int
	Amount          float64
	CreatedAt       time.Time
	DeletedAt       *time.Time
}
