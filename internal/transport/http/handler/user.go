package handler

import (
	"net/http"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/transport/http/jwt"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"github.com/labstack/echo/v4"
)

func SignUpHandler(userService protocol.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req request.SignUp

		if err := c.Bind(&req); err != nil {
			return derror.NewBadRequestError(message.InvalidRequest)
		}

		tokens, err := userService.SignUp(ctx, req)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    tokens,
		})
	}
}

func SignInHandler(userService protocol.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req request.SignIn

		if err := c.Bind(&req); err != nil {
			return derror.NewBadRequestError(message.InvalidRequest)
		}

		tokens, err := userService.SignIn(ctx, req)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    tokens,
		})
	}
}

func RefreshTokenHandler(userService protocol.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID := jwt.Claims(c).UserID

		tokens, err := userService.RefreshToken(ctx, userID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    tokens,
		})
	}
}

func GetProfileHandler(userService protocol.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID := jwt.Claims(c).UserID

		profile, err := userService.GetProfile(ctx, userID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    profile,
		})
	}
}

func EditProfileHandler(userService protocol.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var req request.EditProfile
		if err := c.Bind(&req); err != nil {
			return derror.NewBadRequestError(message.InvalidRequest)
		}

		req.UserID = jwt.Claims(c).UserID

		if err := userService.EditProfile(ctx, req); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    nil,
		})
	}
}
