package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anderson89marques/bank/internal/core/domain"
)

const OperationTypeTableName = "operation_type"

type OperationTypeRepository struct {
	DB *sql.DB
}

func NewOperationTypeRepository(db *sql.DB) *OperationTypeRepository {
	return &OperationTypeRepository{
		DB: db,
	}
}

func (r *OperationTypeRepository) FindByID(ctx context.Context, id int) (*domain.OperationType, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    SELECT * FROM %s
    WHERE id = $1;
    `, OperationTypeTableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var operationType domain.OperationType
	args := []any{id}
	err = stmt.QueryRowContext(ctx, args...).Scan(&operationType.ID, &operationType.Description, &operationType.CreatedAt, &operationType.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("operation_type info not found: %w", err)
		} else {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
	}
	return &operationType, nil
}

func (r *OperationTypeRepository) List(ctx context.Context) ([]*domain.OperationType, error) {
	tx, err := r.DB.BeginTx(ctx, nil) // Start a transaction with the received context
	defer tx.Rollback()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := fmt.Sprintf(`
    SELECT * FROM %s;
  `, OperationTypeTableName)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var operations []*domain.OperationType
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query operations: %w", err)
	}

	for rows.Next() {
		var operationType domain.OperationType
		if err := rows.Scan(&operationType.ID, &operationType.Description, &operationType.CreatedAt, &operationType.DeletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		operations = append(operations, &operationType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during interation: %w", err)
	}
	return operations, nil
}
