package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
)

type AccountTransaction struct {
	cli *sql.DB
}

func (repo *AccountTransaction) Insert(ctx context.Context, transaction *entity.AccountTransaction) error {
	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	query := `
		INSERT INTO public.account_transaction (
			transaction_group_id, financial_account_id, amount, balance,
			description, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING transaction_id;
	`

	_, err = tx.ExecContext(ctx, query,
		transaction.TransactionGroupID,
		transaction.FinancialAccountID,
		transaction.Amount,
		transaction.Balance,
		transaction.Description,
		transaction.Status,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit: %w", err)
	}

	return nil
}

func (repo *AccountTransaction) GetByID(ctx context.Context, transactionID int64) (*entity.AccountTransaction, error) {
	query := `
		SELECT 
			transaction_id, transaction_group_id, financial_account_id, amount,
			balance, description, status, created_at, updated_at, deleted_at
		FROM 
			public.account_transaction
		WHERE 
			transaction_id = $1
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("PrepareContext: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, transactionID)
	transaction := &entity.AccountTransaction{}
	if err := row.Scan(
		&transaction.TransactionID,
		&transaction.TransactionGroupID,
		&transaction.FinancialAccountID,
		&transaction.Amount,
		&transaction.Balance,
		&transaction.Description,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Consider returning a custom error if needed
		}
		return nil, fmt.Errorf("Scan: %w", err)
	}
	return transaction, nil
}

func (repo *AccountTransaction) Delete(ctx context.Context, transactionID int64) error {
	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	query := `
		DELETE FROM public.account_transaction WHERE transaction_id = $1;
	`

	_, err = tx.ExecContext(ctx, query, transactionID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit: %w", err)
	}

	return nil
}

func (repo *AccountTransaction) ListByAccountID(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error) {
	query := `
		SELECT 
			transaction_id, transaction_group_id, financial_account_id, amount,
			balance, description, status, created_at, updated_at, deleted_at
		FROM 
			public.account_transaction
		WHERE 
			financial_account_id = $1
	`

	rows, err := repo.cli.QueryContext(ctx, query, accountID)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()

	var transactions []*entity.AccountTransaction
	for rows.Next() {
		transaction := &entity.AccountTransaction{}
		if err := rows.Scan(
			&transaction.TransactionID,
			&transaction.TransactionGroupID,
			&transaction.FinancialAccountID,
			&transaction.Amount,
			&transaction.Balance,
			&transaction.Description,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repo *AccountTransaction) ListByTransactionGroupID(ctx context.Context, groupID int) ([]*entity.AccountTransaction, error) {
	query := `
		SELECT 
			transaction_id, transaction_group_id, financial_account_id, amount,
			balance, description, status, created_at, updated_at, deleted_at
		FROM 
			public.account_transaction
		WHERE 
			transaction_group_id = $1
	`

	rows, err := repo.cli.QueryContext(ctx, query, groupID)
	if err != nil {
		return nil, fmt.Errorf("QueryContext: %w", err)
	}
	defer rows.Close()

	var transactions []*entity.AccountTransaction
	for rows.Next() {
		transaction := &entity.AccountTransaction{}
		if err := rows.Scan(
			&transaction.TransactionID,
			&transaction.TransactionGroupID,
			&transaction.FinancialAccountID,
			&transaction.Amount,
			&transaction.Balance,
			&transaction.Description,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repo *AccountTransaction) BeginTx(ctx context.Context) (context.Context, error) {
	// TODO: Implement
	return nil, nil
}

func (repo *AccountTransaction) CommitTx(ctx context.Context) error {
	// TODO: Implement
	return nil
}

func (repo *AccountTransaction) RollbackTx(ctx context.Context) error {
	// TODO: Implement
	return nil
}

func (repo *AccountTransaction) DebitAccount(ctx context.Context, accountID int, amount float64) error {
	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	// Retrieve current balance
	var currentBalance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM financial_account WHERE id = $1 FOR UPDATE", accountID).Scan(&currentBalance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("QueryRowContext: %w", err)
	}

	// Check if sufficient balance
	if currentBalance < amount {
		tx.Rollback()
		return fmt.Errorf("Insufficient funds")
	}

	// Update balance
	newBalance := currentBalance - amount
	_, err = tx.ExecContext(ctx, "UPDATE financial_account SET balance = $1 WHERE id = $2", newBalance, accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	// Insert transaction
	_, err = tx.ExecContext(ctx, `
		INSERT INTO account_transaction (financial_account_id, amount, balance, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, 'Debit', 'Completed', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`, accountID, -amount, newBalance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	return tx.Commit()
}

func (repo *AccountTransaction) CreditAccount(ctx context.Context, accountID int, amount float64) error {
	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	// Retrieve current balance
	var currentBalance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM financial_account WHERE id = $1 FOR UPDATE", accountID).Scan(&currentBalance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("QueryRowContext: %w", err)
	}

	// Update balance
	newBalance := currentBalance + amount
	_, err = tx.ExecContext(ctx, "UPDATE financial_account SET balance = $1 WHERE id = $2", newBalance, accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	// Insert transaction
	_, err = tx.ExecContext(ctx, `
		INSERT INTO account_transaction (financial_account_id, amount, balance, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, 'Credit', 'Completed', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`, accountID, amount, newBalance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	return tx.Commit()
}

func (repo *AccountTransaction) CreateNewTransactionGroup(ctx context.Context) (int, error) {
	var newID int
	query := "INSERT INTO transaction_groups DEFAULT VALUES RETURNING id;"

	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("BeginTx: %w", err)
	}

	err = tx.QueryRowContext(ctx, query).Scan(&newID)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("QueryRowContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("Commit: %w", err)
	}

	return newID, nil
}

func (repo *AccountTransaction) Update(ctx context.Context, transaction *entity.AccountTransaction) error {
	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	query := `
		UPDATE public.account_transaction
		SET
			transaction_group_id = $2,
			financial_account_id = $3,
			amount = $4,
			balance = $5,
			description = $6,
			status = $7,
			updated_at = CURRENT_TIMESTAMP
		WHERE 
			transaction_id = $1;
	`

	_, err = tx.ExecContext(ctx, query,
		transaction.TransactionID,
		transaction.TransactionGroupID,
		transaction.FinancialAccountID,
		transaction.Amount,
		transaction.Balance,
		transaction.Description,
		transaction.Status,
	)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ExecContext: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit: %w", err)
	}

	return nil
}

func (repo *AccountTransaction) GetCurrentBalance(ctx context.Context, accountID int) (float64, error) {
	return 0.0 ,nil
}




// type AccountTransactionRepository interface {
// 	Insert(ctx context.Context, transaction *entity.AccountTransaction) error
// 	Delete(ctx context.Context, transactionID int64) error
// 	GetByID(ctx context.Context, transactionID int64) (*entity.AccountTransaction, error)
// 	ListByAccountID(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error)
// 	ListByTransactionGroupID(ctx context.Context, groupID int) ([]*entity.AccountTransaction, error)
// 	CreateNewTransactionGroup(ctx context.Context) (int, error)
// 	Update(ctx context.Context, transaction *entity.AccountTransaction) error
// 	GetCurrentBalance(ctx context.Context, accountID int) (float64, error)

// 	BeginTx(ctx context.Context) (context.Context, error) // Begins a new transaction
// 	CommitTx(ctx context.Context) error                   // Commits the transaction
// 	RollbackTx(ctx context.Context) error                 // Rollbacks the transaction

// 	// Transfer related methods
// 	DebitAccount(ctx context.Context, accountID int, amount float64) error
// 	CreditAccount(ctx context.Context, accountID int, amount float64) error
// }
