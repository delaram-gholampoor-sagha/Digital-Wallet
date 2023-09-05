package accounttransaction

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

func (s *Service) RegisterTransaction(ctx context.Context, req *request.RegisterTransactionRequest) (*response.RegisterTransactionResponse, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) DeleteTransaction(ctx context.Context, transactionID int64) error {
	// TODO: Implement
	return nil
}

func (s *Service) GetTransactionByID(ctx context.Context, transactionID int64) (*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) ListTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) ListTransactionsByGroupID(ctx context.Context, groupID int) ([]*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) CancelTransaction(ctx context.Context, transactionID int64) error {
	// TODO: Implement
	return nil
}

func (s *Service) Transfer(ctx context.Context, req request.TransferRequest) (res response.TransferResponse, err error) {
	// TODO: Implement
	return res, nil
}
