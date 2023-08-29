package user

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"go.uber.org/zap"
)

type Service struct {
	cfg      config.JWT
	logger   *zap.SugaredLogger
	userRepo protocol.UserRepository
	hasher   protocol.Hasher
	tokenGen protocol.TokenGenerator
}

func New(cfg config.JWT, logger *zap.SugaredLogger, userRepo protocol.UserRepository, hasher protocol.Hasher, tokenGen protocol.TokenGenerator) *Service {
	return &Service{
		cfg:      cfg,
		logger:   logger,
		userRepo: userRepo,
		hasher:   hasher,
		tokenGen: tokenGen,
	}
}
