package request

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type AddBranch struct {
	BankID      int
	BranchName  string
	BranchCode  string
	Address     string
	City        string
	Province    string
	PostalCode  string
	PhoneNumber string
	Status      enum.BankBranchStatus
}

func (req *AddBranch) Validate() error {
	// Check if BankID is provided and valid
	if req.BankID <= 0 {
		return errors.New("Invalid Bank ID")
	}

	// Check if BankID is provided and valid
	if req.BankID <= 0 {
		return errors.New("Invalid Bank ID")
	}

	// Check branch name, max length 100
	if len(req.BranchName) == 0 || len(req.BranchName) > 100 {
		return errors.New("Branch name must be between 1 and 100 characters")
	}

	// Check branch code, max length 10
	if len(req.BranchCode) > 10 {
		return errors.New("Branch code must be 10 characters or fewer")
	}

	// Check address, max length 255
	if len(req.Address) > 255 {
		return errors.New("Address must be 255 characters or fewer")
	}

	// Check city, max length 50
	if len(req.City) > 50 {
		return errors.New("City must be 50 characters or fewer")
	}

	// Check province, max length 50
	if len(req.Province) > 50 {
		return errors.New("Province must be 50 characters or fewer")
	}

	// Check postal code, max length 10
	if len(req.PostalCode) > 10 {
		return errors.New("Postal code must be 10 characters or fewer")
	}

	// Check phone number, max length 20
	if len(req.PhoneNumber) > 20 {
		return errors.New("Phone number must be 20 characters or fewer")
	}

	// Check if the phone number contains only digits, spaces, or hyphens
	if matched, _ := regexp.MatchString("^[0-9-\\s]+$", req.PhoneNumber); !matched {
		return errors.New("Phone number can contain only digits, spaces, and hyphens")
	}

	// Check if status is either 'active' or 'inactive'
	if req.Status != enum.BankBranchActive && req.Status != enum.BankBranchInactive {
		return errors.New("Status must be either 'active' or 'inactive'")
	}

	return nil
}

type UpdateBranch struct {
	BranchID    int
	BankID      int // if you want to allow moving branches between banks
	BranchName  string
	BranchCode  string
	Address     string
	City        string
	Province    string
	PostalCode  string
	PhoneNumber string
	Status      enum.BankBranchStatus
}

func (ub *UpdateBranch) Validate() error {
	if ub.BranchID <= 0 {
		return errors.New("BranchID must be greater than zero")
	}

	if ub.BankID <= 0 {
		return errors.New("BankID must be greater than zero")
	}

	if utf8.RuneCountInString(ub.BranchName) == 0 || utf8.RuneCountInString(ub.BranchName) > 100 {
		return errors.New("BranchName must be between 1 and 100 characters")
	}

	if utf8.RuneCountInString(ub.BranchCode) == 0 || utf8.RuneCountInString(ub.BranchCode) > 10 {
		return errors.New("BranchCode must be between 1 and 10 characters")
	}

	if utf8.RuneCountInString(ub.Address) > 255 {
		return errors.New("Address must be 255 characters or less")
	}

	if utf8.RuneCountInString(ub.City) > 50 {
		return errors.New("City must be 50 characters or less")
	}

	if utf8.RuneCountInString(ub.Province) > 50 {
		return errors.New("Province must be 50 characters or less")
	}

	if utf8.RuneCountInString(ub.PostalCode) > 10 {
		return errors.New("PostalCode must be 10 characters or less")
	}

	if utf8.RuneCountInString(ub.PhoneNumber) > 20 {
		return errors.New("PhoneNumber must be 20 characters or less")
	}

	if ub.Status != enum.BankBranchActive && ub.Status != enum.BankBranchInactive {
		return fmt.Errorf("Status must be either 'active' or 'inactive', got %v", ub.Status)
	}

	return nil
}
