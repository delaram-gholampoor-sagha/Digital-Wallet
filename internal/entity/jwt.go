package entity

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID int  `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func (c *JWTClaims) Valid() error {
	now := time.Now()
	if c.ExpiresAt == nil {
		return fmt.Errorf("expiration time is missing")
	}
	if now.After(c.ExpiresAt.Time) {
		return fmt.Errorf("token has expired")
	}
	// Add other validation checks if required
	return nil
}
