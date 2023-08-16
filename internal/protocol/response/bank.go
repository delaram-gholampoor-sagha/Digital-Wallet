package response

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type GetBank struct {
	BankID    int64
	Name      string
	BankCode  string
	Status    enum.BankStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RegisterBank struct {
	BankID int64
}
