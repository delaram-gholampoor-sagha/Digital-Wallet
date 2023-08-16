package request

import "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"

type UpdateBank struct {
	BankID   int64
	Name     string
	BankCode string
	Status   enum.BankStatus
}


type RegisterBank struct {
	Name     string
	BankCode string
	Status   enum.BankStatus
}