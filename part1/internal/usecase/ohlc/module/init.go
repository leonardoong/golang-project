package module

import (
	rProducer "part1/internal/repository/mq"
	rRedis "part1/internal/repository/redis"
	usecase "part1/internal/usecase/ohlc"
)

type ohlcUsecase struct {
	redisRepo    rRedis.Repository
	producerRepo rProducer.Repository
}

func New(
	redisRepo rRedis.Repository,
	producerRepo rProducer.Repository,
) usecase.OhlcUsecase {
	return &ohlcUsecase{
		redisRepo:    redisRepo,
		producerRepo: producerRepo,
	}
}
