package http

import (
	"errors"
	"net/http"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/translation"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func errorHandler(logger *zap.SugaredLogger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {

		if c.Response().Committed {
			return
		}

		language, ok := c.Get(translation.Locale).(translation.Language)
		if !ok {
			language = translation.Farsi
		}
		// TODO use for translation
		_ = language

		derr, ok := err.(*derror.Error)
		if !ok {
			switch {
			case errors.As(err, &jwt.ValidationError{}):
				derr = &derror.Error{
					Message: message.InvalidToken,
					Code:    http.StatusUnauthorized,
				}
			//case errors.Is(err, echojwt.ErrJWTInvalid):
			//	derr = &derror.Error{
			//		Message: message.InvalidToken,
			//		Code:    http.StatusUnauthorized,
			//	}
			default:
				logger.Errorw("transport.http.errorHandler: err not derror", "error", err)
				derr = &derror.Error{
					Message: message.InternalSystemError,
					Code:    http.StatusInternalServerError,
				}
			}
		}

		if c.Request().Method == http.MethodHead {
			err = c.NoContent(derr.Code)
		} else {
			err = c.JSON(derr.Code, protocol.Error{
				Message: derr.Message,
			})
		}

		if err != nil {
			logger.Errorw("transport.http.errorHandler", "error", err.Error())
		}
	}
}
