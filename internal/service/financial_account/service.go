package financialaccount

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"go.uber.org/zap"
)

type Service struct {
	cfg                  config.JWT
	logger               *zap.SugaredLogger
	tokenGen             protocol.TokenGenerator
	bankService          protocol.Bank
	bankBranchService    protocol.BankBranch
	userService          protocol.User
	currencyService      protocol.Currency
	financialAccountRepo protocol.FinancialAccountRepository
}

func New(cfg config.JWT,
	logger *zap.SugaredLogger,
	tokenGen protocol.TokenGenerator,
	financialAccountRepo protocol.FinancialAccountRepository,
	bankService protocol.Bank,
	bankBranchService protocol.BankBranch,
	userService protocol.User,
	currencyService protocol.Currency,
) *Service {
	return &Service{
		cfg:                  cfg,
		logger:               logger,
		tokenGen:             tokenGen,
		financialAccountRepo: financialAccountRepo,
		bankService:          bankService,
		bankBranchService:    bankBranchService,
		userService:          userService,
		currencyService:      currencyService,
	}
}
