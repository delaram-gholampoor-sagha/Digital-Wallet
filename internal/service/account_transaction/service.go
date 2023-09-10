package accounttransaction

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"go.uber.org/zap"
)

type Service struct {
	cfg                     config.JWT
	logger                  *zap.SugaredLogger
	tokenGen                protocol.TokenGenerator
	financialAccountService protocol.FinancialAccount
	accountTransactionRepo  protocol.AccountTransactionRepository
}

func New(
	cfg config.JWT,
	logger *zap.SugaredLogger,
	tokenGen protocol.TokenGenerator,
	financialAccountService protocol.FinancialAccount,
	accountTransactionRepo protocol.AccountTransactionRepository,

) *Service {
	return &Service{
		cfg:                     cfg,
		logger:                  logger,
		tokenGen:                tokenGen,
		financialAccountService: financialAccountService,
		accountTransactionRepo:  accountTransactionRepo,
	}
}
