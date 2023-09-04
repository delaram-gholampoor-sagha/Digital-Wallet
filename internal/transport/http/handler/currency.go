package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CurrencyHandler struct {
	logger          *zap.SugaredLogger
	currencyService protocol.Currency
}

func NewCurrencyHandler(logger *zap.SugaredLogger, currencyService protocol.Currency) *CurrencyHandler {
	return &CurrencyHandler{
		logger:          logger,
		currencyService: currencyService,
	}
}

func (h *CurrencyHandler) AddCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.AddCurrency
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request in AddCurrencyHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request format"))
	}

	res, err := h.currencyService.AddCurrency(ctx, req)
	if err != nil {
		h.logger.Error("Failed to execute AddCurrency service method", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully added new currency", zap.Any("response", res))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency added successfully",
		Data:    res,
	})
}

func (h *CurrencyHandler) UpdateCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UpdateCurrency
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request in UpdateCurrencyHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request format"))
	}

	if err := h.currencyService.UpdateCurrency(ctx, req); err != nil {
		h.logger.Error("Failed to execute UpdateCurrency service method", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully updated currency")
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency updated successfully",
	})
}

func (h *CurrencyHandler) DeleteCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	currencyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid currency ID in DeleteCurrencyHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid currency ID"))
	}

	if err := h.currencyService.DeleteCurrency(ctx, currencyID); err != nil {
		h.logger.Error("Failed to execute DeleteCurrency service method", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully deleted currency", zap.Int("currencyID", currencyID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency deleted successfully",
	})
}

func (h *CurrencyHandler) GetCurrencyByIDHandler(c echo.Context) error {
	ctx := c.Request().Context()

	currencyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid currency ID in GetCurrencyByIDHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid currency ID"))
	}

	currency, err := h.currencyService.GetCurrency(ctx, currencyID)
	if err != nil {
		h.logger.Error("Failed to retrieve currency in GetCurrencyByIDHandler", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved currency by ID", zap.Int("currencyID", currencyID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency retrieved successfully",
		Data:    currency,
	})
}

func (h *CurrencyHandler) GetCurrencyByNameHandler(c echo.Context) error {
	ctx := c.Request().Context()

	currencyName := c.Param("name")

	currency, err := h.currencyService.GetCurrencyByName(ctx, enum.CurrencyName(currencyName))
	if err != nil {
		h.logger.Error("Failed to retrieve currency by name in GetCurrencyByNameHandler", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved currency by name", zap.String("currencyName", currencyName))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency retrieved successfully",
		Data:    currency,
	})
}

func (h *CurrencyHandler) ListCurrenciesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	currencies, err := h.currencyService.ListCurrencies(ctx)
	if err != nil {
		h.logger.Error("Failed to list currencies in ListCurrenciesHandler", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully listed all currencies")
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currencies retrieved successfully",
		Data:    currencies,
	})
}

func (h *CurrencyHandler) GetExchangeRateHandler(c echo.Context) error {
	ctx := c.Request().Context()

	fromCode := c.QueryParam("from")
	toCode := c.QueryParam("to")

	rate, err := h.currencyService.GetExchangeRate(ctx, enum.CurrencyCode(fromCode), enum.CurrencyCode(toCode))
	if err != nil {
		h.logger.Error("Failed to retrieve exchange rate in GetExchangeRateHandler", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved exchange rate", zap.String("from", fromCode), zap.String("to", toCode))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Exchange rate retrieved successfully",
		Data:    rate,
	})
}

func (h *CurrencyHandler) BulkUpdateExchangeRatesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req map[enum.CurrencyCode]float64
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request in BulkUpdateExchangeRatesHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
	}

	err := h.currencyService.BulkUpdateExchangeRates(ctx, req)
	if err != nil {
		h.logger.Error("Failed to bulk update exchange rates", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully updated exchange rates in bulk")
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Exchange rates updated successfully",
	})
}

func (h *CurrencyHandler) SearchCurrenciesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	query := c.QueryParam("query")

	currencies, err := h.currencyService.SearchCurrencies(ctx, query)
	if err != nil {
		h.logger.Error("Failed to search currencies in SearchCurrenciesHandler", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully searched currencies", zap.String("query", query))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currencies found successfully",
		Data:    currencies,
	})
}

func (h *CurrencyHandler) ConvertAmountHandler(c echo.Context) error {
	ctx := c.Request().Context()

	fromCode := c.QueryParam("from")
	toCode := c.QueryParam("to")
	amountStr := c.QueryParam("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		h.logger.Error("Failed to parse amount", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid amount"))
	}

	convertedAmount, err := h.currencyService.ConvertAmount(ctx, enum.CurrencyCode(fromCode), enum.CurrencyCode(toCode), amount)
	if err != nil {
		h.logger.Error("Failed to convert amount", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully converted amount", zap.String("from", fromCode), zap.String("to", toCode), zap.Float64("amount", amount))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Amount converted successfully",
		Data:    convertedAmount,
	})
}

func (h *CurrencyHandler) CompareCurrenciesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	firstCode := c.QueryParam("first")
	secondCode := c.QueryParam("second")

	comparison, err := h.currencyService.CompareCurrencies(ctx, enum.CurrencyCode(firstCode), enum.CurrencyCode(secondCode))
	if err != nil {
		h.logger.Error("Failed to compare currencies", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully compared currencies", zap.String("first", firstCode), zap.String("second", secondCode))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currencies compared successfully",
		Data:    comparison,
	})
}

func (h *CurrencyHandler) GetCurrencyTrendsHandler(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.QueryParam("code")
	durationStr := c.QueryParam("duration")
	duration, err := time.ParseDuration(durationStr)

	if err != nil {
		h.logger.Error("Failed to parse duration", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid duration"))
	}

	trends, err := h.currencyService.GetCurrencyTrends(ctx, enum.CurrencyCode(code), duration)
	if err != nil {
		h.logger.Error("Failed to get currency trends", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved currency trends", zap.String("code", code), zap.Duration("duration", duration))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency trends retrieved successfully",
		Data:    trends,
	})
}

func (h *CurrencyHandler) GetStrongestCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	strongestCurrency, err := h.currencyService.GetStrongestCurrency(ctx)
	if err != nil {
		h.logger.Error("Failed to retrieve strongest currency", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved strongest currency")
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Strongest currency retrieved successfully",
		Data:    strongestCurrency,
	})
}

func (h *CurrencyHandler) GetWeakestCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	weakestCurrency, err := h.currencyService.GetWeakestCurrency(ctx)
	if err != nil {
		h.logger.Error("Failed to retrieve weakest currency", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved weakest currency")
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Weakest currency retrieved successfully",
		Data:    weakestCurrency,
	})
}

func (h *CurrencyHandler) NotifyUsersOnExchangeRateChangeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var threshold float64
	if err := c.Bind(&threshold); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
	}

	if err := h.currencyService.NotifyUsersOnExchangeRateChange(ctx, threshold); err != nil {
		h.logger.Error("Failed to notify users", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully notified users on exchange rate change", zap.Float64("threshold", threshold))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Users notified successfully",
	})
}

func (h *CurrencyHandler) GetCountriesUsingCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.QueryParam("code")
	countries, err := h.currencyService.GetCountriesUsingCurrency(ctx, enum.CurrencyCode(code))
	if err != nil {
		h.logger.Error("Failed to retrieve countries using currency", zap.String("code", code), zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved countries using currency", zap.String("code", code))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Countries using currency retrieved successfully",
		Data:    countries,
	})
}
