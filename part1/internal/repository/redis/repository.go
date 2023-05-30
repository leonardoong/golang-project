package redis

import "context"

type Repository interface {
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}
