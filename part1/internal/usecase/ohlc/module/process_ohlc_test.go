package module

import (
	"context"
	"errors"
	"part1/internal/model"
	"testing"

	ohlcDatalogic "part1/internal/datalogic/ohlc"
)

func Test_ProcessOhlc(t *testing.T) {
	ohlcDatalogicMock := new(ohlcDatalogic.MockDatalogic)

	ctx := context.Background()

	type args struct {
		req model.Transaction
	}

	request := model.Transaction{
		Type:             "E",
		Quantity:         "10",
		StockCode:        "BBCA",
		ExecutedQuantity: "100",
		ExecutionPrice:   "10000",
	}

	ohlc := model.Ohlc{
		StockCode: "BBCA",
		OpenPrice: 9000,
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
	}{
		{
			name: "failed get redis",
			args: args{
				req: request,
			},
			mock: func() {
				ohlcDatalogicMock.On("GetOhlcDataFromRedis", ctx, request.StockCode).
					Return(false, model.Ohlc{}, errors.New("error redis")).
					Once()
			},
			wantErr: true,
		},
		{
			name: "failed calculate ohlc",
			args: args{
				req: request,
			},
			mock: func() {
				exist := true
				ohlcDatalogicMock.On("GetOhlcDataFromRedis", ctx, request.StockCode).
					Return(exist, ohlc, nil).
					Once()

				ohlcDatalogicMock.On("CalculateOhlc", ctx, exist, ohlc, request).
					Return(model.Ohlc{}, errors.New("error calculate")).
					Once()
			},
			wantErr: true,
		},
		{
			name: "failed stored redis",
			args: args{
				req: request,
			},
			mock: func() {
				exist := true
				ohlcDatalogicMock.On("GetOhlcDataFromRedis", ctx, request.StockCode).
					Return(exist, ohlc, nil).
					Once()

				ohlcDatalogicMock.On("CalculateOhlc", ctx, exist, ohlc, request).
					Return(summary, nil).
					Once()

				ohlcDatalogicMock.On("StoreOhlcData", ctx, summary).
					Return(errors.New("error store data redis")).
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
				ohlcDatalogicMock.On("GetOhlcDataFromRedis", ctx, request.StockCode).
					Return(exist, ohlc, nil).
					Once()

				ohlcDatalogicMock.On("CalculateOhlc", ctx, exist, ohlc, request).
					Return(summary, nil).
					Once()

				ohlcDatalogicMock.On("StoreOhlcData", ctx, summary).
					Return(nil).
					Once()
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

			err := u.ProcessOhlc(ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessOhlc failed want err : %v, err %v \n", tt.wantErr, err)
			}
		})
	}
}
