package handler

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/labstack/echo/v4"
)


func RegisterBankHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}


func GetBankByIDHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}


func GetBankByCodeHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}

func GetBankByNameHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}


func UpdateBankDetailsHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}


func ListAllBanksHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}

func ListBanksByStatusHandler(bankService protocol.Bank) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Implement
		return nil
	}
}
