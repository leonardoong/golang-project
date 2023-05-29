package redis

import "context"

type Repository interface {
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
}
