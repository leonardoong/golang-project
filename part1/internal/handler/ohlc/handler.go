package ohlc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "part1/internal/model/proto"
)

type OhlcHandler interface {
	InitData(ctx context.Context, void *empty.Empty) (*pb.InitDataResponse, error)
	GetSummary(ctx context.Context, req *pb.GetSummaryRequest) (resp *pb.GetSummaryResponse, err error)
}
