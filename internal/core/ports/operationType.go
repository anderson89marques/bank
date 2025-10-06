package ports

import (
	"context"

	"github.com/anderson89marques/bank/internal/core/domain"
)

type OperationTypeRepository interface {
	FindByID(ctx context.Context, id int) (*domain.OperationType, error)
	List(ctx context.Context) ([]*domain.OperationType, error)
}

type OperationTypeService interface {
	Apply(ctx context.Context, operationID int, amount float64) (float64, error)
}

// Operation Type Strategies
type AmountStrategy interface {
	Apply(amount float64) float64
}
