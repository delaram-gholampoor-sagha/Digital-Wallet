package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type FinancialCard struct {
	CardID         int
	AccountID      int
	CardNumber     string
	CardType       enum.FinancialCardType
	ExpirationDate time.Time
	CardHolderName string
	CVV            string
	Status         enum.FinancialCardStatus
	IssuedDate     time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}
