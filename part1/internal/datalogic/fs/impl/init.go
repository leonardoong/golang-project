package impl

import (
	fsDatalogic "part1/internal/datalogic/fs"

	kafkaProducer "part1/internal/repository/mq"
)

type datalogic struct {
	producer kafkaProducer.Repository
}

func New(
	producer kafkaProducer.Repository,
) fsDatalogic.Datalogic {
	return &datalogic{
		producer: producer,
	}
}
