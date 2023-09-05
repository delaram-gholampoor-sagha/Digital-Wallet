package request

import (
	"errors"
	"regexp"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

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

func (req *RegisterFinancialAccount) Validate() error {
	// Check if UserID is provided and valid
	if req.UserID <= 0 {
		return errors.New("Invalid User ID")
	}

	// Check if CurrencyID is provided and valid
	if req.CurrencyID <= 0 {
		return errors.New("Invalid Currency ID")
	}

	// Check if BankID is provided and valid
	if req.BankID <= 0 {
		return errors.New("Invalid Bank ID")
	}

	// Check if BranchID is provided and valid
	if req.BranchID <= 0 {
		return errors.New("Invalid Branch ID")
	}

	// Check AccountNumber, length should be 20
	if len(req.AccountNumber) != 20 {
		return errors.New("Account number must be exactly 20 characters")
	}

	// Check ShabaNumber, length should be 26
	if len(req.ShabaNumber) != 26 {
		return errors.New("Shaba number must be exactly 26 characters")
	}

	// Check AccountName, max length 100
	if len(req.AccountName) == 0 || len(req.AccountName) > 100 {
		return errors.New("Account name must be between 1 and 100 characters")
	}

	// Validate AccountType
	if req.AccountType != enum.Checking && req.AccountType != enum.Savings {
		return errors.New("Account type must be either 'checking' or 'savings'")
	}

	// Check CurrencyCode, length should be 3
	if len(req.CurrencyCode) != 3 {
		return errors.New("Currency code must be exactly 3 characters")
	}

	// Validate Status
	if req.Status != enum.Verified && req.Status != enum.Unverified {
		return errors.New("Status must be either 'verified' or 'unverified'")
	}

	return nil
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

func (u *UpdateFinancialAccount) Validate() error {
	// Check for non-zero or negative IDs
	if u.AccountID <= 0 || u.UserID <= 0 || u.CurrencyID <= 0 || u.BankID <= 0 || u.BranchID <= 0 {
		return errors.New("IDs must be greater than zero")
	}

	// Check for AccountNumber length
	if len(u.AccountNumber) == 0 || len(u.AccountNumber) > 20 {
		return errors.New("AccountNumber must be between 1 to 20 characters")
	}

	// Check for ShabaNumber length and uniqueness (you can check uniqueness in the DB)
	if len(u.ShabaNumber) != 26 {
		return errors.New("ShabaNumber must be exactly 26 characters")
	}

	// Regex for ShabaNumber format validation, if you have a specific format
	match, _ := regexp.MatchString("fiiiil", u.ShabaNumber)
	if !match {
		return errors.New("ShabaNumber format is invalid")
	}

	// Check for AccountName length
	if len(u.AccountName) > 100 {
		return errors.New("AccountName must not exceed 100 characters")
	}

	// Check for valid AccountType
	if u.AccountType != enum.Checking && u.AccountType != enum.Savings {
		return errors.New("Invalid AccountType")
	}

	// Check for valid CurrencyCode length
	if len(u.CurrencyCode) != 3 {
		return errors.New("CurrencyCode must be exactly 3 characters")
	}

	// Check for valid status
	if u.Status != enum.Verified && u.Status != enum.Unverified {
		return errors.New("Invalid status")
	}

	return nil
}
