package fs

import (
	"context"
	"part1/internal/model"
)

type Datalogic interface {
	GetOhlcDataFromRedis(ctx context.Context, stockCode string) (bool, model.Ohlc, error)
	CalculateOhlc(ctx context.Context, prevExist bool, prevOhlc model.Ohlc, req model.Transaction) (model.Ohlc, error)
	StoreOhlcData(ctx context.Context, summary model.Ohlc) error
}
