package handler

import (
	"net/http"
	"strconv"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"github.com/labstack/echo/v4"
)

func RegisterBankHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req request.RegisterBank

		if err := c.Bind(&req); err != nil {
			return derror.NewBadRequestError(message.InvalidRequest)
		}

		resp, err := bankService.RegisterBank(ctx, req)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    resp,
		})
	}
}

func GetBankByIDHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		bankIDStr := c.Param("bankID")
		bankID, err := strconv.Atoi(bankIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid bank ID"))
		}

		resp, err := bankService.GetBankByID(ctx, bankID)
		if err != nil {
			if derror.IsNotFound(err) {
				return c.JSON(http.StatusNotFound, derror.NewNotFoundError("Bank not found"))
			}
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}

func GetBankByCodeHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		bankCode := c.Param("bankCode")
		resp, err := bankService.GetBankByCode(ctx, bankCode)
		if err != nil {
			if derror.IsNotFound(err) {
				return c.JSON(http.StatusNotFound, derror.NewNotFoundError("Bank not found"))
			}
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}

func GetBankByNameHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		bankName := c.Param("bankName")
		resp, err := bankService.GetBankByName(ctx, bankName)
		if err != nil {
			if derror.IsNotFound(err) {
				return c.JSON(http.StatusNotFound, derror.NewNotFoundError("Bank not found"))
			}
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}

func UpdateBankDetailsHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		var req request.UpdateBank

		if err := c.Bind(&req); err != nil {
			return derror.NewBadRequestError(message.InvalidRequest)
		}

		err := bankService.UpdateBankDetails(ctx, req)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "success",
			Data:    nil,
		})
	}
}

func ListAllBanksHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		banks, err := bankService.ListAllBanks(ctx)
		if err != nil {

			if derror.IsInternalError(err) {
				return c.JSON(http.StatusInternalServerError, protocol.Error{
					Message: "Internal System Error",
				})
			}

			return c.JSON(http.StatusInternalServerError, protocol.Error{
				Message: "An unexpected error occurred",
			})
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    banks,
		})
	}
}

func ListBanksByStatusHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		statusStr := c.QueryParam("status")
		statusInt, err := strconv.ParseUint(statusStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid status"))
		}

		status := uint(statusInt)
		if status != uint(enum.BankInactive) && status != uint(enum.BankActive) {
			return c.JSON(http.StatusBadRequest, derror.NewBadRequestError("Invalid status"))
		}

		resp, err := bankService.ListBanksByStatus(ctx, enum.BankStatus(status))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, derror.NewInternalSystemError())
		}

		return c.JSON(http.StatusOK, protocol.Success{
			Message: "Success",
			Data:    resp,
		})
	}
}
