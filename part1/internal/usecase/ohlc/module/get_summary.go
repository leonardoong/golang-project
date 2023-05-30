package module

import (
	"context"
	"errors"
	"part1/internal/model"
)

func (u *ohlcUsecase) GetSummary(ctx context.Context, req model.GetSummaryRequest) (resp model.GetSummaryResponse, err error) {
	exist, ohlcSummary, err := u.ohlcDatalogic.GetOhlcDataFromRedis(ctx, req.StockName)
	if err != nil {
		return resp, err
	}

	if !exist {
		return resp, errors.New("call init data first or stock not exist")
	}

	resp = model.GetSummaryResponse{
		StockName:    ohlcSummary.StockCode,
		PrevPrice:    ohlcSummary.PreviousPrice,
		OpenPrice:    ohlcSummary.OpenPrice,
		HighestPrice: ohlcSummary.HighestPrice,
		LowestPrice:  ohlcSummary.LowestPrice,
		ClosePrice:   ohlcSummary.ClosePrice,
		Volume:       ohlcSummary.Volume,
		Value:        ohlcSummary.Value,
		AveragePrice: ohlcSummary.AveragePrice,
	}

	return resp, err
}
