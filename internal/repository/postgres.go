package repository

import (
	"context"
	"errors"
	"fmt"
	"gw-exchanger/config"
	"gw-exchanger/internal/db"
	"gw-exchanger/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresRepository struct {
	q *db.Queries
}

func (r *PostgresRepository) GetAllCurrencies(ctx context.Context) ([]db.AppCurrency, error) {
	rows, err := r.q.GetAllCurrencies(ctx)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	rates := make(pkg.ExchangeRatesMap, len(rows))

	for _, row := range rows {
		rates[row.Code] = row.Rate
	}

	return rows, nil
}

func (r *PostgresRepository) GetTwoCurrencies(ctx context.Context, fromCode, toCode pkg.Currency) (from, to *db.AppCurrency, err error) {
	rows, err := r.q.GetTwoCurrencies(ctx, db.GetTwoCurrenciesParams{
		Code:   fromCode,
		Code_2: toCode,
	})
	if err != nil {
		zap.L().Error(err.Error())
		return nil, nil, err
	}

	if len(rows) != 2 {
		err = errors.New("expected exactly two currencies")
		zap.L().Error(err.Error())
		return nil, nil, err
	}

	if rows[0].Code == fromCode {
		return &rows[0], &rows[1], nil
	} else {
		return &rows[1], &rows[0], nil
	}
}

func NewPostgresRepository(ctx context.Context, cfg *config.DatabaseConfig) (r *PostgresRepository, closeConnection func()) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		zap.L().Fatal("unable to create pgxpool", zap.Error(err))
	}

	if err = pool.Ping(ctx); err != nil {
		zap.L().Fatal("unable to ping database", zap.Error(err))
	}

	zap.L().Info("connected to PostgreSQL")

	queries := db.New(pool)

	return &PostgresRepository{
		q: queries,
	}, pool.Close
}
