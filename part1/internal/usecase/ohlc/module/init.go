package module

import (
	rProducer "part1/internal/repository/mq"
	rRedis "part1/internal/repository/redis"
	usecase "part1/internal/usecase/ohlc"

	fsDatalogic "part1/internal/datalogic/fs"
	ohlcDatalogic "part1/internal/datalogic/ohlc"
)

type ohlcUsecase struct {
	redisRepo     rRedis.Repository
	producerRepo  rProducer.Repository
	fsDatalogic   fsDatalogic.Datalogic
	ohlcDatalogic ohlcDatalogic.Datalogic
}

func New(
	redisRepo rRedis.Repository,
	producerRepo rProducer.Repository,
	fsDatalogic fsDatalogic.Datalogic,
	ohlcDatalogic ohlcDatalogic.Datalogic,
) usecase.OhlcUsecase {
	return &ohlcUsecase{
		redisRepo:     redisRepo,
		producerRepo:  producerRepo,
		fsDatalogic:   fsDatalogic,
		ohlcDatalogic: ohlcDatalogic,
	}
}
