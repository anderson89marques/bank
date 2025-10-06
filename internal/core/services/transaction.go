package services

import (
	"context"

	"github.com/anderson89marques/bank/internal/core/domain"
	"github.com/anderson89marques/bank/internal/core/ports"
)

type TransactionService struct {
	transactionRepo  ports.TransactionRepository
	operationTypeSrv ports.OperationTypeService
}

func NewTransactionService(transactionRepo ports.TransactionRepository, operationTypeSrv ports.OperationTypeService) *TransactionService {
	return &TransactionService{
		transactionRepo:  transactionRepo,
		operationTypeSrv: operationTypeSrv,
	}
}

func (s *TransactionService) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	amount, err := s.operationTypeSrv.Apply(ctx, transaction.OperationTypeID, transaction.Amount)
	if err != nil {
		return nil, err
	}
	transaction.Amount = amount
	newTransaction, err := s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (s *TransactionService) List(ctx context.Context) ([]*domain.Transaction, error) {
	accounts, err := s.transactionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
