package ports

import (
	"context"

	"github.com/anderson89marques/bank/internal/core/domain"
)

type AccountRepository interface {
	Create(ctx context.Context, account *domain.Account) (*domain.Account, error)
	FindByID(ctx context.Context, id int) (*domain.Account, error)
	List(ctx context.Context) ([]*domain.Account, error)
}

type AccountService interface {
	Create(ctx context.Context, account *domain.Account) (*domain.Account, error)
	FindByID(ctx context.Context, id int) (*domain.Account, error)
	List(ctx context.Context) ([]*domain.Account, error)
}
