package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type FinancialAccountTransaction struct {
	TransactionID      int64
	TransactionGroupID int
	FinancialAccountID int64
	Amount             float64
	Balance            float64
	Description        string
	Status             enum.FinancialAccountTransactionStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
