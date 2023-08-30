package bank

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"go.uber.org/zap"
)

type Service struct {
	cfg      config.JWT
	logger   *zap.SugaredLogger
	bankRepo protocol.BankRepository
	tokenGen protocol.TokenGenerator
}

func New(cfg config.JWT, logger *zap.SugaredLogger, bankRepo protocol.BankRepository, tokenGen protocol.TokenGenerator) *Service {
	return &Service{
		cfg:      cfg,
		logger:   logger,
		bankRepo: bankRepo,
		tokenGen: tokenGen,
	}
}
