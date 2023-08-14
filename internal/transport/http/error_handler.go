package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func errorHandler(logger *zap.SugaredLogger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {

	}
}
