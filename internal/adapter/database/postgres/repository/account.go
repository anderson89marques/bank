package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anderson89marques/bank/internal/core/domain"
)

const TableName = "account"

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (r *AccountRepository) Create(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    INSERT INTO %s (document)
    VALUES ($1)
    RETURNING *;
  `, TableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	args := []any{account.Document}
	err = stmt.QueryRowContext(ctx, args...).Scan(&account.ID, &account.Document, &account.CreatedAt, &account.DeletedAt)
	if err != nil {
		fmt.Printf("error inserting file %v \n", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit: %w", err)
	}
	return account, nil
}

func (r *AccountRepository) FindByID(ctx context.Context, id int) (*domain.Account, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    SELECT * FROM %s
    WHERE id = $1;
    `, TableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var account domain.Account
	args := []any{id}
	err = stmt.QueryRowContext(ctx, args...).Scan(&account.ID, &account.Document, &account.CreatedAt, &account.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account info not found: %w", err)
		} else {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
	}
	return &account, nil
}

func (r *AccountRepository) List(ctx context.Context) ([]*domain.Account, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    SELECT * FROM %s;
  `, TableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var accounts []*domain.Account
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query accounts: %w", err)
	}

	for rows.Next() {
		var account domain.Account
		if err := rows.Scan(&account.ID, &account.Document, &account.CreatedAt, &account.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		accounts = append(accounts, &account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during interation: %w", err)
	}
	return accounts, nil
}
