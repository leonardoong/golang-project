package impl

import (
	ohlcDatalogic "part1/internal/datalogic/ohlc"

	redisRepo "part1/internal/repository/redis"
)

type datalogic struct {
	redisRepo redisRepo.Repository
}

func New(
	redisRepo redisRepo.Repository,
) ohlcDatalogic.Datalogic {
	return &datalogic{
		redisRepo: redisRepo,
	}
}
