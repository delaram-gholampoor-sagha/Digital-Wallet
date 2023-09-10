package request

import (
	"errors"
	"math"
)

type RegisterTransactionRequest struct {
	TransactionGroupID int
	FinancialAccountID int
	Amount             float64
	Description        *string
}

func (req *RegisterTransactionRequest) Validate() error {
	// Check if TransactionGroupID is provided and valid
	if req.TransactionGroupID <= 0 {
		return errors.New("Invalid Transaction Group ID")
	}

	// Check if FinancialAccountID is provided and valid
	if req.FinancialAccountID <= 0 {
		return errors.New("Invalid Financial Account ID")
	}

	// Check if Amount is within range.
	// In this case, we use 15 digits with 2 decimal places as in your DECIMAL(15, 2) field.
	// Make sure to adapt as per your business logic.
	if req.Amount < -9999999999999.99 || req.Amount > 9999999999999.99 {
		return errors.New("Amount out of range")
	}

	// Optionally, you can check for a zero amount if your business logic does not allow it.
	if req.Amount == 0 {
		return errors.New("Amount cannot be zero")
	}

	// Check if Description is provided and its length.
	// Make it optional if your business logic allows it.
	if req.Description != nil {
		descLen := len(*req.Description)
		if descLen == 0 || descLen > 255 { // Assuming 255 as a reasonable maximum length for a text description.
			return errors.New("Description must be between 1 and 255 characters")
		}
	}

	return nil
}

type TransferRequest struct {
	SenderAccountID   int
	ReceiverAccountID int
	Amount            float64
	Description       string
}

func (req *TransferRequest) Validate() error {
	if req.SenderAccountID <= 0 {
		return errors.New("invalid sender account ID")
	}

	if req.ReceiverAccountID <= 0 {
		return errors.New("invalid receiver account ID")
	}

	if req.Amount <= 0 {
		return errors.New("invalid transfer amount")
	}

	// Check if the amount has more than 2 decimal places
	if math.Round(req.Amount*100) != req.Amount*100 {
		return errors.New("amount can have up to 2 decimal places")
	}

	// Assuming a max description length of 255
	if len(req.Description) > 255 {
		return errors.New("description too long")
	}

	return nil
}
