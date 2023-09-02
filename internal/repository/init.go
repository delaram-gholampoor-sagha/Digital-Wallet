package repository

import "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"

func NewUser(database protocol.Database) *User {
	return &User{cli: database.DB()}
}

func NewBank(database protocol.Database) *Bank {
	return &Bank{cli: database.DB()}
}

func NewBankBranch(database protocol.Database) *BankBranch {
	return &BankBranch{cli: database.DB()}
}

