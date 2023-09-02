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

type BankBranchHandler struct {
	logger            *zap.SugaredLogger
	bankBranchService protocol.BankBranch
}

func NewBranchHandler(logger *zap.SugaredLogger, bankBranchService protocol.BankBranch) *BankBranchHandler {
	return &BankBranchHandler{logger: logger, bankBranchService: bankBranchService}
}

func (h *BankBranchHandler) AddBranchHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req request.AddBranch

		if err := c.Bind(&req); err != nil {
			h.logger.Error("Failed to bind request",
				zap.Error(err),
				zap.String("handler", "AddBranchHandler"),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
		}

		resp, err := bankBranchService.AddBranch(ctx, req)
		if err != nil {
			h.logger.Error("Failed to add branch",
				zap.Error(err),
				zap.String("handler", "AddBranchHandler"),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Branch added successfully",
			Data:    resp,
		})
	}
}

func (h *BankBranchHandler) GetBranchByIDHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		branchID, err := extractIntParam(c, "id")
		if err != nil {

			h.logger.Error("Invalid branch ID",
				zap.Error(err),
				zap.String("handler", "GetBranchByIDHandler"),
				zap.String("parameter", "id"),
				zap.String("value", c.Param("id")),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid branch ID"))
		}

		h.logger.Info("Attempting to get branch by ID",
			zap.Int("branchID", branchID),
			zap.String("handler", "GetBranchByIDHandler"),
		)

		resp, err := bankBranchService.GetBranchByID(ctx, branchID)
		if err != nil {
			h.logger.Error("Failed to get branch by ID",
				zap.Error(err),
				zap.Int("branchID", branchID),
				zap.String("handler", "GetBranchByIDHandler"),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		h.logger.Info("Successfully retrieved branch by ID",
			zap.Int("branchID", branchID),
			zap.String("handler", "GetBranchByIDHandler"),
		)

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}

func extractIntParam(c echo.Context, paramName string) (int, error) {
	strValue := c.Param(paramName)
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func (h *BankBranchHandler) GetBranchByNameHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		branchName := c.Param("name")

		resp, err := bankBranchService.GetBranchByName(ctx, branchName)
		if err != nil {
			h.logger.Error("Failed to get branch by name",
				zap.Error(err),
				zap.String("handler", "GetBranchByNameHandler"),
				zap.String("branchName", branchName),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}

func (h *BankBranchHandler) GetBranchByCodeHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		branchCode := c.Param("code")

		resp, err := bankBranchService.GetBranchByCode(ctx, branchCode)
		if err != nil {
			h.logger.Error("Failed to get branch by code",
				zap.Error(err),
				zap.String("handler", "GetBranchByCodeHandler"),
				zap.String("branchCode", branchCode),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}

func (h *BankBranchHandler) UpdateBranchHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req request.UpdateBranch

		if err := c.Bind(&req); err != nil {
			h.logger.Error("Failed to bind request",
				zap.Error(err),
				zap.String("handler", "UpdateBranchHandler"),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid request"))
		}

		err := bankBranchService.UpdateBranch(ctx, req)
		if err != nil {
			h.logger.Error("Failed to update branch",
				zap.Error(err),
				zap.String("handler", "UpdateBranchHandler"),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Branch updated successfully",
			Data:    nil,
		})
	}
}

func (h *BankBranchHandler) DeleteBranchHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		branchIDStr := c.Param("id")
		branchID, err := strconv.Atoi(branchIDStr)
		if err != nil {
			h.logger.Error("Invalid branch ID",
				zap.Error(err),
				zap.String("handler", "DeleteBranchHandler"),
				zap.String("branchID", branchIDStr),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid branch ID"))
		}

		err = bankBranchService.DeleteBranch(ctx, branchID)
		if err != nil {
			h.logger.Error("Failed to delete branch",
				zap.Error(err),
				zap.String("handler", "DeleteBranchHandler"),
				zap.Int("branchID", branchID),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Branch deleted successfully",
			Data:    nil,
		})
	}
}

func (h *BankBranchHandler) ListAllBranchesHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		branches, err := bankBranchService.ListAllBranches(ctx)
		if err != nil {
			h.logger.Error("Failed to list all branches",
				zap.Error(err),
				zap.String("handler", "ListAllBranchesHandler"),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    branches,
		})
	}
}

func (h *BankBranchHandler) ListBranchesByStatusHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		statusStr := c.Param("status")
		statusInt, err := strconv.ParseUint(statusStr, 10, 32)
		if err != nil {
			h.logger.Error("Invalid status value",
				zap.Error(err),
				zap.String("handler", "ListBranchesByStatusHandler"),
				zap.String("status", statusStr),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid status value"))
		}

		status := enum.BankBranchStatus(statusInt)

		if status != enum.BankBranchActive && status != enum.BankBranchInactive {
			h.logger.Warn("Invalid enum value for status",
				zap.String("handler", "ListBranchesByStatusHandler"),
				zap.Uint64("status", statusInt),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid status value"))
		}

		branches, err := bankBranchService.ListBranchesByStatus(ctx, status)
		if err != nil {
			h.logger.Error("Failed to list branches by status",
				zap.Error(err),
				zap.String("handler", "ListBranchesByStatusHandler"),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    branches,
		})
	}
}

func (h *BankBranchHandler) ListBranchesByBankIDHandler(bankBranchService protocol.BankBranch) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		bankIDStr := c.Param("id")
		bankID, err := strconv.Atoi(bankIDStr)
		if err != nil {
			h.logger.Error("Invalid bank ID",
				zap.Error(err),
				zap.String("handler", "ListBranchesByBankIDHandler"),
				zap.String("bankID", bankIDStr),
			)
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid bank ID"))
		}

		branches, err := bankBranchService.ListBranchesByBankID(ctx, bankID)
		if err != nil {
			h.logger.Error("Failed to list branches by bank ID",
				zap.Error(err),
				zap.String("handler", "ListBranchesByBankIDHandler"),
			)
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    branches,
		})
	}
}
