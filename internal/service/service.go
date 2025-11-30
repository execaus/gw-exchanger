package service

import (
	"context"
	"errors"
	"gw-exchanger/internal/repository"
	"gw-exchanger/pkg"

	"go.uber.org/zap"
)

type Service struct {
	r repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		r: repo,
	}
}

func (s *Service) GetExchangeRates(ctx context.Context) (pkg.ExchangeRatesMap, error) {
	rows, err := s.r.GetAllCurrencies(ctx)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	m := make(pkg.ExchangeRatesMap, len(rows))
	for _, row := range rows {
		m[row.Code] = row.Rate
	}

	return m, err
}

func (s *Service) GetExchangeRateForCurrency(ctx context.Context, from, to pkg.Currency) (pkg.Rate, error) {
	fromCurrency, toCurrency, err := s.r.GetTwoCurrencies(ctx, from, to)
	if err != nil {
		zap.L().Error(err.Error())
		return 0, err
	}

	if fromCurrency.Rate == 0 || toCurrency.Rate == 0 {
		err = errors.New("invalid rate for currency: zero")
		zap.L().Error(err.Error())
		return 0, err
	}

	return toCurrency.Rate / fromCurrency.Rate, nil
}
