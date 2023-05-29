package impl

import (
	"github.com/redis/go-redis/v9"

	redisRepo "part1/internal/repository/redis"
)

type repository struct {
	redisConn *redis.Client
}

func New(
	redisConn *redis.Client,
) redisRepo.Repository {
	return &repository{
		redisConn: redisConn,
	}
}
