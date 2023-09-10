package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type FinancialAccount struct {
	cli *sql.DB
}

func (f *FinancialAccount) GetByID(ctx context.Context, accountID int) (*entity.FinancialAccount, error) {
	query := `
		SELECT 
			account_id, user_id, currency_id, bank_id, branch_id, account_number, shaba_number,
			account_name, account_type, currency_code, status, created_at, updated_at, deleted_at
		FROM 
			financial_account
		WHERE 
			account_id = $1
	`

	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetByID.PrepareContext: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, accountID)

	var account entity.FinancialAccount
	if err := row.Scan(
		&account.AccountID,
		&account.UserID,
		&account.CurrencyID,
		&account.BankID,
		&account.BranchID,
		&account.AccountNumber,
		&account.ShabaNumber,
		&account.AccountName,
		&account.AccountType,
		&account.CurrencyCode,
		&account.Status,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("FinancialAccount not found")
		}
		return nil, fmt.Errorf("repository.FinancialAccount.GetByID.Scan: %w", err)
	}

	return &account, nil
}

func (f *FinancialAccount) GetByUserID(ctx context.Context, userID int) ([]entity.FinancialAccount, error) {
	query := `
		SELECT 
			account_id, user_id, currency_id, bank_id, branch_id, account_number, shaba_number,
			account_name, account_type, currency_code, status, created_at, updated_at, deleted_at
		FROM 
			financial_account
		WHERE 
			user_id = $1
	`

	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetByUserID.PrepareContext: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetByUserID: %w", err)
	}
	defer rows.Close()

	var accounts []entity.FinancialAccount
	for rows.Next() {
		var account entity.FinancialAccount
		if err := rows.Scan(
			&account.AccountID,
			&account.UserID,
			&account.CurrencyID,
			&account.BankID,
			&account.BranchID,
			&account.AccountNumber,
			&account.ShabaNumber,
			&account.AccountName,
			&account.AccountType,
			&account.CurrencyCode,
			&account.Status,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("repository.FinancialAccount.GetByUserID.Scan: %w", err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (f *FinancialAccount) IsAccountExist(ctx context.Context, accountID int) (bool, error) {
	query := `
        SELECT COUNT(*) 
        FROM financial_account 
        WHERE account_id = $1
    `
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("repository.FinancialAccount.IsAccountExist.Prepare: %w", err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRowContext(ctx, accountID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("repository.FinancialAccount.IsAccountExist: %w", err)
	}
	return count > 0, nil
}

func (f *FinancialAccount) Insert(ctx context.Context, account *entity.FinancialAccount) error {
	query := `
		INSERT INTO financial_account (
			user_id, currency_id, bank_id, branch_id, account_number, shaba_number,
			account_name, account_type, currency_code, status, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.Insert.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		account.UserID, account.CurrencyID, account.BankID, account.BranchID,
		account.AccountNumber, account.ShabaNumber, account.AccountName,
		account.AccountType, account.CurrencyCode, account.Status,
		time.Now(), time.Now(),
	)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.Insert.ExecContext: %w", err)
	}

	return nil
}

func (f *FinancialAccount) Update(ctx context.Context, account *entity.FinancialAccount) error {
	query := `
		UPDATE financial_account
		SET 
			user_id = $1, currency_id = $2, bank_id = $3, branch_id = $4, account_number = $5,
			shaba_number = $6, account_name = $7, account_type = $8, currency_code = $9,
			status = $10, updated_at = $11
		WHERE 
			account_id = $12
	`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.Update.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		account.UserID, account.CurrencyID, account.BankID, account.BranchID,
		account.AccountNumber, account.ShabaNumber, account.AccountName,
		account.AccountType, account.CurrencyCode, account.Status,
		time.Now(), account.AccountID,
	)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.Update.ExecContext: %w", err)
	}

	return nil
}

func (f *FinancialAccount) Delete(ctx context.Context, accountID int) error {
	query := `
        DELETE FROM financial_account 
        WHERE account_id = $1
    `
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.Delete.Prepare: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, accountID)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.Delete: %w", err)
	}
	return nil
}

func (f *FinancialAccount) ListByStatus(ctx context.Context, status enum.FinancialAccountStatus) ([]*entity.FinancialAccount, error) {
	query := `
	SELECT * 
	FROM financial_account 
	WHERE status = $1
`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.ListByStatus.Prepare: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, string(status))
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.ListByStatus: %w", err)
	}
	defer rows.Close()

	var accounts []*entity.FinancialAccount
	for rows.Next() {
		var account entity.FinancialAccount
		err = rows.Scan(&account.AccountID, &account.UserID, &account.CurrencyID, &account.BankID,
			&account.BranchID, &account.AccountNumber, &account.ShabaNumber, &account.AccountName,
			&account.AccountType, &account.CurrencyCode, &account.Status, &account.CreatedAt,
			&account.UpdatedAt, &account.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("repository.FinancialAccount.ListByStatus.Scan: %w", err)
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}

func (f *FinancialAccount) GetByShabaNumber(ctx context.Context, shabaNumber string) (*entity.FinancialAccount, error) {
	query := `
        SELECT * 
        FROM financial_account 
        WHERE shaba_number = $1
    `
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetByShabaNumber.Prepare: %w", err)
	}
	defer stmt.Close()

	var account entity.FinancialAccount
	err = stmt.QueryRowContext(ctx, shabaNumber).Scan(&account.AccountID, &account.UserID, &account.CurrencyID,
		&account.BankID, &account.BranchID, &account.AccountNumber,
		&account.ShabaNumber, &account.AccountName, &account.AccountType,
		&account.CurrencyCode, &account.Status, &account.CreatedAt,
		&account.UpdatedAt, &account.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetByShabaNumber: %w", err)
	}
	return &account, nil
}

func (f *FinancialAccount) GetTransactions(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error) {
	query := `
	SELECT * 
	FROM financial_account_transaction 
	WHERE account_id = $1
`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetTransactions.Prepare: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetTransactions: %w", err)
	}
	defer rows.Close()

	var transactions []*entity.AccountTransaction
	for rows.Next() {
		var transaction entity.AccountTransaction

		err := rows.Scan(
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
		)
		if err != nil {
			return nil, fmt.Errorf("repository.FinancialAccount.GetTransactions.Scan: %w", err)
		}

		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

func (f *FinancialAccount) ListByType(ctx context.Context, accountType string) ([]*entity.FinancialAccount, error) {
	query := `
        SELECT * 
        FROM financial_account 
        WHERE account_type = $1
    `
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.ListByType.Prepare: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, accountType)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.ListByType: %w", err)
	}
	defer rows.Close()

	var accounts []*entity.FinancialAccount
	for rows.Next() {
		var account entity.FinancialAccount
		err := rows.Scan(
			&account.AccountID,
			&account.UserID,
			&account.CurrencyID,
			&account.BankID,
			&account.BranchID,
			&account.AccountNumber,
			&account.ShabaNumber,
			&account.AccountName,
			&account.AccountType,
			&account.CurrencyCode,
			&account.Status,
			&account.CreatedAt,
			&account.UpdatedAt,
			&account.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("repository.FinancialAccount.ListByType.Scan: %w", err)
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (f *FinancialAccount) GetCurrencyByAccountID(ctx context.Context, accountID int) (*entity.Currency, error) {
	query := `
	SELECT c.* 
	FROM currency c 
	JOIN financial_account f ON f.currency_id = c.currency_id 
	WHERE f.account_id = $1
`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetCurrencyByAccountID.Prepare: %w", err)
	}
	defer stmt.Close()

	var currency entity.Currency
	err = stmt.QueryRowContext(ctx, accountID).Scan(
		&currency.CurrencyID,
		&currency.CurrencyCode,
		&currency.CurrencyName,
		&currency.Symbol,
		&currency.ExchangeRate,
		&currency.CreatedAt,
		&currency.UpdatedAt,
		&currency.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetCurrencyByAccountID: %w", err)
	}

	return &currency, nil
}

func (f *FinancialAccount) GetBranchByAccountID(ctx context.Context, accountID int) (*entity.BankBranch, error) {
	query := `
	SELECT b.* 
	FROM bank_branches b 
	JOIN financial_account f ON f.branch_id = b.branch_id 
	WHERE f.account_id = $1
`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetBranchByAccountID.Prepare: %w", err)
	}
	defer stmt.Close()

	var branch entity.BankBranch
	err = stmt.QueryRowContext(ctx, accountID).Scan(
		&branch.BranchID,
		&branch.BankID,
		&branch.BranchName,
		&branch.BranchCode,
		&branch.Address,
		&branch.City,
		&branch.Province,
		&branch.PostalCode,
		&branch.PhoneNumber,
		&branch.Status,
		&branch.CreatedAt,
		&branch.UpdatedAt,
		&branch.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetBranchByAccountID: %w", err)
	}

	return &branch, nil
}

func (f *FinancialAccount) GetBankByAccountID(ctx context.Context, accountID int) (*entity.Bank, error) {
	query := `
	SELECT b.* 
	FROM banks b 
	JOIN financial_account f ON f.bank_id = b.bank_id 
	WHERE f.account_id = $1
`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetBankByAccountID.Prepare: %w", err)
	}
	defer stmt.Close()

	var bank entity.Bank
	err = stmt.QueryRowContext(ctx, accountID).Scan(
		&bank.BankID,
		&bank.Name,
		&bank.BankCode,
		&bank.Status,
		&bank.CreatedAt,
		&bank.UpdatedAt,
		&bank.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialAccount.GetBankByAccountID: %w", err)
	}

	return &bank, nil
}

func (f *FinancialAccount) UpdateStatus(ctx context.Context, accountID int, status string) error {
	query := `
		UPDATE financial_account
		SET status = $2, updated_at = CURRENT_TIMESTAMP
		WHERE account_id = $1
	`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.UpdateStatus.Prepare: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, accountID, status)
	if err != nil {
		return fmt.Errorf("repository.FinancialAccount.UpdateStatus: %w", err)
	}

	return nil
}

func (f *FinancialAccount) GetAccountStatus(ctx context.Context, accountID int) (enum.FinancialAccountStatus, error) {
	query := `
		SELECT status
		FROM financial_account
		WHERE account_id = $1
	`
	stmt, err := f.cli.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("repository.FinancialAccount.GetAccountStatus.Prepare: %w", err)
	}
	defer stmt.Close()

	var status enum.FinancialAccountStatus
	err = stmt.QueryRowContext(ctx, accountID).Scan(&status)
	if err != nil {
		return 0, fmt.Errorf("repository.FinancialAccount.GetAccountStatus: %w", err)
	}

	return status, nil
}
