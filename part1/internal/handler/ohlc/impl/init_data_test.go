package impl

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"part1/internal/model"
	pb "part1/internal/model/proto"
	ohlcUsecase "part1/internal/usecase/ohlc"

	"github.com/golang/protobuf/ptypes/empty"
)

func Test_InitData(t *testing.T) {
	ohlcUsecase := new(ohlcUsecase.MockOhlcUsecase)

	var (
		void *empty.Empty
		ctx  = context.Background()
	)

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
		want    *pb.InitDataResponse
	}{
		{
			name: "failed get data json",
			mock: func() {
				ohlcUsecase.On("InitData", ctx).
					Return(model.InitDataResponse{
						Success:      false,
						ErrorMessage: "error init data",
					}, errors.New("error init data")).
					Once()
			},
			want: &pb.InitDataResponse{
				ResponseHeader: &pb.ResponseHeader{
					Success:      false,
					ErrorMessage: "error init data",
					ProcessTime:  0,
				},
			},
		},
		{
			name: "success",
			mock: func() {
				ohlcUsecase.On("InitData", ctx).
					Return(model.InitDataResponse{
						Success: true,
					}, nil).
					Once()
			},
			want: &pb.InitDataResponse{
				ResponseHeader: &pb.ResponseHeader{
					Success:     true,
					ProcessTime: 0,
				},
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

			got, err := h.InitData(ctx, void)
			// normalize process time
			got.ResponseHeader.ProcessTime = 0
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitData failed got : %v, err %v \n", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("InitData failed want err : %v, err %v \n", tt.wantErr, err)
			}
		})
	}
}
