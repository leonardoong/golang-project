package impl

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"part1/internal/constant"
	"part1/internal/model"
	redisRepo "part1/internal/repository/redis"
)

func Test_StoreOhlc(t *testing.T) {
	redisRepoMock := new(redisRepo.MockRepository)

	type args struct {
		summary model.Ohlc
	}

	req := args{
		summary: model.Ohlc{
			StockCode:       "BBCA",
			PreviousPrice:   8000,
			OpenPrice:       8050,
			HighestPrice:    8050,
			LowestPrice:     8050,
			ClosePrice:      8050,
			AveragePrice:    8050,
			Volume:          100,
			Value:           805000,
			LastOrderNumber: 1231231232,
		},
	}
	key := fmt.Sprintf(constant.OhlcRedisKey, req.summary.StockCode)

	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name:    "fail HSet",
			args:    req,
			wantErr: true,
			mock: func() {

				redisRepoMock.
					On("HSet", context.Background(), key, req.summary).
					Return(int64(0), errors.New("err HSet")).
					Once()
			},
		},
		{
			name:    "success",
			args:    req,
			wantErr: false,
			mock: func() {
				redisRepoMock.
					On("HSet", context.Background(), key, req.summary).
					Return(int64(0), nil).
					Once()
			},
		},
	}

	for _, tt := range tests {
		d := datalogic{
			redisRepo: redisRepoMock,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			err := d.StoreOhlcData(context.Background(), tt.args.summary)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreOhlcData failed want err : %v, got err : %v", tt.wantErr, err)
			}
		})
	}
}
