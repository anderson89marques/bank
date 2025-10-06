package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anderson89marques/bank/internal/core/domain"
)

const TransactionTableName = "transaction"

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    INSERT INTO %s (account_id, operation_type_id, amount)
    VALUES ($1, $2, $3)
    RETURNING *;
  `, TransactionTableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	args := []any{transaction.AccountID, transaction.OperationTypeID, transaction.Amount}
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&transaction.ID, &transaction.AccountID, &transaction.OperationTypeID,
		&transaction.Amount, &transaction.CreatedAt, &transaction.DeletedAt)
	if err != nil {
		fmt.Printf("error inserting file %v \n", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit: %w", err)
	}
	return transaction, nil
}

func (r *TransactionRepository) List(ctx context.Context) ([]*domain.Transaction, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    SELECT * FROM %s;
  `, TransactionTableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var transactions []*domain.Transaction
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query accounts: %w", err)
	}

	for rows.Next() {
		var transaction domain.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.OperationTypeID,
			&transaction.Amount, &transaction.CreatedAt, &transaction.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		transactions = append(transactions, &transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during interation: %w", err)
	}
	return transactions, nil
}
