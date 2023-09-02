package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type Bank interface {
	RegisterBank(ctx context.Context, req request.RegisterBank) (response.RegisterBank, error)
	GetBankByID(ctx context.Context, bankID int) (response.GetBank, error)
	GetBankByCode(ctx context.Context, bankCode string) (response.GetBank, error)
	GetBankByName(ctx context.Context, bankName string) (response.GetBank, error)
	IsBankExist(ctx context.Context, bankID int) (bool, error)
	UpdateBankDetails(ctx context.Context, req request.UpdateBank) error
	ListAllBanks(ctx context.Context) ([]*entity.Bank, error)
	ListBanksByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error)
}

type BankRepository interface {
	GetByID(ctx context.Context, bankID int) (*entity.Bank, error)
	GetByCode(ctx context.Context, bankCode string) (*entity.Bank, error)
	GetByName(ctx context.Context, bankName string) (*entity.Bank, error)
	Insert(ctx context.Context, bank *entity.Bank) error
	Update(ctx context.Context, bank *entity.Bank) error
	Delete(ctx context.Context, bankID int) error
	ListAll(ctx context.Context) ([]*entity.Bank, error)
	ListByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error)
	IsBankCodeExist(ctx context.Context, bankCode string) (bool, error)
	IsBankExist(ctx context.Context, bankID int) (bool, error)
}
