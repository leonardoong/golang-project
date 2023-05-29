package impl

import (
	"context"
)

func (r *repository) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.redisConn.HSet(ctx, key, values...).Result()
}
