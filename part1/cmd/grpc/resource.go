package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Resource struct {
	RedisConn *redis.Client
}

func initResource(cfg Config) (Resource, error) {
	fmt.Printf("cfg == %v \n", cfg.Redis.Host)
	fmt.Printf("cfg 2 == %v \n", cfg.Redis.Password)

	redisConn := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	return Resource{
		RedisConn: redisConn,
	}, nil
}
