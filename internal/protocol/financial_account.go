package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type FinancialAccount interface {
	CreateAccount(ctx context.Context, req request.RegisterFinancialAccount) (response.RegisterFinancialAccount, error)
	GetAccountByID(ctx context.Context, accountID int) (response.GetFinancialAccount, error)
	IsAccountExist(ctx context.Context, accountID int) (bool, error)
	UpdateAccount(ctx context.Context, req request.UpdateFinancialAccount) error
	DeleteAccount(ctx context.Context, accountID int) error
	ListAccountsByUserID(ctx context.Context, userID int) ([]entity.FinancialAccount, error)
	ListAccountsByStatus(ctx context.Context, status enum.FinancialAccountStatus) ([]*entity.FinancialAccount, error)
	GetAccountStatus(ctx context.Context, accountID int) (enum.FinancialAccountStatus, error)

	VerifyAccount(ctx context.Context, accountID int) error
	GetAccountByShaba(ctx context.Context, shabaNumber string) (response.GetFinancialAccount, error)
	ListAccountsByType(ctx context.Context, accountType string) ([]*entity.FinancialAccount, error)
	GetAccountCurrency(ctx context.Context, accountID int) (response.GetCurrency, error)
	GetBranchForAccount(ctx context.Context, accountID int) (response.GetBankBranch, error)
	GetBankForAccount(ctx context.Context, accountID int) (response.GetBank, error)
}

type FinancialAccountRepository interface {
	GetByID(ctx context.Context, accountID int) (*entity.FinancialAccount, error)
	GetByUserID(ctx context.Context, userID int) ([]entity.FinancialAccount, error)
	IsAccountExist(ctx context.Context, accountID int) (bool, error)
	Insert(ctx context.Context, account *entity.FinancialAccount) error
	Update(ctx context.Context, account *entity.FinancialAccount) error
	Delete(ctx context.Context, accountID int) error
	ListByStatus(ctx context.Context, status enum.FinancialAccountStatus) ([]*entity.FinancialAccount, error)
	UpdateStatus(ctx context.Context, accountID int, status string) error
	GetAccountStatus(ctx context.Context, accountID int) (enum.FinancialAccountStatus, error)

	GetByShabaNumber(ctx context.Context, shabaNumber string) (*entity.FinancialAccount, error)
	GetTransactions(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error)
	ListByType(ctx context.Context, accountType string) ([]*entity.FinancialAccount, error)
	GetCurrencyByAccountID(ctx context.Context, accountID int) (*entity.Currency, error)
	GetBranchByAccountID(ctx context.Context, accountID int) (*entity.BankBranch, error)
	GetBankByAccountID(ctx context.Context, accountID int) (*entity.Bank, error)
}
