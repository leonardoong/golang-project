package main

import (
	ohlcUsecase "part1/internal/usecase/ohlc"
	ohlcUsecaseModule "part1/internal/usecase/ohlc/module"

	rKafkaImpl "part1/internal/repository/mq/impl"
	rRedisImpl "part1/internal/repository/redis/impl"
)

type Usecase struct {
	OhlcUsecase ohlcUsecase.OhlcUsecase
}

func initUsecase(res *Resource) Usecase {
	return Usecase{
		OhlcUsecase: initOhlcUsecase(res),
	}
}

func initOhlcUsecase(res *Resource) ohlcUsecase.OhlcUsecase {
	redisRepo := rRedisImpl.New(res.RedisConn)
	producerRepo := rKafkaImpl.New(res.KafkaProducer)

	return ohlcUsecaseModule.New(redisRepo, producerRepo)
}
