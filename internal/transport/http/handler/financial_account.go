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

type FinancialAccountHandler struct {
	logger                  *zap.SugaredLogger
	FinancialAccountService protocol.FinancialAccount
}

func NewFinancialAccountHandler(logger *zap.SugaredLogger, FinancialAccountService protocol.FinancialAccount) *FinancialAccountHandler {
	return &FinancialAccountHandler{
		logger:                  logger,
		FinancialAccountService: FinancialAccountService,
	}
}

func (h *FinancialAccountHandler) CreateAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("CreateAccountHandler invoked")

	var req request.RegisterFinancialAccount
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request in CreateAccountHandler", zap.Error(err), zap.String("context", "binding"))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request payload"))
	}

	h.logger.Debug("CreateAccount request payload validated", zap.Any("requestData", req))

	resp, err := h.FinancialAccountService.CreateAccount(ctx, req)
	if err != nil {
		h.logger.Error("Failed to create account", zap.Error(err), zap.String("context", "service_call"))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Account successfully created", zap.Any("response", resp))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Account created successfully",
		Data:    resp,
	})
}

func (h *FinancialAccountHandler) GetAccountByIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("GetAccountByIDHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in GetAccountByIDHandler", zap.Error(err), zap.String("context", "validation"))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	h.logger.Debug("Valid accountID", zap.Int("accountID", accountID))

	resp, err := h.FinancialAccountService.GetAccountByID(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to get account by ID", zap.Error(err), zap.String("context", "service_call"))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Account successfully retrieved", zap.Int("accountID", accountID), zap.Any("response", resp))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Account retrieved successfully",
		Data:    resp,
	})
}

func (h *FinancialAccountHandler) UpdateAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("UpdateAccountHandler invoked")

	var req request.UpdateFinancialAccount
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to bind request in UpdateAccountHandler", zap.Error(err), zap.String("context", "binding"))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request payload"))
	}

	h.logger.Debug("UpdateAccount request payload validated", zap.Any("requestData", req))

	err := h.FinancialAccountService.UpdateAccount(ctx, req)
	if err != nil {
		h.logger.Error("Failed to update account", zap.Error(err), zap.String("context", "service_call"))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Account successfully updated", zap.Any("requestData", req))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Account updated successfully",
	})
}

func (h *FinancialAccountHandler) DeleteAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("DeleteAccountHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in DeleteAccountHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	if err := h.FinancialAccountService.DeleteAccount(ctx, accountID); err != nil {
		h.logger.Error("Failed to delete account", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully deleted account", zap.Int("accountID", accountID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Account deleted successfully",
	})
}

func (h *FinancialAccountHandler) ListAccountsByUserIDHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("ListAccountsByUserIDHandler invoked")

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		h.logger.Error("Invalid user ID in ListAccountsByUserIDHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid user ID"))
	}

	accounts, err := h.FinancialAccountService.ListAccountsByUserID(ctx, userID)
	if err != nil {
		h.logger.Error("Failed to list accounts by user ID", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully listed accounts by user ID", zap.Int("userID", userID), zap.Int("accountCount", len(accounts)))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Accounts retrieved successfully",
		Data:    accounts,
	})
}

func (h *FinancialAccountHandler) ListAccountsByStatusHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("ListAccountsByStatusHandler invoked")

	statusStr := c.QueryParam("status")

	var statusMap = map[string]enum.FinancialAccountStatus{
		"Verified":   enum.Verified,
		"Unverified": enum.Unverified,
	}

	var enumToStringMap = map[enum.FinancialAccountStatus]string{
		enum.Verified:  "Verified",
		enum.Unverified: "Unverified",
	}

	statusEnum, ok := statusMap[statusStr]

	if !ok {
		h.logger.Error("Invalid account status in ListAccountsByStatusHandler", zap.String("status", statusStr))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account status"))
	}

	accounts, err := h.FinancialAccountService.ListAccountsByStatus(ctx, statusEnum)
	if err != nil {
		h.logger.Error("Failed to list accounts by status", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully listed accounts by status", zap.String("status", enumToStringMap[statusEnum]), zap.Int("accountCount", len(accounts)))

	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Accounts retrieved successfully",
		Data:    accounts,
	})
}

func (h *FinancialAccountHandler) VerifyAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("VerifyAccountHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in VerifyAccountHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	if err := h.FinancialAccountService.VerifyAccount(ctx, accountID); err != nil {
		h.logger.Error("Failed to verify account", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully verified account", zap.Int("accountID", accountID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Account verified successfully",
	})
}

func (h *FinancialAccountHandler) GetAccountByShabaHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("GetAccountByShabaHandler invoked")

	shabaNumber := c.Param("shabaNumber")

	account, err := h.FinancialAccountService.GetAccountByShaba(ctx, shabaNumber)
	if err != nil {
		h.logger.Error("Failed to retrieve account by Shaba number", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved account by Shaba number", zap.String("shabaNumber", shabaNumber))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Account retrieved successfully",
		Data:    account,
	})
}



func (h *FinancialAccountHandler) ListAccountsByTypeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("ListAccountsByTypeHandler invoked")

	accountType := c.Param("type")

	accounts, err := h.FinancialAccountService.ListAccountsByType(ctx, accountType)
	if err != nil {
		h.logger.Error("Failed to list accounts by type", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully listed accounts by type", zap.String("type", accountType), zap.Int("accountCount", len(accounts)))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Accounts listed successfully",
		Data:    accounts,
	})
}

func (h *FinancialAccountHandler) GetAccountCurrencyHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("GetAccountCurrencyHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in GetAccountCurrencyHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	currency, err := h.FinancialAccountService.GetAccountCurrency(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to get currency for account", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved currency for account", zap.Int("accountID", accountID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Currency retrieved successfully",
		Data:    currency,
	})
}

func (h *FinancialAccountHandler) GetBranchForAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("GetBranchForAccountHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in GetBranchForAccountHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	branch, err := h.FinancialAccountService.GetBranchForAccount(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to get branch for account", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved branch for account", zap.Int("accountID", accountID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Branch retrieved successfully",
		Data:    branch,
	})
}

func (h *FinancialAccountHandler) GetBankForAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	h.logger.Info("GetBankForAccountHandler invoked")

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid account ID in GetBankForAccountHandler", zap.Error(err))
		return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid account ID"))
	}

	bank, err := h.FinancialAccountService.GetBankForAccount(ctx, accountID)
	if err != nil {
		h.logger.Error("Failed to get bank for account", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
	}

	h.logger.Info("Successfully retrieved bank for account", zap.Int("accountID", accountID))
	return c.JSON(http.StatusOK, protocol.Success{
		Message: "Bank retrieved successfully",
		Data:    bank,
	})
}
