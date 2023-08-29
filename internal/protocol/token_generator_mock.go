package protocol

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockTokenGenerator struct {
	mock.Mock
}

func (m *MockTokenGenerator) GenerateToken(claims entity.JWTClaims, secret string) (string, error) {
	args := m.Called(claims, secret)
	return args.String(0), args.Error(1)
}
