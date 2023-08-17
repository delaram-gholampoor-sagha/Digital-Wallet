package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type AccountTransactionService interface {
	RegisterTransaction(ctx context.Context, req *request.RegisterTransactionRequest) (*response.RegisterTransactionResponse, error)
	DeleteTransaction(ctx context.Context, transactionID int64) error
	GetTransactionByID(ctx context.Context, transactionID int64) (*entity.FinancialAccountTransaction, error)
	ListTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialAccountTransaction, error)
	ListTransactionsByGroupID(ctx context.Context, groupID int) ([]*entity.FinancialAccountTransaction, error)
	CancelTransaction(ctx context.Context, transactionID int64) error
	Transfer(ctx context.Context, req request.TransferRequest) (res response.TransferResponse, err error)
}

type AccountTransactionRepository interface {
	Insert(ctx context.Context, transaction *entity.FinancialAccountTransaction) error
	Delete(ctx context.Context, transactionID int64) error
	GetByID(ctx context.Context, transactionID int64) (*entity.FinancialAccountTransaction, error)
	ListByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialAccountTransaction, error)
	ListByTransactionGroupID(ctx context.Context, groupID int) ([]*entity.FinancialAccountTransaction, error)

	BeginTx(ctx context.Context) (context.Context, error) // Begins a new transaction
	CommitTx(ctx context.Context) error                   // Commits the transaction
	RollbackTx(ctx context.Context) error                 // Rollbacks the transaction

	// Transfer related methods
	DebitAccount(ctx context.Context, accountID int, amount float64) error
	CreditAccount(ctx context.Context, accountID int, amount float64) error
}
