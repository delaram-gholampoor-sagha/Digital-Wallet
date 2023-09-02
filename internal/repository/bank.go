package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
)

type Bank struct {
	cli *sql.DB
}

func NewBankRepository(cli *sql.DB) *Bank {
	return &Bank{cli}
}

func (repo *Bank) GetByID(ctx context.Context, bankID int) (*entity.Bank, error) {
	query := `
		SELECT
			bank_id, name, bank_code, status,
			created_at, updated_at, deleted_at
		FROM
			bank
		WHERE
			bank_id = $1 AND deleted_at IS NULL
	`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Bank.GetByID.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, bankID)

	var bank entity.Bank
	if err := row.Scan(&bank.BankID, &bank.Name, &bank.BankCode, &bank.Status, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, derror.NewNotFoundError("Bank not found")
		}
		return nil, fmt.Errorf("repository.Bank.GetByID.Scan: %w", err)
	}

	return &bank, nil
}

func (repo *Bank) GetByCode(ctx context.Context, bankCode string) (*entity.Bank, error) {
	query := `
		SELECT
			bank_id, name, bank_code, status,
			created_at, updated_at, deleted_at
		FROM
			bank
		WHERE
			bank_code = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Bank.GetByCode.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, bankCode)

	var bank entity.Bank
	if err := row.Scan(&bank.BankID, &bank.Name, &bank.BankCode, &bank.Status, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, derror.NewNotFoundError("Bank not found")
		}
		return nil, fmt.Errorf("repository.Bank.GetByCode.Scan: %w", err)
	}

	return &bank, nil
}

func (repo *Bank) GetByName(ctx context.Context, bankName string) (*entity.Bank, error) {
	query := `
		SELECT
			bank_id, name, bank_code, status,
			created_at, updated_at, deleted_at
		FROM
			public.bank
		WHERE
			name = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Bank.GetByName.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, bankName)

	var bank entity.Bank
	if err := row.Scan(&bank.BankID, &bank.Name, &bank.BankCode, &bank.Status, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, derror.NewNotFoundError("Bank not found")
		}
		return nil, fmt.Errorf("repository.Bank.GetByName.Scan: %w", err)
	}

	return &bank, nil
}

func (repo *Bank) Insert(ctx context.Context, bank *entity.Bank) error {
	query := `
		INSERT INTO bank (name, bank_code, status, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING bank_id
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.Bank.Insert.PrepareContext: %w", err)
	}

	err = stmt.QueryRowContext(ctx, bank.Name, bank.BankCode, bank.Status).Scan(&bank.BankID)
	if err != nil {
		return fmt.Errorf("repository.Bank.Insert.QueryRowContext: %w", err)
	}

	return nil
}

func (repo *Bank) Update(ctx context.Context, bank *entity.Bank) error {
	query := `
		UPDATE bank
		SET name = $1, bank_code = $2, status = $3, updated_at = CURRENT_TIMESTAMP
		WHERE bank_id = $4
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.Bank.Update.PrepareContext: %w", err)
	}

	_, err = stmt.ExecContext(ctx, bank.Name, bank.BankCode, bank.Status, bank.BankID)
	if err != nil {
		return fmt.Errorf("repository.Bank.Update.ExecContext: %w", err)
	}

	return nil
}

func (repo *Bank) Delete(ctx context.Context, bankID int) error {
	query := `
		UPDATE bank
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE bank_id = $1
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.Bank.Delete.PrepareContext: %w", err)
	}

	_, err = stmt.ExecContext(ctx, bankID)
	if err != nil {
		return fmt.Errorf("repository.Bank.Delete.ExecContext: %w", err)
	}

	return nil
}

func (repo *Bank) ListAll(ctx context.Context) ([]*entity.Bank, error) {
	query := `
		SELECT bank_id, name, bank_code, status, created_at, updated_at, deleted_at
		FROM bank
		WHERE deleted_at IS NULL
	`

	rows, err := repo.cli.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Bank.ListAll.QueryContext: %w", err)
	}
	defer rows.Close()

	var banks []*entity.Bank
	for rows.Next() {
		var bank entity.Bank
		if err := rows.Scan(&bank.BankID, &bank.Name, &bank.BankCode, &bank.Status, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt); err != nil {
			return nil, fmt.Errorf("repository.Bank.ListAll.Scan: %w", err)
		}
		banks = append(banks, &bank)
	}

	return banks, nil
}

func (repo *Bank) ListByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error) {
	query := `
		SELECT bank_id, name, bank_code, status, created_at, updated_at, deleted_at
		FROM bank
		WHERE status = $1 AND deleted_at IS NULL
	`

	rows, err := repo.cli.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("repository.Bank.ListByStatus.QueryContext: %w", err)
	}
	defer rows.Close()

	var banks []*entity.Bank
	for rows.Next() {
		var bank entity.Bank
		if err := rows.Scan(&bank.BankID, &bank.Name, &bank.BankCode, &bank.Status, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt); err != nil {
			return nil, fmt.Errorf("repository.Bank.ListByStatus.Scan: %w", err)
		}
		banks = append(banks, &bank)
	}

	return banks, nil
}

func (repo *Bank) IsBankCodeExist(ctx context.Context, bankCode string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM bank
			WHERE bank_code = $1 AND deleted_at IS NULL
		)
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("repository.Bank.IsBankCodeExist.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, bankCode)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("repository.Bank.IsBankCodeExist.Scan: %w", err)
	}

	return exists, nil
}

func (repo *Bank) IsBankExist(ctx context.Context, bankID int) (bool, error) {

	return false, nil
}
