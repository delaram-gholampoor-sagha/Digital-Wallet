package financialaccount

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

func (s *Service) CreateAccount(ctx context.Context, req request.RegisterFinancialAccount) (response.RegisterFinancialAccount, error) {
	// TODO: Implement
	return response.RegisterFinancialAccount{}, nil
}

func (s *Service) GetAccountByID(ctx context.Context, accountID int) (response.GetFinancialAccount, error) {
	// TODO: Implement
	return response.GetFinancialAccount{}, nil
}

func (s *Service) IsAccountExist(ctx context.Context, accountID int) (bool, error) {
	// TODO: Implement
	return false, nil
}

func (s *Service) UpdateAccount(ctx context.Context, req request.UpdateFinancialAccount) error {
	// TODO: Implement
	return nil
}

func (s *Service) DeleteAccount(ctx context.Context, accountID int) error {
	// TODO: Implement
	return nil
}

func (s *Service) ListAccountsByUserID(ctx context.Context, userID int) ([]*entity.FinancialAccount, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) ListAccountsByStatus(ctx context.Context, status enum.FinancialAccountStatus) ([]*entity.FinancialAccount, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) VerifyAccount(ctx context.Context, accountID int) error {
	// TODO: Implement
	return nil
}

func (s *Service) GetAccountByShaba(ctx context.Context, shabaNumber string) (response.GetFinancialAccount, error) {
	// TODO: Implement
	return response.GetFinancialAccount{}, nil
}

func (s *Service) GetAccountTransactionHistory(ctx context.Context, accountID int) ([]*entity.FinancialAccountTransaction, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) ListAccountsByType(ctx context.Context, accountType string) ([]*entity.FinancialAccount, error) {
	// TODO: Implement
	return nil, nil
}

func (s *Service) GetAccountCurrency(ctx context.Context, accountID int) (response.GetCurrency, error) {
	// TODO: Implement
	return response.GetCurrency{}, nil
}

func (s *Service) GetBranchForAccount(ctx context.Context, accountID int) (response.GetBankBranch, error) {
	// TODO: Implement
	return response.GetBankBranch{}, nil
}

func (s *Service) GetBankForAccount(ctx context.Context, accountID int) (response.GetBank, error) {
	// TODO: Implement
	return response.GetBank{}, nil
}
