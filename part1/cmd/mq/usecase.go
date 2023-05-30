package main

import (
	"part1/internal/common"
	ohlcUsecase "part1/internal/usecase/ohlc"
	ohlcUsecaseModule "part1/internal/usecase/ohlc/module"

	rKafkaImpl "part1/internal/repository/mq/impl"
	rRedisImpl "part1/internal/repository/redis/impl"

	dFs "part1/internal/datalogic/fs/impl"
	dOhlc "part1/internal/datalogic/ohlc/impl"
)

type Usecase struct {
	OhlcUsecase ohlcUsecase.OhlcUsecase
}

func initUsecase(res *common.Resource) Usecase {
	return Usecase{
		OhlcUsecase: initOhlcUsecase(res),
	}
}

func initOhlcUsecase(res *common.Resource) ohlcUsecase.OhlcUsecase {
	redisRepo := rRedisImpl.New(res.RedisConn)
	producerRepo := rKafkaImpl.New(res.KafkaProducer)

	dFs := dFs.New(producerRepo)
	dOhlc := dOhlc.New(redisRepo)

	return ohlcUsecaseModule.New(redisRepo, producerRepo, dFs, dOhlc)
}
