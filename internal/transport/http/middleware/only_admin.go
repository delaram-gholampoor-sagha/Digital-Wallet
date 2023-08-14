package middleware

import (
	"net/http"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/jwt"
	"github.com/labstack/echo/v4"
)

func OnlyAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !jwt.Claims(c).Admin {
				return c.NoContent(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
