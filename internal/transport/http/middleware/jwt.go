package middleware

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWT(secret string) echo.MiddlewareFunc {
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JWTClaims)
		},
		SigningKey: []byte(secret),
	}

	return echojwt.WithConfig(jwtConfig)
}
