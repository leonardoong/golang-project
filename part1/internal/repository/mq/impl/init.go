package impl

import (
	kafkaRepo "part1/internal/repository/mq"

	"github.com/Shopify/sarama"
)

type repository struct {
	kafkaProducer sarama.SyncProducer
}

func New(
	kafkaProducer sarama.SyncProducer,
) kafkaRepo.Repository {
	return &repository{
		kafkaProducer: kafkaProducer,
	}
}
