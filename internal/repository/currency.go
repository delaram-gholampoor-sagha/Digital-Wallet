package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type Currency struct {
	cli *sql.DB
}

func (repo *Currency) Insert(ctx context.Context, currency entity.Currency) error {
	query := `
		INSERT INTO currency (currency_code, currency_name, symbol, exchange_rate, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING currency_id
	`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.Currency.Insert.PrepareContext: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		currency.CurrencyCode,
		currency.CurrencyName,
		currency.Symbol,
		currency.ExchangeRate,
		currency.CreatedAt,
		currency.UpdatedAt,
	).Scan(&currency.CurrencyID)
	if err != nil {
		return fmt.Errorf("repository.Currency.Insert.QueryRowContext: %w", err)
	}

	return nil
}

func (repo *Currency) Update(ctx context.Context, currency entity.Currency) error {
	query := `
		UPDATE currency
		SET currency_code = $1, currency_name = $2, symbol = $3, exchange_rate = $4, updated_at = $5
		WHERE currency_id = $6
	`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.Currency.Update.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		currency.CurrencyCode,
		currency.CurrencyName,
		currency.Symbol,
		currency.ExchangeRate,
		currency.UpdatedAt,
		currency.CurrencyID,
	)
	if err != nil {
		return fmt.Errorf("repository.Currency.Update.ExecContext: %w", err)
	}

	return nil
}

func (repo *Currency) Delete(ctx context.Context, currencyID int) error {
	query := "UPDATE currency SET deleted_at = CURRENT_TIMESTAMP WHERE currency_id = $1"
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.Currency.Delete.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, currencyID)
	if err != nil {
		return fmt.Errorf("repository.Currency.Delete.ExecContext: %w", err)
	}

	return nil
}

func (repo *Currency) Get(ctx context.Context, currencyID int) (entity.Currency, error) {
	query := "SELECT * FROM currency WHERE currency_id = $1"
	row := repo.cli.QueryRowContext(ctx, query, currencyID)

	var currency entity.Currency
	err := row.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
	if err != nil {
		return entity.Currency{}, fmt.Errorf("repository.Currency.Get: %w", err)
	}

	return currency, nil
}

func (repo *Currency) GetByCode(ctx context.Context, code enum.CurrencyCode) (entity.Currency, error) {
	query := "SELECT * FROM currency WHERE currency_code = $1"
	row := repo.cli.QueryRowContext(ctx, query, code)

	var currency entity.Currency
	err := row.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
	if err != nil {
		return entity.Currency{}, fmt.Errorf("repository.Currency.GetByCode: %w", err)
	}

	return currency, nil
}

func (repo *Currency) GetByName(ctx context.Context, name enum.CurrencyName) (entity.Currency, error) {
	query := "SELECT * FROM currency WHERE currency_name = $1"
	row := repo.cli.QueryRowContext(ctx, query, name)

	var currency entity.Currency
	err := row.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
	if err != nil {
		return entity.Currency{}, fmt.Errorf("repository.Currency.GetByName: %w", err)
	}

	return currency, nil
}

func (repo *Currency) List(ctx context.Context) ([]entity.Currency, error) {
	query := "SELECT * FROM currency"
	rows, err := repo.cli.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.List.QueryContext: %w", err)
	}
	defer rows.Close()

	var currencies []entity.Currency
	for rows.Next() {
		var currency entity.Currency
		err = rows.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("repository.Currency.List.Scan: %w", err)
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (repo *Currency) IsCodeExist(ctx context.Context, code enum.CurrencyCode) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM currency WHERE currency_code = $1)"
	var exists bool
	err := repo.cli.QueryRowContext(ctx, query, code).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("repository.Currency.IsCodeExist.QueryRowContext: %w", err)
	}

	return exists, nil
}

func (repo *Currency) BulkInsert(ctx context.Context, currencies []entity.Currency) error {
	query := `
		INSERT INTO currency (currency_code, currency_name, symbol, exchange_rate, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repository.Currency.BulkInsert.BeginTx: %w", err)
	}

	for _, currency := range currencies {
		_, err = tx.ExecContext(ctx, query, currency.CurrencyCode, currency.CurrencyName, currency.Symbol, currency.ExchangeRate, time.Now(), time.Now())
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("repository.Currency.BulkInsert.ExecContext: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("repository.Currency.BulkInsert.Commit: %w", err)
	}

	return nil
}

func (repo *Currency) BulkUpdate(ctx context.Context, currencies []entity.Currency) error {
	query := `
		UPDATE currency
		SET currency_name = $1, symbol = $2, exchange_rate = $3, updated_at = $4
		WHERE currency_code = $5
	`

	tx, err := repo.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repository.Currency.BulkUpdate.BeginTx: %w", err)
	}

	for _, currency := range currencies {
		_, err = tx.ExecContext(ctx, query, currency.CurrencyName, currency.Symbol, currency.ExchangeRate, time.Now(), currency.CurrencyCode)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("repository.Currency.BulkUpdate.ExecContext: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("repository.Currency.BulkUpdate.Commit: %w", err)
	}

	return nil
}

func (repo *Currency) Search(ctx context.Context, query string) ([]entity.Currency, error) {
	sqlQuery := `
		SELECT * FROM currency
		WHERE LOWER(currency_name) LIKE LOWER($1)
		OR LOWER(currency_code) LIKE LOWER($1)
	`
	rows, err := repo.cli.QueryContext(ctx, sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.Search.QueryContext: %w", err)
	}
	defer rows.Close()

	var currencies []entity.Currency
	for rows.Next() {
		var currency entity.Currency
		err = rows.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("repository.Currency.Search.Scan: %w", err)
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (repo *Currency) GetLatestExchangeRates(ctx context.Context) (map[enum.CurrencyCode]float64, error) {
	query := "SELECT currency_code, exchange_rate FROM currency"
	rows, err := repo.cli.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.GetLatestExchangeRates.QueryContext: %w", err)
	}
	defer rows.Close()

	exchangeRates := make(map[enum.CurrencyCode]float64)
	for rows.Next() {
		var code enum.CurrencyCode
		var rate float64
		err = rows.Scan(&code, &rate)
		if err != nil {
			return nil, fmt.Errorf("repository.Currency.GetLatestExchangeRates.Scan: %w", err)
		}
		exchangeRates[code] = rate
	}

	return exchangeRates, nil
}

func (repo *Currency) GetBetweenDates(ctx context.Context, code enum.CurrencyCode, startDate, endDate time.Time) ([]entity.Currency, error) {
	query := `
		SELECT * FROM currency
		WHERE currency_code = $1
		AND (updated_at BETWEEN $2 AND $3)
	`
	rows, err := repo.cli.QueryContext(ctx, query, code, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.GetBetweenDates: %w", err)
	}
	defer rows.Close()

	var currencies []entity.Currency
	for rows.Next() {
		var currency entity.Currency
		err = rows.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("repository.Currency.GetBetweenDates.Scan: %w", err)
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (repo *Currency) GetTopNCurrencies(ctx context.Context, n int) ([]entity.Currency, error) {
	query := `
		SELECT * FROM currency
		ORDER BY exchange_rate DESC
		LIMIT $1
	`
	rows, err := repo.cli.QueryContext(ctx, query, n)
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.GetTopNCurrencies: %w", err)
	}
	defer rows.Close()

	var currencies []entity.Currency
	for rows.Next() {
		var currency entity.Currency
		err = rows.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("repository.Currency.GetTopNCurrencies.Scan: %w", err)
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (repo *Currency) GetBottomNCurrencies(ctx context.Context, n int) ([]entity.Currency, error) {
	query := `
		SELECT * FROM currency
		ORDER BY exchange_rate ASC
		LIMIT $1
	`
	rows, err := repo.cli.QueryContext(ctx, query, n)
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.GetBottomNCurrencies: %w", err)
	}
	defer rows.Close()

	var currencies []entity.Currency
	for rows.Next() {
		var currency entity.Currency
		err = rows.Scan(&currency.CurrencyID, &currency.CurrencyCode, &currency.CurrencyName, &currency.Symbol, &currency.ExchangeRate, &currency.CreatedAt, &currency.UpdatedAt, &currency.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("repository.Currency.GetBottomNCurrencies.Scan: %w", err)
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (repo *Currency) GetAverageExchangeRate(ctx context.Context, code enum.CurrencyCode) (float64, error) {
	query := `
		SELECT AVG(exchange_rate) FROM currency
		WHERE currency_code = $1
	`
	var avg float64
	err := repo.cli.QueryRowContext(ctx, query, code).Scan(&avg)
	if err != nil {
		return 0, fmt.Errorf("repository.Currency.GetAverageExchangeRate: %w", err)
	}
	return math.Round(avg*100) / 100, nil // Rounding to two decimal places
}

func (repo *Currency) ListCountriesByCurrencyCode(ctx context.Context, code enum.CurrencyCode) ([]string, error) {
	query := `
		SELECT DISTINCT country FROM country_currency_mapping
		WHERE currency_code = $1
	`
	rows, err := repo.cli.QueryContext(ctx, query, code)
	if err != nil {
		return nil, fmt.Errorf("repository.Currency.ListCountriesByCurrencyCode: %w", err)
	}
	defer rows.Close()

	var countries []string
	for rows.Next() {
		var country string
		if err := rows.Scan(&country); err != nil {
			return nil, fmt.Errorf("repository.Currency.ListCountriesByCurrencyCode.Scan: %w", err)
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func (repo *Currency) IsCurrrencyExist(ctx context.Context, currencyID int) (bool, error) {

	return false, nil
}
