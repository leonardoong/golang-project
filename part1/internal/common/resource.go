package common

import (
	"github.com/redis/go-redis/v9"

	"github.com/Shopify/sarama"
)

type Resource struct {
	RedisConn     *redis.Client
	KafkaProducer sarama.SyncProducer
}

func InitResource(cfg Config, skipProducer bool) (Resource, error) {
	redisConn := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	var (
		kafkaConn sarama.SyncProducer
		err       error
	)

	if !skipProducer {
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 10
		config.Producer.Return.Successes = true
		config.Producer.MaxMessageBytes = 200000000

		kafkaConn, err = sarama.NewSyncProducer([]string{cfg.Kafka.Broker}, config)
		if err != nil {
			panic(err)
		}
	}

	return Resource{
		RedisConn:     redisConn,
		KafkaProducer: kafkaConn,
	}, nil
}
