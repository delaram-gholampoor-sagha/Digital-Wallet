package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterBankHandler(t *testing.T) {
	e := echo.New()

	t.Run("Successfully create a new bank", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		mockResponse := response.RegisterBank{BankID: 1}
		mockBankService.On("RegisterBank", mock.Anything, mock.Anything).Return(mockResponse, nil)

		reqBody := `{"Name": "Bank A", "BankCode": "A", "Status": "Active"}`
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(reqBody)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := RegisterBankHandler(mockBankService)(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response response.RegisterBank
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, int64(1), response.BankID)
	})

	t.Run("Fail with bad request", func(t *testing.T) {
		// My mock setup and test logic for a bad request case
		// For example, return a bad request if the "Name" field is empty.
	})

	t.Run("Fail with internal server error", func(t *testing.T) {
		// My mock setup and test logic for an internal server error case
		// For example, mock an error from the service method and verify the handler returns a 500 status.
	})
}

func TestGetBankByIDHandler(t *testing.T) {

	e := echo.New()

	t.Run("Successfully get bank by ID", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := GetBankByIDHandler(mockBankService)

		bankID := 1
		expectedResponse := response.GetBank{BankID: int64(bankID), Name: "Test Bank", BankCode: "TST"}

		mockBankService.On("GetBankByID", mock.Anything, bankID).Return(expectedResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/banks/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var res response.GetBank
			if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
				t.Errorf("Failed to unmarshal response: %s", err)
			}

			assert.Equal(t, expectedResponse, res)
		}

		mockBankService.AssertExpectations(t)
	})

	t.Run("Fail with not found", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := GetBankByIDHandler(mockBankService)

		bankID := 1

		mockBankService.On("GetBankByID", mock.Anything, bankID).Return(response.GetBank{}, errors.New("not found"))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/banks/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}

		mockBankService.AssertExpectations(t)
	})

	t.Run("Fail with internal server error", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := GetBankByIDHandler(mockBankService)

		bankID := 1

		mockBankService.On("GetBankByID", mock.Anything, bankID).Return(response.GetBank{}, errors.New("internal error"))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/banks/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}

		mockBankService.AssertExpectations(t)
	})
}

func TestGetBankByCodeHandler(t *testing.T) {
	
	e := echo.New()

	t.Run("Successfully get bank by code", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := GetBankByCodeHandler(mockBankService)

		bankCode := "TST"
		expectedResponse := response.GetBank{BankID: 1, Name: "Test Bank", BankCode: "TST"}

		mockBankService.On("GetBankByCode", mock.Anything, bankCode).Return(expectedResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/banks/code/:code")
		c.SetParamNames("code")
		c.SetParamValues("TST")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var res response.GetBank
			if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
				t.Errorf("Failed to unmarshal response: %s", err)
			}

			assert.Equal(t, expectedResponse, res)
		}

		mockBankService.AssertExpectations(t)
	})

}

func TestGetBankByNameHandler(t *testing.T) {
	// Create an Echo instance for testing
	e := echo.New()

	t.Run("Successfully get bank by name", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := GetBankByNameHandler(mockBankService)

		bankName := "Test Bank"
		expectedResponse := response.GetBank{BankID: 1, Name: "Test Bank", BankCode: "TST"}

		mockBankService.On("GetBankByName", mock.Anything, bankName).Return(expectedResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/banks/name/:name")
		c.SetParamNames("name")
		c.SetParamValues("Test Bank")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var res response.GetBank
			if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
				t.Errorf("Failed to unmarshal response: %s", err)
			}

			assert.Equal(t, expectedResponse, res)
		}

		mockBankService.AssertExpectations(t)
	})

}

func TestUpdateBankDetailsHandler(t *testing.T) {
	e := echo.New()

	t.Run("Successfully update bank details", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := UpdateBankDetailsHandler(mockBankService)

		reqBody := request.UpdateBank{
			BankID:   1,
			Name:     "Updated Bank",
			BankCode: "UPD",
			Status:   enum.BankActive,
		}
		mockBankService.On("UpdateBankDetails", mock.Anything, reqBody).Return(nil)

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBufferString(`{"BankID": 1, "Name": "Updated Bank", "BankCode": "UPD", "Status": "Active"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}

		mockBankService.AssertExpectations(t)
	})

}

func TestListAllBanksHandler(t *testing.T) {
	e := echo.New()

	t.Run("Successfully list all banks", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := ListAllBanksHandler(mockBankService)

		expectedResponse := []*entity.Bank{
			{BankID: 1, Name: "Test Bank", BankCode: "TST"},
			{BankID: 2, Name: "Another Bank", BankCode: "ANB"},
		}

		mockBankService.On("ListAllBanks", mock.Anything).Return(expectedResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var res []*entity.Bank
			if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
				t.Errorf("Failed to unmarshal response: %s", err)
			}

			assert.Equal(t, expectedResponse, res)
		}

		mockBankService.AssertExpectations(t)
	})

	
}

func TestListBanksByStatusHandler(t *testing.T) {
	e := echo.New()

	t.Run("Successfully list banks by status", func(t *testing.T) {
		mockBankService := new(protocol.MockBankService)
		h := ListBanksByStatusHandler(mockBankService)

		status := enum.Active
		expectedResponse := []*entity.Bank{
			{BankID: 1, Name: "Active Bank 1", BankCode: "AB1", Status: enum.BankActive},
			{BankID: 2, Name: "Active Bank 2", BankCode: "AB2", Status: enum.BankActive},
		}

		mockBankService.On("ListBanksByStatus", mock.Anything, status).Return(expectedResponse, nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/banks/status/:status")
		c.SetParamNames("status")
		c.SetParamValues("Active")

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var res []*entity.Bank
			if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
				t.Errorf("Failed to unmarshal response: %s", err)
			}

			assert.Equal(t, expectedResponse, res)
		}

		mockBankService.AssertExpectations(t)
	})

}
