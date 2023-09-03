package financialcard

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"go.uber.org/zap"
)

type Service struct {
	cfg                     config.JWT
	logger                  *zap.SugaredLogger
	financialCardRepo       protocol.FinancialCardRepository
	financialAccountService protocol.FinancialAccount
	tokenGen                protocol.TokenGenerator
}

func New(cfg config.JWT, logger *zap.SugaredLogger, financialCard protocol.FinancialCardRepository, tokenGen protocol.TokenGenerator, FinancialAccount protocol.FinancialAccount) *Service {
	return &Service{
		cfg:                     cfg,
		logger:                  logger,
		financialAccountService: FinancialAccount,
		financialCardRepo:       financialCard,
		tokenGen:                tokenGen,
	}
}
