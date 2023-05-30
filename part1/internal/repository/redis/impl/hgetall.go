package impl

import (
	"context"
)

func (r *repository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.redisConn.HGetAll(ctx, key).Result()
}
