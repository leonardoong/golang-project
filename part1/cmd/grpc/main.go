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
	server := grpc.NewServer()

	ohlcService := handlerOhlc.New()
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
