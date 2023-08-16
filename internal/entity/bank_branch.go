package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type BankBranch struct {
	BranchID    int
	BankID      int
	BranchName  string
	BranchCode  string
	Address     string
	City        string
	Province    string
	PostalCode  string
	PhoneNumber string
	Status      enum.BankBranchStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
