package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type Bank struct {
	BankID    int64
	Name      string
	BankCode  string
	Status    enum.BankStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
