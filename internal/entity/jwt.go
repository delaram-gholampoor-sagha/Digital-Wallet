package entity

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID int  `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}
