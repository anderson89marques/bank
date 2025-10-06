package ports

import (
	"context"

	"github.com/anderson89marques/bank/internal/core/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, account *domain.Transaction) (*domain.Transaction, error)
	List(ctx context.Context) ([]*domain.Transaction, error)
}

type TransactionService interface {
	Create(ctx context.Context, account *domain.Transaction) (*domain.Transaction, error)
	List(ctx context.Context) ([]*domain.Transaction, error)
}
