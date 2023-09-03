package request

import (
	"errors"
	"regexp"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type RegisterFinancialCard struct {
	AccountID      int
	CardNumber     string
	CardType       enum.FinancialCardType
	ExpirationDate time.Time
	CardHolderName string
	CVV            string
	Status         enum.FinancialCardStatus
}

func (r *RegisterFinancialCard) Validate() error {
	if r.AccountID <= 0 {
		return errors.New("invalid AccountID")
	}

	if len(r.CardNumber) < 16 || len(r.CardNumber) > 19 {
		return errors.New("invalid CardNumber")
	}

	if !enum.IsValidFinancialCardType(r.CardType) { // Assuming you have a function to validate the enum
		return errors.New("invalid CardType")
	}

	if time.Now().After(r.ExpirationDate) {
		return errors.New("invalid ExpirationDate")
	}

	if len(r.CardHolderName) > 100 {
		return errors.New("invalid CardHolderName")
	}

	isCVVValid, _ := regexp.MatchString(`^\d{3,4}$`, r.CVV)
	if !isCVVValid {
		return errors.New("invalid CVV")
	}

	if !enum.IsValidFinancialCardStatus(r.Status) { // Assuming you have a function to validate the enum
		return errors.New("invalid Status")
	}

	return nil
}

type UpdateFinancialCard struct {
	CardID         int64
	CardNumber     string
	CardType       enum.FinancialCardType
	ExpirationDate time.Time
	CardHolderName string
	CVV            string
	Status         enum.FinancialCardStatus
}

func (u *UpdateFinancialCard) Validate() error {
	// Validate CardID
	if u.CardID <= 0 {
		return errors.New("invalid CardID")
	}

	// Validate CardNumber
	if len(u.CardNumber) < 16 || len(u.CardNumber) > 19 {
		return errors.New("invalid CardNumber length")
	}

	// Validate CardType
	if !enum.IsValidFinancialCardType(u.CardType) {
		return errors.New("invalid CardType")
	}

	// Validate ExpirationDate
	if u.ExpirationDate.Before(time.Now()) {
		return errors.New("invalid ExpirationDate")
	}

	// Validate CardHolderName
	if len(u.CardHolderName) == 0 || len(u.CardHolderName) > 100 {
		return errors.New("invalid CardHolderName length")
	}

	// Validate CVV
	if len(u.CVV) != 3 && len(u.CVV) != 4 {
		return errors.New("invalid CVV")
	}

	// Validate Status
	if !enum.IsValidFinancialCardStatus(u.Status) {
		return errors.New("invalid Status")
	}

	return nil
}
