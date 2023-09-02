package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/stretchr/testify/mock"
)

type MockBankRepo struct {
	mock.Mock
}

func (m *MockBankRepo) GetByID(ctx context.Context, bankID int) (*entity.Bank, error) {
	args := m.Called(ctx, bankID)
	return args.Get(0).(*entity.Bank), args.Error(1)
}

func (m *MockBankRepo) GetByCode(ctx context.Context, bankCode string) (*entity.Bank, error) {
	args := m.Called(ctx, bankCode)
	return args.Get(0).(*entity.Bank), args.Error(1)
}

func (m *MockBankRepo) GetByName(ctx context.Context, bankName string) (*entity.Bank, error) {
	args := m.Called(ctx, bankName)
	return args.Get(0).(*entity.Bank), args.Error(1)
}

func (m *MockBankRepo) Insert(ctx context.Context, bank *entity.Bank) error {
	args := m.Called(ctx, bank)
	return args.Error(0)
}

func (m *MockBankRepo) Update(ctx context.Context, bank *entity.Bank) error {
	args := m.Called(ctx, bank)
	return args.Error(0)
}

func (m *MockBankRepo) Delete(ctx context.Context, bankID int) error {
	args := m.Called(ctx, bankID)
	return args.Error(0)
}

func (m *MockBankRepo) ListAll(ctx context.Context) ([]*entity.Bank, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Bank), args.Error(1)
}

func (m *MockBankRepo) ListByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error) {
	args := m.Called(ctx, status)
	return args.Get(0).([]*entity.Bank), args.Error(1)
}

func (m *MockBankRepo) IsBankCodeExist(ctx context.Context, bankCode string) (bool, error) {
	args := m.Called(ctx, bankCode)
	return args.Bool(0), args.Error(1)
}


func (m *MockBankRepo) IsBankExist(ctx context.Context, bankID int) (bool, error) {
	args := m.Called(ctx, bankID)
	return args.Bool(0), args.Error(1)
}

type MockBankService struct {
	mock.Mock
}

func (m *MockBankService) RegisterBank(ctx context.Context, req request.RegisterBank) (response.RegisterBank, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(response.RegisterBank), args.Error(1)
}

func (m *MockBankService) GetBankByID(ctx context.Context, bankID int) (response.GetBank, error) {
	args := m.Called(ctx, bankID)
	return args.Get(0).(response.GetBank), args.Error(1)
}

func (m *MockBankService) GetBankByCode(ctx context.Context, bankCode string) (response.GetBank, error) {
	args := m.Called(ctx, bankCode)
	return args.Get(0).(response.GetBank), args.Error(1)
}

func (m *MockBankService) GetBankByName(ctx context.Context, bankName string) (response.GetBank, error) {
	args := m.Called(ctx, bankName)
	return args.Get(0).(response.GetBank), args.Error(1)
}

func (m *MockBankService) UpdateBankDetails(ctx context.Context, req request.UpdateBank) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockBankService) ListAllBanks(ctx context.Context) ([]*entity.Bank, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Bank), args.Error(1)
}

func (m *MockBankService) ListBanksByStatus(ctx context.Context, status enum.BankStatus) ([]*entity.Bank, error) {
	args := m.Called(ctx, status)
	return args.Get(0).([]*entity.Bank), args.Error(1)
}

func (m *MockBankService) IsBankExist(ctx context.Context, bankID int) (bool, error) {
	args := m.Called(ctx, bankID)
	return args.Bool(0), args.Error(1)
}
