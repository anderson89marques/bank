package services

import (
	"context"

	"github.com/anderson89marques/bank/internal/core/domain"
	"github.com/anderson89marques/bank/internal/core/ports"
)

type AccountService struct {
	repo ports.AccountRepository
}

func NewAccountService(repo ports.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) Create(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	newAccount, err := s.repo.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

func (s *AccountService) FindByID(ctx context.Context, id int) (*domain.Account, error) {
	account, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountService) List(ctx context.Context) ([]*domain.Account, error) {
	accounts, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
