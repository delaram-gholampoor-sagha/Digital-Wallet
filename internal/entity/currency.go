package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type Currency struct {
	CurrencyID   int
	CurrencyCode enum.CurrencyCode
	CurrencyName enum.CurrencyName
	Symbol       enum.CurrencySymbol
	ExchangeRate *float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
