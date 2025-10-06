package services

import (
	"context"
	"fmt"

	"github.com/anderson89marques/bank/internal/core/ports"
)

type OperationTypeService struct {
	repo ports.OperationTypeRepository
}

func NewOperationTypeService(repo ports.OperationTypeRepository) *OperationTypeService {
	return &OperationTypeService{
		repo: repo,
	}
}

func (o *OperationTypeService) Apply(ctx context.Context, operationTypeID int, amount float64) (float64, error) {
	operationType, err := o.repo.FindByID(context.Background(), operationTypeID)
	if err != nil {
		return 0, err
	}

	strategy, ok := strategyMap[operationType.Description]
	if !ok {
		return 0, fmt.Errorf("strategy not found for OperationType: %s", operationType.Description)
	}
	adjustedAmount := strategy.Apply(amount)
	return adjustedAmount, nil
}

// OperationType strategies
type DebitStrategy struct{}

func (d *DebitStrategy) Apply(amount float64) float64 {
	return -amount
}

type CreditStrategy struct{}

func (c *CreditStrategy) Apply(amount float64) float64 {
	return amount
}

var strategyMap = map[string]ports.AmountStrategy{
	"Normal Purchase":            &DebitStrategy{},
	"Purchase with installments": &DebitStrategy{},
	"Withdrawal":                 &DebitStrategy{},
	"Credit Voucher":             &CreditStrategy{},
}
