package request

import "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"

type RegisterFinancialAccount struct {
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
}

type UpdateFinancialAccount struct {
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
}
