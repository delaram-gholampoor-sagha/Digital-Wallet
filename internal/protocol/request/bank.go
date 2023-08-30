package request

import (
	"errors"
	"regexp"
	"strings"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type UpdateBank struct {
	BankID   int64
	Name     string
	BankCode string
	Status   enum.BankStatus
}

func (u *UpdateBank) Validate() error {
	if u.BankID == 0 {
		return errors.New("BankID must be set")
	}
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("Name must not be empty")
	}
	if len(u.Name) > 100 {
		return errors.New("Name must not exceed 100 characters")
	}
	if strings.TrimSpace(u.BankCode) == "" {
		return errors.New("BankCode must not be empty")
	}
	if len(u.BankCode) > 10 {
		return errors.New("BankCode must not exceed 10 characters")
	}
	if !isValidStatus(u.Status) {
		return errors.New("Invalid Status")
	}
	return nil
}

// Assuming enum.BankStatus is an enum of string types
func isValidStatus(status enum.BankStatus) bool {
	return status == enum.BankActive || status == enum.BankInactive
}

type RegisterBank struct {
	Name     string
	BankCode string
	Status   enum.BankStatus
}

func (r *RegisterBank) Validate() error {
	// Validate Name: Must not be empty and must be less than 100 characters
	if r.Name == "" || len(r.Name) > 100 {
		return errors.New("invalid Name: must not be empty and must be less than 100 characters")
	}

	// Validate BankCode: Must not be empty and must be less than 10 characters
	if r.BankCode == "" || len(r.BankCode) > 10 {
		return errors.New("invalid BankCode: must not be empty and must be less than 10 characters")
	}

	// regular expression  to enforce a specific format for BankCode
	var re = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(r.BankCode) {
		return errors.New("invalid BankCode: must only contain alphanumeric characters")
	}

	// Validate Status: Must be either 'active' or 'inactive'
	if r.Status != enum.BankActive && r.Status != enum.BankInactive {
		return errors.New("invalid Status: must be either 'active' or 'inactive'")
	}

	return nil
}
