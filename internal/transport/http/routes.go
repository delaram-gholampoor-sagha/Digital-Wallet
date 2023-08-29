package http

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/handler"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/middleware"
)

func (s *Server) register(secret string, userService protocol.User) {
	auth := s.echo.Group("/auth")
	auth.POST("/sign-up", handler.SignUpHandler(userService))
	auth.POST("/sign-in", handler.SignInHandler(userService))
	auth.POST("/refresh", handler.RefreshTokenHandler(userService), middleware.JWT(secret))

	user := s.echo.Group("/account", middleware.JWT(secret))
	user.GET("profile", handler.GetProfileHandler(userService))
	user.PUT("profile", handler.EditProfileHandler(userService))
}
