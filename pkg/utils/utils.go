package utils

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Real implementations of Hasher and TokenGenerator
type BcryptHasher struct{}

func (h BcryptHasher) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (h BcryptHasher) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

// token generator
type JWTTokenGenerator struct{}

func (t JWTTokenGenerator) GenerateToken(claims entity.JWTClaims, secret string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}
