package bank

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)


func (b *Service) RegisterBank(ctx context.Context, req request.RegisterBank) (response.RegisterBank, error) {
	// TODO: Implement
	return response.RegisterBank{}, nil
}


func (b *Service) GetBankByID(ctx context.Context, bankID int) (response.GetBank, error) {
	// TODO: Implement
	return response.GetBank{}, nil
}


func (b *Service) GetBankByCode(ctx context.Context, bankCode string) (response.GetBank, error) {
	// TODO: Implement
	return response.GetBank{}, nil
}


func (b *Service) GetBankByName(ctx context.Context, bankName string) (response.GetBank, error) {
	// TODO: Implement
	return response.GetBank{}, nil
}


func (b *Service) UpdateBankDetails(ctx context.Context, req request.UpdateBank) error {
	// TODO: Implement
	return nil
}


func (b *Service) ListAllBanks(ctx context.Context) ([]*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}


func (b *Service) ListBanksByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error) {
	// TODO: Implement
	return nil, nil
}
