package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handlerOhlc "part1/internal/handler/ohlc/impl"

	pb "part1/internal/model/proto"

	"part1/internal/common"
)

func main() {
	cfg, err := common.ReadConfig()
	if err != nil {
		log.Fatalf("failed to read config, error : %v\n", err)
	}

	res, err := common.InitResource(cfg, false)
	if err != nil {
		log.Fatalf("failed to init resource, error : %v\n", err)
	}

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
