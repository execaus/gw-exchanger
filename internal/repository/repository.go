package repository

import (
	"context"
	"gw-exchanger/pkg"
)

type Repository interface {
	GetAllRates(ctx context.Context) (pkg.ExchangeRatesMap, error)
	GetRate(ctx context.Context, currency pkg.Currency) (pkg.Rate, error)
}
