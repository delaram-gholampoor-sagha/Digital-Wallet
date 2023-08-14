package jwt

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Claims(c echo.Context) *entity.JWTClaims {
	return c.Get("user").(*jwt.Token).Claims.(*entity.JWTClaims)
}
