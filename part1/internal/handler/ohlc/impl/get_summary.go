package impl

import (
	"context"
	"part1/internal/model"
	pb "part1/internal/model/proto"
	"time"
)

func (h *handler) GetSummary(ctx context.Context, req *pb.GetSummaryRequest) (resp *pb.GetSummaryResponse, err error) {
	startTime := time.Now()

	summary, err := h.ohlcUsecase.GetSummary(ctx, model.GetSummaryRequest{
		StockName: req.GetStockName(),
	})
	if err != nil {
		return &pb.GetSummaryResponse{
			ResponseHeader: &pb.ResponseHeader{
				Success:      false,
				ErrorMessage: err.Error(),
				ProcessTime:  float64(time.Since(startTime).Milliseconds()),
			},
		}, nil
	}

	resp = &pb.GetSummaryResponse{
		ResponseHeader: &pb.ResponseHeader{
			Success:     true,
			ProcessTime: float64(time.Since(startTime).Milliseconds()),
		},
		StockName:     summary.StockName,
		PreviousPrice: summary.PrevPrice,
		OpenPrice:     summary.OpenPrice,
		HighestPrice:  summary.HighestPrice,
		LowestPrice:   summary.LowestPrice,
		ClosePrice:    summary.ClosePrice,
		Volume:        summary.Volume,
		Value:         summary.Value,
		AveragePrice:  summary.AveragePrice,
	}

	return resp, err
}
