package impl

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"part1/internal/model"
	pb "part1/internal/model/proto"
	ohlcUsecase "part1/internal/usecase/ohlc"
)

func Test_GetSummary(t *testing.T) {
	ohlcUsecase := new(ohlcUsecase.MockOhlcUsecase)

	var (
		request = &pb.GetSummaryRequest{
			StockName: "BBCA",
		}
		ctx = context.Background()
	)

	type args struct {
		req *pb.GetSummaryRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
		want    *pb.GetSummaryResponse
	}{
		{
			name: "failed get data json",
			args: args{
				req: request,
			},
			mock: func() {
				ohlcUsecase.On("GetSummary", ctx, model.GetSummaryRequest{
					StockName: request.StockName,
				}).
					Return(model.GetSummaryResponse{}, errors.New("error get summary")).
					Once()
			},
			want: &pb.GetSummaryResponse{
				ResponseHeader: &pb.ResponseHeader{
					Success:      false,
					ErrorMessage: "error get summary",
					ProcessTime:  0,
				},
			},
		},
		{
			name: "success",
			args: args{
				req: request,
			},
			mock: func() {
				ohlcUsecase.On("GetSummary", ctx, model.GetSummaryRequest{
					StockName: request.StockName,
				}).
					Return(model.GetSummaryResponse{
						StockName:    "BBCA",
						OpenPrice:    9000,
						PrevPrice:    9050,
						HighestPrice: 10500,
						LowestPrice:  9000,
						ClosePrice:   10200,
						AveragePrice: 10000,
						Volume:       2000000,
						Value:        1000000,
					}, nil).
					Once()
			},
			want: &pb.GetSummaryResponse{
				ResponseHeader: &pb.ResponseHeader{
					Success:     true,
					ProcessTime: 0,
				},
				StockName:     "BBCA",
				OpenPrice:     9000,
				PreviousPrice: 9050,
				HighestPrice:  10500,
				LowestPrice:   9000,
				ClosePrice:    10200,
				AveragePrice:  10000,
				Volume:        2000000,
				Value:         1000000,
			},
		},
	}

	for _, tt := range tests {
		h := handler{
			ohlcUsecase: ohlcUsecase,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, err := h.GetSummary(ctx, tt.args.req)
			// normalize process time
			got.ResponseHeader.ProcessTime = 0
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSummary failed got : %v, err %v \n", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummary failed want err : %v, err %v \n", tt.wantErr, err)
			}
		})
	}
}
