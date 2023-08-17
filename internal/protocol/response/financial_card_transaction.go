package response

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type RegisterCardTransactionResponse struct {
	TransactionID int
	Status        enum.FinancialCardTransactionStatus
	CreatedAt     time.Time
}

type Transfer struct {
	SenderTx   entity.FinancialCardTransaction
	ReceiverTx entity.FinancialCardTransaction
}
