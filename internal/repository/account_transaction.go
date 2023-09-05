package repository

import (
	"context"
	"database/sql"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
)

type AccountTransaction struct {
	cli *sql.DB
}

func (repo *AccountTransaction) Insert(ctx context.Context, transaction *entity.FinancialAccountTransaction) error {
	// TODO: Implement
	return nil
}

func (repo *AccountTransaction) Delete(ctx context.Context, transactionID int64) error {
	// TODO: Implement
	return nil
}

func (repo *AccountTransaction) GetByID(ctx context.Context, transactionID int64) (*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
}

func (repo *AccountTransaction) ListByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
}

func (repo *AccountTransaction) ListByTransactionGroupID(ctx context.Context, groupID int) ([]*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
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
	// TODO: Implement
	return nil
}

func (repo *AccountTransaction) CreditAccount(ctx context.Context, accountID int, amount float64) error {
	// TODO: Implement
	return nil
}
