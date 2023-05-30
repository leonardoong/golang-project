package ohlc

import (
	"context"
	"part1/internal/model"
)

type OhlcUsecase interface {
	InitData(ctx context.Context) (resp model.InitDataResponse, err error)
	ProcessOhlc(ctx context.Context, req model.Transaction) error
	GetSummary(ctx context.Context, req model.GetSummaryRequest) (resp model.GetSummaryResponse, err error)
}
