package repository

import (
	"context"
	"gw-exchanger/config"
	"gw-exchanger/internal/db"
	"gw-exchanger/pkg"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repository interface {
	GetAllCurrencies(ctx context.Context) ([]db.AppCurrency, error)
	GetTwoCurrencies(ctx context.Context, first, second pkg.Currency) (from, to *db.AppCurrency, err error)
}

func NewRepository(ctx context.Context, cfg *config.DatabaseConfig) (r Repository, closeConnection func()) {
	return NewPostgresRepository(ctx, cfg)
}
