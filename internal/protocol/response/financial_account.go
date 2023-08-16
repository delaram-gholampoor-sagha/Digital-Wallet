package response

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type RegisterFinancialAccount struct {
	AccountID int
}

type GetFinancialAccount struct {
	AccountID     int
	UserID        int
	CurrencyID    int
	BankID        int
	BranchID      int
	AccountNumber string
	ShabaNumber   string
	AccountName   string
	AccountType   enum.FinancialAccountType
	CurrencyCode  string
	Status        enum.FinancialAccountStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

type GetCurrency struct {
	CurrencyID   int
	CurrencyCode string
	CurrencyName string
	Symbol       string
	ExchangeRate *float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type GetBankBranch struct {
	BranchID   int
	BankID     int
	BranchCode string
	BranchName string
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
