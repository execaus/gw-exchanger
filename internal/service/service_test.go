package service

import (
	"context"
	"errors"
	"testing"

	"gw-exchanger/internal/db"
	mock_repository "gw-exchanger/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetExchangeRateForCurrency_ValidRates_Success(t *testing.T) {
	ctx := context.Background()
	currencyFrom := "USD"
	currencyTo := "EUR"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	srv := NewService(mockRepo)

	from := &db.AppCurrency{Rate: 1.52}
	to := &db.AppCurrency{Rate: 1.23}

	mockRepo.EXPECT().GetTwoCurrencies(ctx, currencyFrom, currencyTo).Return(from, to, nil)

	rate, err := srv.GetExchangeRateForCurrency(ctx, currencyFrom, currencyTo)
	assert.NoError(t, err)
	assert.Equal(t, to.Rate/from.Rate, rate)
}

func TestGetExchangeRateForCurrency_RepoError_ReturnsError(t *testing.T) {
	ctx := context.Background()
	currencyFrom := "USD"
	currencyTo := "EUR"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	srv := NewService(mockRepo)

	mockRepo.EXPECT().
		GetTwoCurrencies(ctx, currencyFrom, currencyTo).
		Return((*db.AppCurrency)(nil), (*db.AppCurrency)(nil), errors.New("failed to get rate"))

	rate, err := srv.GetExchangeRateForCurrency(ctx, currencyFrom, currencyTo)
	assert.Error(t, err)
	assert.Equal(t, float32(0), rate)
	assert.Equal(t, "failed to get rate", err.Error())
}

func TestGetExchangeRateForCurrency_ZeroRate_Error(t *testing.T) {
	ctx := context.Background()
	currencyFrom := "USD"
	currencyTo := "EUR"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	srv := NewService(mockRepo)

	mockRepo.EXPECT().
		GetTwoCurrencies(ctx, currencyFrom, currencyTo).
		Return(&db.AppCurrency{Rate: 0}, &db.AppCurrency{Rate: 0}, nil)

	rate, err := srv.GetExchangeRateForCurrency(ctx, currencyFrom, currencyTo)

	assert.Error(t, err)
	assert.Equal(t, float32(0), rate)
	assert.Error(t, err)
}
