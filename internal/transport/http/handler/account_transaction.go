package handler

import (
	"net/http"
	"strconv"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AccountTransactionHandler struct {
	logger                    *zap.SugaredLogger
	accountTransactionService protocol.AccountTransaction
}

func NewAccountTransactionHandler(logger *zap.SugaredLogger, accountTransactionService protocol.AccountTransaction) *AccountTransactionHandler {
	return &AccountTransactionHandler{logger: logger, accountTransactionService: accountTransactionService}
}

func (h *AccountTransactionHandler) RegisterTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.RegisterTransactionRequest

	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request",
			zap.Error(err),
			zap.String("handler", "RegisterTransactionHandler"),
		)
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
	}

	resp, err := h.accountTransactionService.RegisterTransaction(ctx, &req)
	if err != nil {
		h.logger.Error("Failed to register transaction",
			zap.Error(err),
			zap.String("handler", "RegisterTransactionHandler"),
		)
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Transaction registered successfully",
		Data:    resp,
	})
}

func (h *AccountTransactionHandler) DeleteTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()
	transactionIDStr := c.Param("transactionID")
	transactionID, err := strconv.ParseInt(transactionIDStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid transaction ID",
			zap.Error(err),
			zap.String("handler", "DeleteTransactionHandler"),
			zap.String("transactionID", transactionIDStr),
		)
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid transaction ID"))
	}

	err = h.accountTransactionService.DeleteTransaction(ctx, transactionID)
	if err != nil {
		h.logger.Error("Failed to delete transaction",
			zap.Error(err),
			zap.String("handler", "DeleteTransactionHandler"),
			zap.Int64("transactionID", transactionID),
		)
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Transaction deleted successfully",
		Data:    nil,
	})
}
func (h *AccountTransactionHandler) GetTransactionByIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseInt(transactionIDStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid transaction ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid transaction ID"))
	}

	resp, err := h.accountTransactionService.GetTransactionByID(ctx, transactionID)
	if err != nil {
		h.logger.Error("Failed to get transaction by ID", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Success",
		Data:    resp,
	})
}

func (h *AccountTransactionHandler) ListTransactionsByAccountIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	accountIDStr := c.Param("id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		h.logger.Error("Invalid account ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	resp, err := h.accountTransactionService.ListTransactionsByAccountID(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to list transactions by account ID", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Success",
		Data:    resp,
	})
}

func (h *AccountTransactionHandler) ListTransactionsByGroupIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	groupIDStr := c.Param("id")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		h.logger.Error("Invalid group ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid group ID"))
	}

	resp, err := h.accountTransactionService.ListTransactionsByGroupID(ctx, groupID)
	if err != nil {
		h.logger.Error("Failed to list transactions by group ID", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Success",
		Data:    resp,
	})
}

func (h *AccountTransactionHandler) CancelTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseInt(transactionIDStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid transaction ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid transaction ID"))
	}

	err = h.accountTransactionService.CancelTransaction(ctx, transactionID)
	if err != nil {
		h.logger.Error("Failed to cancel transaction", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Transaction canceled successfully",
		Data:    nil,
	})
}

func (h *AccountTransactionHandler) TransferHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.TransferRequest

	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
	}

	resp, err := h.accountTransactionService.Transfer(ctx, req)
	if err != nil {
		h.logger.Error("Failed to initiate transfer", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Transfer initiated successfully",
		Data:    resp,
	})
}

func (h *AccountTransactionHandler) GetAccountTransactionHistoryHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("GetAccountTransactionHistoryHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in GetAccountTransactionHistoryHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	transactions, err := h.accountTransactionService.GetAccountTransactionHistory(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to retrieve account transaction history", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved account transaction history", zap.Int("accountID", accountID), zap.Int("transactionCount", len(transactions)))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Transaction history retrieved successfully",
		Data:    transactions,
	})
}
