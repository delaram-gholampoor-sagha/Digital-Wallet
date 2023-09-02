package bank

import (
	"context"
	"errors"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
)

func (b *Service) RegisterBank(ctx context.Context, req request.RegisterBank) (response.RegisterBank, error) {
	if err := req.Validate(); err != nil {
		return response.RegisterBank{}, err
	}

	exists, err := b.bankRepo.IsBankCodeExist(ctx, req.BankCode)
	if err != nil {
		b.logger.Errorw("service.bank.RegisterBank.bankRepo.IsBankCodeExist", "error", err)
		return response.RegisterBank{}, derror.NewInternalSystemError()
	}
	if exists {
		return response.RegisterBank{}, derror.NewBadRequestError("Bank code already exists.")
	}

	bank := entity.Bank{
		Name:     req.Name,
		BankCode: req.BankCode,
		Status:   req.Status,
	}
	if err := b.bankRepo.Insert(ctx, &bank); err != nil {
		b.logger.Errorw("Failed to insert bank", "error", err)
		return response.RegisterBank{}, derror.NewInternalSystemError()
	}
	return response.RegisterBank{}, nil
}

func (b *Service) GetBankByID(ctx context.Context, bankID int) (response.GetBank, error) {
	bank, err := b.bankRepo.GetByID(ctx, bankID)
	if err != nil {
		b.logger.Errorw("Failed to get bank by ID", "error", err)
		return response.GetBank{}, derror.NewInternalSystemError()
	}

	return response.GetBank{
		BankID:    bank.BankID,
		Name:      bank.Name,
		BankCode:  bank.BankCode,
		Status:    bank.Status,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
		DeletedAt: bank.DeletedAt,
	}, nil
}

func (b *Service) GetBankByCode(ctx context.Context, bankCode string) (response.GetBank, error) {
	if bankCode == "" {
		return response.GetBank{}, errors.New("bankCode cannot be empty")
	}

	bank, err := b.bankRepo.GetByCode(ctx, bankCode)
	if err != nil {
		b.logger.Errorw("Failed to retrieve bank by code", "bankCode", bankCode, "error", err.Error())
		return response.GetBank{}, derror.NewInternalSystemError()
	}

	if bank == nil {
		return response.GetBank{}, derror.NewNotFoundError(fmt.Sprintf("Bank with code %s not found", bankCode))
	}

	return response.GetBank{
		BankID:    bank.BankID,
		Name:      bank.Name,
		BankCode:  bank.BankCode,
		Status:    bank.Status,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
		DeletedAt: bank.DeletedAt,
	}, nil
}

func (b *Service) GetBankByName(ctx context.Context, bankName string) (response.GetBank, error) {
	if bankName == "" {
		return response.GetBank{}, errors.New("bankName cannot be empty")
	}

	bank, err := b.bankRepo.GetByName(ctx, bankName)
	if err != nil {
		b.logger.Errorw("Failed to retrieve bank by name", "bankName", bankName, "error", err.Error())
		return response.GetBank{}, derror.NewInternalSystemError()
	}

	if bank == nil {
		return response.GetBank{}, derror.NewNotFoundError(fmt.Sprintf("Bank with name %s not found", bankName))
	}

	return response.GetBank{
		BankID:    bank.BankID,
		Name:      bank.Name,
		BankCode:  bank.BankCode,
		Status:    bank.Status,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
		DeletedAt: bank.DeletedAt,
	}, nil
}

func (b *Service) UpdateBankDetails(ctx context.Context, req request.UpdateBank) error {

	if err := req.Validate(); err != nil {
		return err
	}

	bank := &entity.Bank{
		BankID:   req.BankID,
		Name:     req.Name,
		BankCode: req.BankCode,
		Status:   req.Status,
	}

	if err := b.bankRepo.Update(ctx, bank); err != nil {
		b.logger.Errorw("service.bank.UpdateBankDetails.bankRepo.Update", "error", err.Error())
		return derror.NewInternalSystemError()
	}
	return nil
}

func (b *Service) ListAllBanks(ctx context.Context) ([]*entity.Bank, error) {
	banks, err := b.bankRepo.ListAll(ctx)
	if err != nil {
		b.logger.Errorw("service.bank.ListAllBanks.bankRepo.ListAll", "error", err.Error())
		return nil, derror.NewInternalSystemError()
	}
	return banks, nil
}

func (b *Service) ListBanksByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error) {
	banks, err := b.bankRepo.ListByStatus(ctx, status)
	if err != nil {
		b.logger.Errorw("service.bank.ListBanksByStatus.bankRepo.ListByStatus", "error", err.Error())
		return nil, derror.NewInternalSystemError()
	}
	return banks, nil
}

func (b *Service) IsBankExist(ctx context.Context, bankID int) (bool, error) {

	return false, nil
}
