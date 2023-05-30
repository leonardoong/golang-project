package module

import (
	"context"
	"errors"
	"part1/internal/model"
	"reflect"
	"testing"

	ohlcDatalogic "part1/internal/datalogic/ohlc"
)

func Test_GetSummary(t *testing.T) {
	ohlcDatalogicMock := new(ohlcDatalogic.MockDatalogic)

	ctx := context.Background()

	type args struct {
		req model.GetSummaryRequest
	}

	request := model.GetSummaryRequest{
		StockName: "BBCA",
	}

	summary := model.Ohlc{
		StockCode:       "BBCA",
		OpenPrice:       9000,
		PreviousPrice:   9050,
		HighestPrice:    10500,
		LowestPrice:     9000,
		ClosePrice:      10200,
		AveragePrice:    10000,
		Volume:          2000000,
		Value:           1000000,
		LastOrderNumber: 123123123,
	}

	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
		want    model.GetSummaryResponse
	}{
		{
			name: "failed get redis",
			args: args{
				req: request,
			},
			mock: func() {
				ohlcDatalogicMock.On("GetOhlcDataFromRedis", ctx, request.StockName).
					Return(false, model.Ohlc{}, errors.New("error redis")).
					Once()
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				req: request,
			},
			mock: func() {
				exist := true
				ohlcDatalogicMock.On("GetOhlcDataFromRedis", ctx, request.StockName).
					Return(exist, summary, nil).
					Once()
			},
			want: model.GetSummaryResponse{
				StockName:    "BBCA",
				OpenPrice:    9000,
				PrevPrice:    9050,
				HighestPrice: 10500,
				LowestPrice:  9000,
				ClosePrice:   10200,
				AveragePrice: 10000,
				Volume:       2000000,
				Value:        1000000,
			},
		},
	}

	for _, tt := range tests {
		u := ohlcUsecase{
			ohlcDatalogic: ohlcDatalogicMock,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, err := u.GetSummary(ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSummary failed got : %v, err %v \n", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummary failed want err : %v, err %v \n", tt.wantErr, err)
			}
		})
	}
}
