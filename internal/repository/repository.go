package repository

import (
	"context"
	"gw-exchanger/internal/db"
	"gw-exchanger/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repository interface {
	GetAllCurrencies(ctx context.Context) ([]db.AppCurrency, error)
	GetTwoCurrencies(ctx context.Context, first, second pkg.Currency) (from, to *db.AppCurrency, err error)
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return NewPostgresRepository(pool)
}
