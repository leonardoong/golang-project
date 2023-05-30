package impl

import (
	"context"
	"math"
	"part1/internal/model"
	"strconv"
)

func (d *datalogic) CalculateOhlc(ctx context.Context, prevExist bool, prevOhlc model.Ohlc, req model.Transaction) (res model.Ohlc, err error) {

	if !prevExist {
		if req.Quantity == "" {
			price := int64(0)
			if req.Price != "" {
				price, err = strconv.ParseInt(req.Price, 10, 64)
				if err != nil {
					return
				}
				res.PreviousPrice = price
				res.StockCode = req.StockCode
			}
		}
	} else {
		res = prevOhlc

		// in case type E/P already set redis first
		if req.Quantity == "" {
			price := int64(0)
			if req.Price != "" {
				price, err = strconv.ParseInt(req.Price, 10, 64)
				if err != nil {
					return
				}
				res.PreviousPrice = price
				res.StockCode = req.StockCode
			}
		}

		if req.Type == "E" || req.Type == "P" {
			executionPrice := int64(0)
			if req.ExecutionPrice != "" {
				executionPrice, err = strconv.ParseInt(req.ExecutionPrice, 10, 64)
				if err != nil {
					return
				}
			}

			executedQty := int64(0)
			if req.ExecutedQuantity != "" {
				executedQty, err = strconv.ParseInt(req.ExecutedQuantity, 10, 64)
				if err != nil {
					return
				}
			}

			if prevOhlc.OpenPrice == 0 {
				res.OpenPrice = executionPrice
			}
			if prevOhlc.HighestPrice == 0 || prevOhlc.HighestPrice < executionPrice {
				res.HighestPrice = executionPrice
			}
			if prevOhlc.LowestPrice == 0 || prevOhlc.LowestPrice > executionPrice {
				res.LowestPrice = executionPrice
			}
			res.Volume += executedQty
			res.Value += executedQty * executionPrice

			res.AveragePrice = int64(math.Round(float64(res.Value) / float64(res.Volume)))

			// assuming the last order number will be the last transaction means closing price
			orderNumber := int64(0)
			if req.OrderNumber != "" {
				orderNumber, err = strconv.ParseInt(req.OrderNumber, 10, 64)
				if err != nil {
					return
				}
			}

			if prevOhlc.LastOrderNumber == 0 || prevOhlc.LastOrderNumber < orderNumber {
				res.LastOrderNumber = orderNumber
				res.ClosePrice = executionPrice
			}
		}
	}
	return res, err
}
