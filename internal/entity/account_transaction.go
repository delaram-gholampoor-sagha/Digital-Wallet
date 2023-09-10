package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type AccountTransaction struct {
	TransactionID      int
	TransactionGroupID int
	FinancialAccountID int
	Amount             float64
	Balance            float64
	Description        *string
	Status             enum.AccountTransactionStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
