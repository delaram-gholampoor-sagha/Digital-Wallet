package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type FinancialCardTransactionService interface {
	RegisterTransaction(ctx context.Context, request *request.RegisterCardTransaction) (*response.RegisterCardTransactionResponse, error)
	GetTransactionByID(ctx context.Context, transactionID int64) (*entity.FinancialCardTransaction, error)
	ListTransactionsByCardID(ctx context.Context, cardID int) ([]*entity.FinancialCardTransaction, error)
	ListTransactionsByGroupID(ctx context.Context, groupID int) ([]*entity.FinancialCardTransaction, error)
	Transfer(ctx context.Context, req request.Transfer) (res response.Transfer, err error)
	CancelTransaction(ctx context.Context, transactionID int64) error
}

type CardTransactionRepository interface {
	Insert(ctx context.Context, transaction *entity.FinancialCardTransaction) error
	GetByID(ctx context.Context, transactionID int64) (*entity.FinancialCardTransaction, error)
	ListByTransactionGroupID(ctx context.Context, groupID int) ([]*entity.FinancialAccountTransaction, error)
	ListByCardID(ctx context.Context, cardID int) ([]*entity.FinancialCardTransaction, error)

	BeginTx(ctx context.Context) (context.Context, error) // Begins a new transaction
	CommitTx(ctx context.Context) error                   // Commits the transaction
	RollbackTx(ctx context.Context) error                 // Rollbacks the transaction

	// Transfer related methods
	DebitAccount(ctx context.Context, accountID int, amount float64) error
	CreditAccount(ctx context.Context, accountID int, amount float64) error
}
