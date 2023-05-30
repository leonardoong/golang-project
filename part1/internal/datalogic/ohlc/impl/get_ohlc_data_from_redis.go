package impl

import (
	"context"
	"fmt"
	"strconv"

	"part1/internal/constant"
	"part1/internal/model"
)

func (d *datalogic) GetOhlcDataFromRedis(ctx context.Context, stockCode string) (exist bool, res model.Ohlc, err error) {
	cache, err := d.redisRepo.HGetAll(ctx, fmt.Sprintf(constant.OhlcRedisKey, stockCode))
	if err != nil {
		return false, res, err
	}
	if len(cache) == 0 {
		return false, res, err
	}
	var (
		StockCode       string
		PreviousPrice   int64
		OpenPrice       int64
		HighestPrice    int64
		LowestPrice     int64
		ClosePrice      int64
		AveragePrice    int64
		Volume          int64
		Value           int64
		LastOrderNumber int64
	)

	if val, ok := cache["previous_price"]; ok && val != "" {
		PreviousPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["open_price"]; ok && val != "" {
		OpenPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["highest_price"]; ok && val != "" {
		HighestPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["lowest_price"]; ok && val != "" {
		LowestPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["close_price"]; ok && val != "" {
		ClosePrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["average_price"]; ok && val != "" {
		AveragePrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["volume"]; ok && val != "" {
		Volume, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["value"]; ok && val != "" {
		Value, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["last_order_number"]; ok && val != "" {
		LastOrderNumber, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["stock_code"]; ok && val != "" {
		StockCode = val
	}

	return true, model.Ohlc{
		StockCode:       StockCode,
		PreviousPrice:   PreviousPrice,
		OpenPrice:       OpenPrice,
		HighestPrice:    HighestPrice,
		LowestPrice:     LowestPrice,
		ClosePrice:      ClosePrice,
		AveragePrice:    AveragePrice,
		Volume:          Volume,
		Value:           Value,
		LastOrderNumber: LastOrderNumber,
	}, err
}
