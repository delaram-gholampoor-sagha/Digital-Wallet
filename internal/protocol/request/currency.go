package request

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type AddCurrency struct {
	CurrencyCode enum.CurrencyCode
	CurrencyName enum.CurrencyName
	Symbol       enum.CurrencySymbol
	ExchangeRate *float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UpdateCurrency struct {
	CurrencyID   int
	CurrencyCode enum.CurrencyCode
	CurrencyName enum.CurrencyName
	Symbol       enum.CurrencySymbol
	ExchangeRate *float64
	UpdatedAt    time.Time
}

type BulkUpdateExchangeRates struct {
	Rates map[enum.CurrencyCode]float64
}

type SearchCurrencies struct {
	Query string
}

type ConvertAmount struct {
	FromCode enum.CurrencyCode
	ToCode   enum.CurrencyCode
	Amount   float64
}

type CompareCurrencies struct {
	FirstCode  enum.CurrencyCode
	SecondCode enum.CurrencyCode
}

type GetCurrencyTrends struct {
	Code     enum.CurrencyCode
	Duration time.Duration
}
