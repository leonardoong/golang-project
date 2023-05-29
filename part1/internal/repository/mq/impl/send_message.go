package impl

import (
	"context"

	"github.com/Shopify/sarama"
)

func (r *repository) SendMessage(ctx context.Context, topic, value string) (int32, int64, error) {
	return r.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	})
}
