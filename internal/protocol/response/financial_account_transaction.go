package response

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type RegisterTransactionResponse struct {
	TransactionID int
	Status        enum.FinancialAccountTransactionStatus
	CreatedAt     time.Time
}

type TransferResponse struct {
	SenderTx   entity.FinancialAccountTransaction
	ReceiverTx entity.FinancialAccountTransaction
}
