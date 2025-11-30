package repository

import (
	"context"
	"gw-exchanger/internal/db"
	"gw-exchanger/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresRepository struct {
	q *db.Queries
}

func (r *PostgresRepository) GetAllRates(ctx context.Context) (pkg.ExchangeRatesMap, error) {
	rows, err := r.q.GetAllCurrencies(ctx)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	rates := make(pkg.ExchangeRatesMap, len(rows))

	for _, row := range rows {
		rates[row.Code] = row.Rate
	}

	return rates, nil
}

func (r *PostgresRepository) GetRate(ctx context.Context, currency pkg.Currency) (pkg.Rate, error) {
	row, err := r.q.GetCurrency(ctx, currency)
	if err != nil {
		zap.L().Error(err.Error())
		return 0, err
	}

	return row.Rate, nil
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	queries := db.New(pool)
	return &PostgresRepository{
		q: queries,
	}
}
