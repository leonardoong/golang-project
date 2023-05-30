package module

import (
	"context"
	"part1/internal/model"
)

func (u *ohlcUsecase) ProcessOhlc(ctx context.Context, req model.Transaction) error {
	var (
		err error
	)

	// 1. get redis if exist
	exist, prevOhlc, err := u.ohlcDatalogic.GetOhlcDataFromRedis(ctx, req.StockCode)
	if err != nil {
		return err
	}

	summary, err := u.ohlcDatalogic.CalculateOhlc(ctx, exist, prevOhlc, req)
	if err != nil {
		return err
	}

	err = u.ohlcDatalogic.StoreOhlcData(ctx, summary)

	return err
}
