package protocol

import (
	"context"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type CurrencyService interface {
	AddCurrency(ctx context.Context, req request.AddCurrency) (response.AddCurrency, error)
	UpdateCurrency(ctx context.Context, req request.UpdateCurrency) (response.UpdateCurrency, error)
	DeleteCurrency(ctx context.Context, currencyID int) error
	GetCurrency(ctx context.Context, currencyID int) (entity.Currency, error)
	GetCurrencyByName(ctx context.Context, currencyName enum.CurrencyName) (entity.Currency, error)
	ListCurrencies(ctx context.Context) ([]entity.Currency, error)
	GetExchangeRate(ctx context.Context, fromCode enum.CurrencyCode, toCode enum.CurrencyCode) (float64, error)

	// Update exchange rates for multiple currencies at once.
	BulkUpdateExchangeRates(ctx context.Context, rates map[enum.CurrencyCode]float64) error

	// Search for currencies based on partial matches (e.g., name, code, or symbol).
	SearchCurrencies(ctx context.Context, query string) ([]entity.Currency, error)

	ConvertAmount(ctx context.Context, fromCode enum.CurrencyCode, toCode enum.CurrencyCode, amount float64) (float64, error)

	// Compare two currencies – could be useful for traders or for users trying to decide which currency to transact in.
	CompareCurrencies(ctx context.Context, firstCode enum.CurrencyCode, secondCode enum.CurrencyCode) (response.CurrencyComparison, error)

	// Get the trend of a currency over a specified duration – could be daily, weekly, monthly changes, etc.
	GetCurrencyTrends(ctx context.Context, code enum.CurrencyCode, duration time.Duration) (response.CurrencyTrends, error)

	//  Retrieve the strongest and weakest currency based on the latest exchange rates.
	GetStrongestCurrency(ctx context.Context) (entity.Currency, error)
	GetWeakestCurrency(ctx context.Context) (entity.Currency, error)

	// Notify registered users if a particular currency's exchange rate changes beyond a certain threshold.
	// Useful for traders or businesses to keep an eye on currency fluctuations.
	NotifyUsersOnExchangeRateChange(ctx context.Context, threshold float64) error

	// to retrieve a list of countries that use a specific currency.
	// The return type would typically be a list of country names or country codes.
	// This is particularly useful for currencies that are shared among multiple countries, like the Euro (EUR).
	GetCountriesUsingCurrency(ctx context.Context, code enum.CurrencyCode) ([]string, error)
}

type CurrencyRepository interface {
	Insert(ctx context.Context, c entity.Currency) error
	Update(ctx context.Context, c entity.Currency) error
	Delete(ctx context.Context, currencyID int) error
	Get(ctx context.Context, currencyID int) (entity.Currency, error)
	GetByCode(ctx context.Context, code enum.CurrencyCode) (entity.Currency, error)
	GetByName(ctx context.Context, name enum.CurrencyName) (entity.Currency, error)
	List(ctx context.Context) ([]entity.Currency, error)
	IsCodeExist(ctx context.Context, code enum.CurrencyCode) (bool, error)

	// Insert or update multiple currencies at once, useful for bulk operations, especially if you're syncing with an external source.
	BulkInsert(ctx context.Context, currencies []entity.Currency) error
	BulkUpdate(ctx context.Context, currencies []entity.Currency) error

	// Repository method to search currencies based on various attributes.
	Search(ctx context.Context, query string) ([]entity.Currency, error)

	// Retrieve the latest exchange rates for all currencies.
	GetLatestExchangeRates(ctx context.Context) (map[enum.CurrencyCode]float64, error)

	// Retrieve historical exchange rates of a currency between two dates.
	GetBetweenDates(ctx context.Context, code enum.CurrencyCode, startDate, endDate time.Time) ([]entity.Currency, error)

	// Retrieve top N and bottom N currencies based on the latest exchange rates. Useful for dashboards and analytics.
	GetTopNCurrencies(ctx context.Context, n int) ([]entity.Currency, error)
	GetBottomNCurrencies(ctx context.Context, n int) ([]entity.Currency, error)

	// Retrieve the average exchange rate of a currency over a specific duration.
	GetAverageExchangeRate(ctx context.Context, code enum.CurrencyCode) (float64, error)

	// query a database (or other data source) to find all countries associated with the given currency code.
	// To implement this, you'd need some data structure or table that maps currencies to countries.
	// CREATE TABLE public.currency_country_map (
	//	currency_code CHAR(3) REFERENCES public.currency(currency_code),
	//	country_name TEXT NOT NULL,
	//	PRIMARY KEY (currency_code, country_name)
	// );
	ListCountriesByCurrencyCode(ctx context.Context, code enum.CurrencyCode) ([]string, error)
}
