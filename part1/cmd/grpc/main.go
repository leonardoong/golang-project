package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handlerOhlc "part1/internal/handler/ohlc/impl"

	pb "part1/internal/model/proto"
)

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("failed to read config, error : %v\n", err)
	}

	res, err := initResource(cfg)
	if err != nil {
		log.Fatalf("failed to init resource, error : %v\n", err)
	}

	// err = res.RedisConn.Set(context.Background(), "test", "value", 0).Err()
	// if err != nil {
	// 	log.Fatalf("failed to set redis, error : %v\n", err)
	// }

	// val, err := res.RedisConn.Get(context.Background(), "test").Result()
	// if err != nil {
	// 	log.Fatalf("failed to get redis, error : %v\n", err)
	// }
	// fmt.Println("key", val)

	usecase := initUsecase(&res)

	server := grpc.NewServer()

	ohlcService := handlerOhlc.New(
		usecase.OhlcUsecase,
	)
	pb.RegisterOhlcServer(server, ohlcService)
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to create listener, error : %v\n", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve, error : %v\n", err)
	}
}
