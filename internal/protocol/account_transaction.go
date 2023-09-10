package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type AccountTransaction interface {
	RegisterTransaction(ctx context.Context, req *request.RegisterTransactionRequest) (*response.RegisterTransactionResponse, error)
	DeleteTransaction(ctx context.Context, transactionID int64) error
	GetTransactionByID(ctx context.Context, transactionID int64) (*entity.AccountTransaction, error)
	ListTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error)
	ListTransactionsByGroupID(ctx context.Context, groupID int) ([]*entity.AccountTransaction, error)
	CancelTransaction(ctx context.Context, transactionID int64) error
	Transfer(ctx context.Context, req request.TransferRequest) (res response.TransferResponse, err error)
	GetAccountTransactionHistory(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error)
}

type AccountTransactionRepository interface {
	Insert(ctx context.Context, transaction *entity.AccountTransaction) error
	Delete(ctx context.Context, transactionID int64) error
	GetByID(ctx context.Context, transactionID int64) (*entity.AccountTransaction, error)
	ListByAccountID(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error)
	ListByTransactionGroupID(ctx context.Context, groupID int) ([]*entity.AccountTransaction, error)
	CreateNewTransactionGroup(ctx context.Context) (int, error)
	Update(ctx context.Context, transaction *entity.AccountTransaction) error
	GetCurrentBalance(ctx context.Context, accountID int) (float64, error)

	BeginTx(ctx context.Context) (context.Context, error) // Begins a new transaction
	CommitTx(ctx context.Context) error                   // Commits the transaction
	RollbackTx(ctx context.Context) error                 // Rollbacks the transaction

	// Transfer related methods
	DebitAccount(ctx context.Context, accountID int, amount float64) error
	CreditAccount(ctx context.Context, accountID int, amount float64) error
}
