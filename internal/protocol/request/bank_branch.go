package request

import "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"

type AddBranch struct {
	BankID      int
	BranchName  string
	BranchCode  string
	Address     string
	City        string
	Province    string
	PostalCode  string
	PhoneNumber string
	Status      enum.BankBranchStatus
}

type UpdateBranch struct {
	BranchID    int
	BankID      int // if you want to allow moving branches between banks
	BranchName  string
	BranchCode  string
	Address     string
	City        string
	Province    string
	PostalCode  string
	PhoneNumber string
	Status      enum.BankBranchStatus
}
