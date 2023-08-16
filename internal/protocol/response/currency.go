package response

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type AddCurrency struct {
	CurrencyID int
}

type UpdateCurrency struct {
	CurrencyID int
}

type CurrencyComparison struct {
	FirstCurrency  entity.Currency
	SecondCurrency entity.Currency
	IsStronger     bool
	ComparisonRate float64
}

type CurrencyTrends struct {
	Code   enum.CurrencyCode
	Trends []struct {
		Date         time.Time
		ExchangeRate float64
	}
}

type SearchCurrenciesResult struct {
	Currencies []entity.Currency
}

type ConvertAmountResult struct {
	ConvertedAmount float64
}

// This could be an object showing the latest exchange rate of the two currencies and how they fare against a base currency, or any other information relevant to comparing two currencies.
type CompareCurrenciesResult struct {
	FirstCurrencyRate  float64
	SecondCurrencyRate float64
	ComparisonValue    float64 // e.g., the difference or ratio
}

type GetStrongestCurrencyResult struct {
	Currency entity.Currency
	Rate     float64
}

type GetWeakestCurrencyResult struct {
	Currency entity.Currency
	Rate     float64
}

type GetCountriesUsingCurrencyResult struct {
	Countries []string
}
