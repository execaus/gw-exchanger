package handler

import (
	"context"
	gw_grpc "gw-exchanger/internal/pb/exchange"
	"gw-exchanger/internal/service"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	s *service.Service
	gw_grpc.UnimplementedExchangeServiceServer
}

func (h *Handler) GetExchangeRates(ctx context.Context, _ *emptypb.Empty) (*gw_grpc.ExchangeRatesResponse, error) {
	m, err := h.s.GetExchangeRates(ctx)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	return &gw_grpc.ExchangeRatesResponse{
		Rates: m,
	}, nil
}

func (h *Handler) GetExchangeRateForCurrency(ctx context.Context, in *gw_grpc.CurrencyRequest) (*gw_grpc.ExchangeRateResponse, error) {
	rate, err := h.s.GetExchangeRateForCurrency(ctx, in.FromCurrency, in.ToCurrency)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	return &gw_grpc.ExchangeRateResponse{
		FromCurrency: in.FromCurrency,
		ToCurrency:   in.ToCurrency,
		Rate:         rate,
	}, nil
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		s: srv,
	}
}
