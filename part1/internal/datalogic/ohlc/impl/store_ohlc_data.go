package impl

import (
	"context"
	"fmt"
	"part1/internal/constant"
	"part1/internal/model"
)

func (d *datalogic) StoreOhlcData(ctx context.Context, summary model.Ohlc) error {

	key := fmt.Sprintf(constant.OhlcRedisKey, summary.StockCode)
	_, err := d.redisRepo.HSet(ctx, key, summary)
	if err != nil {
		return err
	}

	return nil
}
