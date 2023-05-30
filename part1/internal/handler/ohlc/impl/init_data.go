package impl

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	pb "part1/internal/model/proto"
)

func (h *handler) InitData(ctx context.Context, voidReq *empty.Empty) (resp *pb.InitDataResponse, err error) {

	startTime := time.Now()

	res, err := h.ohlcUsecase.InitData(ctx)

	return &pb.InitDataResponse{
		ResponseHeader: &pb.ResponseHeader{
			Success:      res.Success,
			ErrorMessage: res.ErrorMessage,
			ProcessTime:  float64(time.Since(startTime).Milliseconds()),
		},
	}, nil
}
