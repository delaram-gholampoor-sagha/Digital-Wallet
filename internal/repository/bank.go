package repository

import (
	"context"
	"database/sql"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type Bank struct {
	cli *sql.DB
}


func NewBankRepository(cli *sql.DB) *Bank {
	return &Bank{cli}
}


func (b *Bank) GetByID(ctx context.Context, bankID int) (*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}

func (b *Bank) GetByCode(ctx context.Context, bankCode string) (*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}


func (b *Bank) GetByName(ctx context.Context, bankName string) (*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}

func (b *Bank) Insert(ctx context.Context, bank *entity.Bank) error {
	// TODO: Implement
	return nil
}

func (b *Bank) Update(ctx context.Context, bank *entity.Bank) error {
	// TODO: Implement
	return nil
}


func (b *Bank) Delete(ctx context.Context, bankID int) error {
	// TODO: Implement
	return nil
}


func (b *Bank) ListAll(ctx context.Context) ([]*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}


func (b *Bank) ListByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}


func (b *Bank) IsBankCodeExist(ctx context.Context, bankCode string) (bool, error) {
	// TODO: Implement
	return false, nil
}
