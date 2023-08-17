package request

import (
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

type UpdateFinancialCard struct {
	CardID         int64
	CardNumber     string
	CardType       enum.FinancialCardType
	ExpirationDate time.Time
	CardHolderName string
	CVV            string
	Status         enum.FinancialCardStatus
}
