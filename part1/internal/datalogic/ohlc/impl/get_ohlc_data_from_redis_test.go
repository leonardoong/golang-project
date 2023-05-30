package impl

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"part1/internal/constant"
	"part1/internal/model"
	redisRepo "part1/internal/repository/redis"
)

func Test_GetOhlcDataFromRedis(t *testing.T) {
	redisRepoMock := new(redisRepo.MockRepository)

	type args struct {
		stockName string
	}

	req := args{
		stockName: "BBCA",
	}

	tests := []struct {
		name      string
		mock      func()
		args      args
		want      model.Ohlc
		wantExist bool
		wantErr   bool
	}{
		{
			name:      "fail HGetAll",
			args:      req,
			wantErr:   true,
			wantExist: false,
			mock: func() {
				redisRepoMock.
					On("HGetAll", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA")).
					Return(nil, errors.New("err HGetAll")).
					Once()
			},
		},
		{
			name:      "no result",
			args:      req,
			wantErr:   false,
			wantExist: false,
			mock: func() {
				redisRepoMock.
					On("HGetAll", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA")).
					Return(map[string]string{}, nil).
					Once()
			},
		},
		{
			name: "success",
			args: req,
			mock: func() {
				redisRepoMock.
					On("HGetAll", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA")).
					Return(map[string]string{
						"previous_price": "8000",
						"open_price":     "8050",
						"highest_price":  "8050",
						"lowest_price":   "8050",
						"close_price":    "8050",
						"average_price":  "8050",
						"volume":         "100",
						"value":          "805000",
					}, nil).
					Once()
			},
			want: model.Ohlc{
				PreviousPrice: 8000,
				OpenPrice:     8050,
				HighestPrice:  8050,
				LowestPrice:   8050,
				ClosePrice:    8050,
				AveragePrice:  8050,
				Volume:        100,
				Value:         805000,
			},
			wantExist: true,
			wantErr:   false,
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

			exist, got, err := d.GetOhlcDataFromRedis(context.Background(), tt.args.stockName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOhlcDataFromRedis test failed. want: %+v, got: %+v", tt.want, got)
			}
			if exist != tt.wantExist {
				t.Errorf("GetOhlcDataFromRedis test failed. want exist: %+v, got: %+v", tt.wantExist, exist)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOhlcDataFromRedis test failed. want err: %+v, got: %+v", tt.wantErr, err)
			}
		})
	}
}
