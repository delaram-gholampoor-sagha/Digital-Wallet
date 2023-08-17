package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type FinancialCardTransaction struct {
	TransactionID      int64
	TransactionGroupID int
	FinancialCardID    int
	Amount             float64
	Balance            float64
	Description        string
	Status             enum.FinancialCardTransactionStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
