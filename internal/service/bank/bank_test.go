package bank

import (
	"context"
	"testing"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func setup() (*Service, *protocol.MockBankRepo) {
	mockRepo := new(protocol.MockBankRepo)

	cfg := config.JWT{
		AccessTokenExp:  time.Minute * 15,
		RefreshTokenExp: time.Hour * 24,
		Secret:          "testSecret",
	}

	tokenGen := utils.JWTTokenGenerator{}

	logger, _ := zap.NewProduction()
	sugaredLogger := logger.Sugar()

	service := &Service{
		bankRepo: mockRepo,
		tokenGen: &tokenGen,
		cfg:      cfg,
		logger:   sugaredLogger,
	}
	return service, mockRepo
}

func TestRegisterBank(t *testing.T) {
	ctx := context.Background()
	service, mockRepo := setup()

	t.Run("Successfully register bank", func(t *testing.T) {
		req := request.RegisterBank{
			Name:     "Test Bank",
			BankCode: "TST",
		}

		mockRepo.On("Insert", ctx, mock.AnythingOfType("*entity.Bank")).Return(nil)

		_, err := service.RegisterBank(ctx, req)
		assert.NoError(t, err)
	})

}

func TestGetBankByID(t *testing.T) {
	ctx := context.Background()
	service, mockRepo := setup()

	t.Run("Successfully get bank by ID", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, 1).Return(&entity.Bank{BankID: 1, Name: "Test Bank", BankCode: "TST"}, nil)

		res, err := service.GetBankByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, "Test Bank", res.Name)
	})

}

func TestGetBankByCode(t *testing.T) {
	ctx := context.Background()
	service, mockRepo := setup()

	t.Run("Successfully get bank by Code", func(t *testing.T) {
		mockRepo.On("GetByCode", ctx, "TST").Return(&entity.Bank{BankID: 1, Name: "Test Bank", BankCode: "TST"}, nil)

		res, err := service.GetBankByCode(ctx, "TST")
		assert.NoError(t, err)
		assert.Equal(t, "Test Bank", res.Name)
	})

}

func TestUpdateBankDetails(t *testing.T) {
	ctx := context.Background()
	service, mockRepo := setup()

	t.Run("Successfully update bank details", func(t *testing.T) {
		req := request.UpdateBank{
			BankID:   1,
			Name:     "Updated Test Bank",
			BankCode: "UTST",
		}

		mockRepo.On("Update", ctx, mock.AnythingOfType("*entity.Bank")).Return(nil)

		err := service.UpdateBankDetails(ctx, req)
		assert.NoError(t, err)
	})

}

func TestListAllBanks(t *testing.T) {
	ctx := context.Background()
	service, mockRepo := setup()

	t.Run("Successfully list all banks", func(t *testing.T) {
		mockRepo.On("ListAll", ctx).Return([]*entity.Bank{{BankID: 1, Name: "Bank1", BankCode: "B1"}}, nil)

		res, err := service.ListAllBanks(ctx)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, "Bank1", res[0].Name)
	})

}

func TestListBanksByStatus(t *testing.T) {
	ctx := context.Background()
	service, mockRepo := setup()

	t.Run("Successfully list banks by status", func(t *testing.T) {
		status := enum.BankActive
		mockRepo.On("ListByStatus", ctx, status).Return([]*entity.Bank{{BankID: 1, Name: "ActiveBank", BankCode: "AB", Status: enum.BankActive}}, nil)

		res, err := service.ListBanksByStatus(ctx, status)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		assert.Equal(t, "ActiveBank", res[0].Name)
	})

}
