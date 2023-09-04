package currency

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"go.uber.org/zap"
)

func (s *Service) AddCurrency(ctx context.Context, req request.AddCurrency) (response.AddCurrency, error) {
	newCurrency := entity.Currency{
		CurrencyCode: req.CurrencyCode,
		CurrencyName: req.CurrencyName,
		Symbol:       req.Symbol,
		ExchangeRate: req.ExchangeRate,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}

	err := s.currencyRepo.Insert(ctx, newCurrency)
	if err != nil {
		s.logger.Error("Failed to add new currency", zap.Error(err))
		return response.AddCurrency{}, err
	}

	s.logger.Info("Successfully added new currency", zap.Any("currency", newCurrency))
	return response.AddCurrency{CurrencyID: newCurrency.CurrencyID}, nil
}

func (s *Service) UpdateCurrency(ctx context.Context, req request.UpdateCurrency) error {
	updatedCurrency := entity.Currency{
		CurrencyID:   req.CurrencyID,
		CurrencyCode: req.CurrencyCode,
		CurrencyName: req.CurrencyName,
		Symbol:       req.Symbol,
		ExchangeRate: req.ExchangeRate,
		UpdatedAt:    req.UpdatedAt,
	}

	err := s.currencyRepo.Update(ctx, updatedCurrency)
	if err != nil {
		s.logger.Error("Failed to update currency", zap.Error(err))
		return err
	}

	s.logger.Info("Successfully updated currency", zap.Int("currencyID", req.CurrencyID))
	return nil
}

func (s *Service) DeleteCurrency(ctx context.Context, currencyID int) error {
	err := s.currencyRepo.Delete(ctx, currencyID)
	if err != nil {
		s.logger.Error("Failed to delete currency", zap.Int("currencyID", currencyID), zap.Error(err))
		return err
	}

	s.logger.Info("Successfully deleted currency", zap.Int("currencyID", currencyID))
	return nil
}

func (s *Service) GetCurrency(ctx context.Context, currencyID int) (entity.Currency, error) {
	currency, err := s.currencyRepo.Get(ctx, currencyID)
	if err != nil {
		s.logger.Error("Failed to fetch currency by ID", zap.Int("currencyID", currencyID), zap.Error(err))
		return entity.Currency{}, err
	}

	if currency.CurrencyID == 0 {
		s.logger.Warn("Currency not found", zap.Int("currencyID", currencyID))
		return entity.Currency{}, fmt.Errorf("currency not found")
	}

	s.logger.Info("Successfully fetched currency by ID", zap.Int("currencyID", currencyID), zap.Any("currency", currency))
	return currency, nil
}

func (s *Service) GetCurrencyByName(ctx context.Context, currencyName enum.CurrencyName) (entity.Currency, error) {
	currency, err := s.currencyRepo.GetByName(ctx, currencyName)
	if err != nil {
		s.logger.Error("Failed to fetch currency by name", zap.String("currencyName", string(currencyName)), zap.Error(err))
		return entity.Currency{}, err
	}

	if currency.CurrencyID == 0 {
		s.logger.Warn("Currency not found", zap.String("currencyName", string(currencyName)))
		return entity.Currency{}, fmt.Errorf("currency not found")
	}

	s.logger.Info("Successfully fetched currency by name", zap.String("currencyName", string(currencyName)), zap.Any("currency", currency))
	return currency, nil
}

func (s *Service) ListCurrencies(ctx context.Context) ([]entity.Currency, error) {
	currencies, err := s.currencyRepo.List(ctx)
	if err != nil {
		s.logger.Error("Failed to list currencies", zap.Error(err))
		return nil, err
	}

	if len(currencies) == 0 {
		s.logger.Warn("No currencies found")
		return []entity.Currency{}, nil
	}

	s.logger.Info("Successfully listed currencies", zap.Int("count", len(currencies)))
	return currencies, nil
}

func (s *Service) GetExchangeRate(ctx context.Context, fromCode enum.CurrencyCode, toCode enum.CurrencyCode) (float64, error) {
	rates, err := s.currencyRepo.GetLatestExchangeRates(ctx)
	if err != nil {
		s.logger.Error("Failed to get latest exchange rates", zap.Error(err))
		return 0.0, err
	}

	fromRate, fromExists := rates[fromCode]
	toRate, toExists := rates[toCode]

	if !fromExists || !toExists {
		s.logger.Warn("Exchange rate not available for one or both currencies", zap.String("fromCode", string(fromCode)), zap.String("toCode", string(toCode)))
		return 0.0, fmt.Errorf("exchange rate not available for one or both currencies")
	}

	exchangeRate := toRate / fromRate
	s.logger.Info("Successfully fetched exchange rate", zap.String("fromCode", string(fromCode)), zap.String("toCode", string(toCode)), zap.Float64("exchangeRate", exchangeRate))
	return exchangeRate, nil
}

func (s *Service) BulkUpdateExchangeRates(ctx context.Context, rates map[enum.CurrencyCode]float64) error {
	var currenciesToUpdate []entity.Currency
	for code, rate := range rates {
		currenciesToUpdate = append(currenciesToUpdate, entity.Currency{
			CurrencyCode: code,
			ExchangeRate: &rate,
			UpdatedAt:    time.Now(),
		})
	}

	if err := s.currencyRepo.BulkUpdate(ctx, currenciesToUpdate); err != nil {
		s.logger.Error("Failed to bulk update exchange rates", zap.Error(err))
		return err
	}

	s.logger.Info("Successfully bulk updated exchange rates", zap.Int("count", len(rates)))
	return nil
}

func (s *Service) SearchCurrencies(ctx context.Context, query string) ([]entity.Currency, error) {
	currencies, err := s.currencyRepo.Search(ctx, query)
	if err != nil {
		s.logger.Error("Failed to search currencies", zap.String("query", query), zap.Error(err))
		return nil, err
	}

	if len(currencies) == 0 {
		s.logger.Warn("No currencies found for the search query", zap.String("query", query))
		return []entity.Currency{}, nil
	}

	s.logger.Info("Successfully searched currencies", zap.String("query", query), zap.Int("count", len(currencies)))
	return currencies, nil
}

func (s *Service) ConvertAmount(ctx context.Context, fromCode enum.CurrencyCode, toCode enum.CurrencyCode, amount float64) (float64, error) {
	exchangeRate, err := s.GetExchangeRate(ctx, fromCode, toCode)
	if err != nil {
		s.logger.Error("Failed to get exchange rate for currency conversion", zap.Error(err))
		return 0.0, err
	}

	convertedAmount := amount * exchangeRate
	s.logger.Info("Successfully converted amount", zap.Float64("convertedAmount", convertedAmount))
	return convertedAmount, nil
}

func (s *Service) CompareCurrencies(ctx context.Context, firstCode enum.CurrencyCode, secondCode enum.CurrencyCode) (response.CurrencyComparison, error) {
	firstCurrency, err1 := s.currencyRepo.GetByCode(ctx, firstCode)
	secondCurrency, err2 := s.currencyRepo.GetByCode(ctx, secondCode)

	if err1 != nil || err2 != nil {
		s.logger.Error("Failed to fetch currencies for comparison", zap.Error(err1), zap.Error(err2))
		return response.CurrencyComparison{}, fmt.Errorf("failed to fetch currencies for comparison: %v %v", err1, err2)
	}

	isStronger := firstCurrency.ExchangeRate != nil && secondCurrency.ExchangeRate != nil && *firstCurrency.ExchangeRate > *secondCurrency.ExchangeRate
	comparisonRate := 0.0
	if firstCurrency.ExchangeRate != nil && secondCurrency.ExchangeRate != nil {
		comparisonRate = *firstCurrency.ExchangeRate / *secondCurrency.ExchangeRate
	}

	result := response.CurrencyComparison{
		FirstCurrency:  firstCurrency,
		SecondCurrency: secondCurrency,
		IsStronger:     isStronger,
		ComparisonRate: comparisonRate,
	}

	s.logger.Info("Successfully compared currencies", zap.Bool("isStronger", isStronger))
	return result, nil
}

func (s *Service) GetCurrencyTrends(ctx context.Context, code enum.CurrencyCode, duration time.Duration) (response.CurrencyTrends, error) {
	endDate := time.Now()
	startDate := endDate.Add(-duration)

	historicalRates, err := s.currencyRepo.GetBetweenDates(ctx, code, startDate, endDate)
	if err != nil {
		s.logger.Error("Failed to fetch historical exchange rates", zap.Error(err))
		return response.CurrencyTrends{}, err
	}

	var trends []struct {
		Date         time.Time
		ExchangeRate float64
	}

	for _, rate := range historicalRates {
		if rate.ExchangeRate != nil {
			trends = append(trends, struct {
				Date         time.Time
				ExchangeRate float64
			}{
				Date:         rate.UpdatedAt,
				ExchangeRate: *rate.ExchangeRate,
			})
		}
	}

	result := response.CurrencyTrends{
		Code:   code,
		Trends: trends,
	}

	s.logger.Info("Successfully fetched currency trends", zap.Int("numberOfEntries", len(trends)))
	return result, nil
}

func (s *Service) GetStrongestCurrency(ctx context.Context) (entity.Currency, error) {
	currencies, err := s.currencyRepo.GetTopNCurrencies(ctx, 1)
	if err != nil {
		s.logger.Error("Failed to fetch strongest currency", zap.Error(err))
		return entity.Currency{}, err
	}

	if len(currencies) == 0 {
		return entity.Currency{}, errors.New("no currencies available")
	}

	s.logger.Info("Successfully fetched the strongest currency", zap.String("CurrencyCode", string(currencies[0].CurrencyCode)))
	return currencies[0], nil
}

func (s *Service) GetWeakestCurrency(ctx context.Context) (entity.Currency, error) {
	currencies, err := s.currencyRepo.GetBottomNCurrencies(ctx, 1)
	if err != nil {
		s.logger.Error("Failed to fetch weakest currency", zap.Error(err))
		return entity.Currency{}, err
	}

	if len(currencies) == 0 {
		return entity.Currency{}, errors.New("no currencies available")
	}

	s.logger.Info("Successfully fetched the weakest currency", zap.String("CurrencyCode", string(currencies[0].CurrencyCode)))
	return currencies[0], nil
}

func (s *Service) NotifyUsersOnExchangeRateChange(ctx context.Context, threshold float64) error {
	// fetching the user data and their notification settings,
	s.logger.Info("Notifying users about exchange rate changes that exceed the threshold", zap.Float64("threshold", threshold))
	return nil
}

func (s *Service) GetCountriesUsingCurrency(ctx context.Context, code enum.CurrencyCode) ([]string, error) {
	countries, err := s.currencyRepo.ListCountriesByCurrencyCode(ctx, code)
	if err != nil {
		s.logger.Error("Failed to fetch countries using the currency", zap.Error(err))
		return nil, err
	}

	s.logger.Info("Successfully fetched countries using the currency", zap.String("CurrencyCode", string(code)), zap.Strings("countries", countries))
	return countries, nil
}
