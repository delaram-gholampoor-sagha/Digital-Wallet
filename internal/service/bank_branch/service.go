package bankbranch

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"go.uber.org/zap"
)

type Service struct {
	cfg            config.JWT
	logger         *zap.SugaredLogger
	bank           protocol.Bank
	bankBranchRepo protocol.BankBranchRepository
	tokenGen       protocol.TokenGenerator
}

func New(cfg config.JWT, logger *zap.SugaredLogger, bankBranch protocol.BankBranchRepository, tokenGen protocol.TokenGenerator, bank protocol.Bank) *Service {
	return &Service{
		cfg:            cfg,
		logger:         logger,
		bank:           bank,
		bankBranchRepo: bankBranch,
		tokenGen:       tokenGen,
	}
}
