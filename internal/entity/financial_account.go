package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type FinancialAccount struct {
	AccountID     int64
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
