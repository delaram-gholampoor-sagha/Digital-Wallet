package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type BankBranch struct {
	cli *sql.DB
}

func (repo *BankBranch) GetByID(ctx context.Context, branchID int) (*entity.BankBranch, error) {
	query := `
		SELECT
			branch_id, bank_id, branch_name, branch_code, address,
			city, province, postal_code, phone_number, status,
			created_at, updated_at
		FROM
			bank_branch
		WHERE
			branch_id = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.BankBranch.GetByID.PrepareContext: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, branchID)

	var branch entity.BankBranch
	if err := row.Scan(
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
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("BankBranch not found")
		}
		return nil, fmt.Errorf("repository.BankBranch.GetByID.Scan: %w", err)
	}

	return &branch, nil
}

func (repo *BankBranch) GetByName(ctx context.Context, branchName string) (*entity.BankBranch, error) {
	query := `
		SELECT
			branch_id, bank_id, branch_name, branch_code, address,
			city, province, postal_code, phone_number, status,
			created_at, updated_at
		FROM
			bank_branch
		WHERE
			branch_name = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.BankBranch.GetByName.PrepareContext: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, branchName)

	var branch entity.BankBranch
	if err := row.Scan(
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
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("BankBranch not found")
		}
		return nil, fmt.Errorf("repository.BankBranch.GetByName.Scan: %w", err)
	}

	return &branch, nil
}

func (repo *BankBranch) GetByCode(ctx context.Context, branchCode string) (*entity.BankBranch, error) {
	query := `
		SELECT
			branch_id, bank_id, branch_name, branch_code, address,
			city, province, postal_code, phone_number, status,
			created_at, updated_at
		FROM
			bank_branch
		WHERE
			branch_code = $1 AND deleted_at IS NULL
	`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.BankBranch.GetByCode.PrepareContext: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, branchCode)

	var branch entity.BankBranch
	if err := row.Scan(
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
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("BankBranch not found")
		}
		return nil, fmt.Errorf("repository.BankBranch.GetByCode.Scan: %w", err)
	}

	return &branch, nil
}

func (repo *BankBranch) Insert(ctx context.Context, branch *entity.BankBranch) error {
	query := `
		INSERT INTO bank_branch (bank_id, branch_name, branch_code, address,
			city, province, postal_code, phone_number, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING branch_id
	`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.BankBranch.Insert.PrepareContext: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		branch.BankID,
		branch.BranchName,
		branch.BranchCode,
		branch.Address,
		branch.City,
		branch.Province,
		branch.PostalCode,
		branch.PhoneNumber,
		branch.Status,
	).Scan(&branch.BranchID)
	if err != nil {
		return fmt.Errorf("repository.BankBranch.Insert.QueryRowContext: %w", err)
	}

	return nil
}

func (repo *BankBranch) Update(ctx context.Context, branch *entity.BankBranch) error {
	query := `
		UPDATE bank_branch
		SET bank_id = $1, branch_name = $2, branch_code = $3, address = $4,
			city = $5, province = $6, postal_code = $7, phone_number = $8, status = $9,
			updated_at = CURRENT_TIMESTAMP
		WHERE branch_id = $10
	`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.BankBranch.Update.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		branch.BankID,
		branch.BranchName,
		branch.BranchCode,
		branch.Address,
		branch.City,
		branch.Province,
		branch.PostalCode,
		branch.PhoneNumber,
		branch.Status,
		branch.BranchID,
	)
	if err != nil {
		return fmt.Errorf("repository.BankBranch.Update.ExecContext: %w", err)
	}

	return nil
}

func (repo *BankBranch) Delete(ctx context.Context, branchID int) error {
	query := "UPDATE bank_branch SET deleted_at = CURRENT_TIMESTAMP WHERE branch_id = $1"
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.BankBranch.Delete.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, branchID)
	if err != nil {
		return fmt.Errorf("repository.BankBranch.Delete.ExecContext: %w", err)
	}

	return nil
}

func (repo *BankBranch) ListAll(ctx context.Context) ([]*entity.BankBranch, error) {
	query := `
		SELECT branch_id, bank_id, branch_name, branch_code, address, city, province, 
			postal_code, phone_number, status, created_at, updated_at
		FROM bank_branch
		WHERE deleted_at IS NULL
	`
	rows, err := repo.cli.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.BankBranch.ListAll.QueryContext: %w", err)
	}
	defer rows.Close()

	var branches []*entity.BankBranch
	for rows.Next() {
		var branch entity.BankBranch
		if err := rows.Scan(
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
		); err != nil {
			return nil, fmt.Errorf("repository.BankBranch.ListAll.Scan: %w", err)
		}
		branches = append(branches, &branch)
	}

	return branches, nil
}

// ListByStatus retrieves all BankBranches with a given status
func (repo *BankBranch) ListByStatus(ctx context.Context, status enum.BankBranchStatus) ([]*entity.BankBranch, error) {
	query := `
		SELECT branch_id, bank_id, branch_name, branch_code, address, city, province, 
			postal_code, phone_number, status, created_at, updated_at
		FROM bank_branch
		WHERE status = $1 AND deleted_at IS NULL
	`
	rows, err := repo.cli.QueryContext(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("repository.BankBranch.ListByStatus.QueryContext: %w", err)
	}
	defer rows.Close()

	var branches []*entity.BankBranch
	for rows.Next() {
		var branch entity.BankBranch
		if err := rows.Scan(
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
		); err != nil {
			return nil, fmt.Errorf("repository.BankBranch.ListByStatus.Scan: %w", err)
		}
		branches = append(branches, &branch)
	}

	return branches, nil
}

func (repo *BankBranch) ListByBankID(ctx context.Context, bankID int) ([]*entity.BankBranch, error) {
	query := `
		SELECT branch_id, bank_id, branch_name, branch_code, address, city, province, 
			postal_code, phone_number, status, created_at, updated_at
		FROM bank_branch
		WHERE bank_id = $1 AND deleted_at IS NULL
	`
	rows, err := repo.cli.QueryContext(ctx, query, bankID)
	if err != nil {
		return nil, fmt.Errorf("repository.BankBranch.ListByBankID.QueryContext: %w", err)
	}
	defer rows.Close()

	var branches []*entity.BankBranch
	for rows.Next() {
		var branch entity.BankBranch
		if err := rows.Scan(
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
		); err != nil {
			return nil, fmt.Errorf("repository.BankBranch.ListByBankID.Scan: %w", err)
		}
		branches = append(branches, &branch)
	}

	return branches, nil
}

func (repo *BankBranch) IsBranchCodeExist(ctx context.Context, branchCode string) (bool, error) {
	query := "SELECT COUNT(*) FROM bank_branch WHERE branch_code = $1 AND deleted_at IS NULL"
	var count int

	err := repo.cli.QueryRowContext(ctx, query, branchCode).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("repository.BankBranch.IsBranchCodeExist.QueryRowContext: %w", err)
	}

	return count > 0, nil
}
