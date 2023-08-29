package protocol

import "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"

// TokenGenerator provides an interface to generate JWT tokens.
type TokenGenerator interface {
	GenerateToken(claims entity.JWTClaims, secret string) (string, error)
}
