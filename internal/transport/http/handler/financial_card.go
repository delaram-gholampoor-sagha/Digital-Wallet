package handler

import (
	"net/http"
	"strconv"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type FinancialCardHandler struct {
	logger               *zap.SugaredLogger
	financialCardService protocol.FinancialCard
}

func NewFinancialCardHandler(logger *zap.SugaredLogger, financialCardService protocol.FinancialCard) *FinancialCardHandler {
	return &FinancialCardHandler{logger: logger, financialCardService: financialCardService}
}

func (h *FinancialCardHandler) RegisterCardHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.RegisterFinancialCard
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
	}

	ID, err := h.financialCardService.RegisterCard(ctx, &req)
	if err != nil {
		h.logger.Error("Failed to register card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Card registered successfully",
		Data:    ID,
	})
}

func (h *FinancialCardHandler) UpdateCardHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UpdateFinancialCard
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
	}

	err := h.financialCardService.UpdateCard(ctx, &req)
	if err != nil {
		h.logger.Error("Failed to update card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Card updated successfully",
	})
}

func (h *FinancialCardHandler) DeleteCardHandler(c echo.Context) error {
	ctx := c.Request().Context()

	cardID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error("Invalid card ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid card ID"))
	}

	err = h.financialCardService.DeleteCard(ctx, cardID)
	if err != nil {
		h.logger.Error("Failed to delete card", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Card deleted successfully",
	})
}

func (h *FinancialCardHandler) GetCardByIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cardIDStr := c.Param("id")
	cardID, err := strconv.ParseInt(cardIDStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid card ID",
			zap.Error(err),
			zap.String("handler", "GetCardByIDHandler"),
		)
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid card ID"))
	}

	card, err := h.financialCardService.GetCardByID(ctx, cardID)
	if err != nil {
		h.logger.Error("Failed to get card by ID",
			zap.Error(err),
			zap.String("handler", "GetCardByIDHandler"),
		)
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Success",
		Data:    card,
	})
}

func (h *FinancialCardHandler) ListCardsByAccountIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	accountIDStr := c.Param("id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		h.logger.Error("Invalid account ID",
			zap.Error(err),
			zap.String("handler", "ListCardsByAccountIDHandler"),
		)
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	cards, err := h.financialCardService.ListCardsByAccountID(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to list cards by account ID",
			zap.Error(err),
			zap.String("handler", "ListCardsByAccountIDHandler"),
		)
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Success",
		Data:    cards,
	})
}

func (h *FinancialCardHandler) ListCardsByTypeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cardTypeStr := c.Param("type")
	cardType, err := strconv.ParseInt(cardTypeStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid card type",
			zap.Error(err),
			zap.String("handler", "ListCardsByTypeHandler"),
		)
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid card type"))
	}

	cards, err := h.financialCardService.ListCardsByType(ctx, enum.FinancialCardType(cardType))
	if err != nil {
		h.logger.Error("Failed to list cards by type",
			zap.Error(err),
			zap.String("handler", "ListCardsByTypeHandler"),
		)
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Success",
		Data:    cards,
	})
}
