package impl

import (
	"context"
	"reflect"
	"testing"

	"part1/internal/model"
	redisRepo "part1/internal/repository/redis"
)

func Test_CalculateOhlc(t *testing.T) {
	redisRepoMock := new(redisRepo.MockRepository)

	type args struct {
		prevExist bool
		prevOhlc  model.Ohlc
		req       model.Transaction
	}

	prevOhlcPrevPrice := model.Ohlc{
		StockCode:     "BBCA",
		PreviousPrice: 8000,
	}

	prevOhlcFull := model.Ohlc{
		StockCode:       "BBCA",
		PreviousPrice:   8000,
		OpenPrice:       8050,
		HighestPrice:    8100,
		LowestPrice:     7900,
		ClosePrice:      8100,
		Volume:          100,
		Value:           810000,
		AveragePrice:    8100,
		LastOrderNumber: 123123,
	}

	reqTxQuantity0 := model.Transaction{
		StockCode: "BBCA",
		Type:      "A",
		Quantity:  "",
		Price:     "8000",
	}

	reqTxTypeE := model.Transaction{
		StockCode:        "BBCA",
		Type:             "E",
		ExecutedQuantity: "100",
		ExecutionPrice:   "8500",
		OrderNumber:      "123122",
	}

	reqTxTypePLowest := model.Transaction{
		StockCode:        "BBCA",
		Type:             "E",
		ExecutedQuantity: "100",
		ExecutionPrice:   "7800",
		OrderNumber:      "123122",
	}

	reqTxTypePClose := model.Transaction{
		StockCode:        "BBCA",
		Type:             "E",
		ExecutedQuantity: "100",
		ExecutionPrice:   "8000",
		OrderNumber:      "123125",
	}

	tests := []struct {
		name    string
		args    args
		want    model.Ohlc
		wantErr bool
	}{
		{
			name: "error parse int",
			args: args{
				prevExist: false,
				prevOhlc:  model.Ohlc{},
				req: model.Transaction{
					StockCode: "BBCA",
					Type:      "A",
					Quantity:  "",
					Price:     "asd",
				},
			},
			want:    model.Ohlc{},
			wantErr: true,
		},
		{
			name: "prev data not exist and quantity = 0",
			args: args{
				prevExist: false,
				prevOhlc:  model.Ohlc{},
				req:       reqTxQuantity0,
			},
			want: model.Ohlc{
				StockCode:     "BBCA",
				PreviousPrice: 8000,
			},
		},
		{
			name: "prev data exist and type = A",
			args: args{
				prevExist: true,
				prevOhlc:  prevOhlcPrevPrice,
				req:       reqTxQuantity0,
			},
			want: model.Ohlc{
				StockCode:     "BBCA",
				PreviousPrice: 8000,
			},
		},
		{
			name: "prev data exist only prev price and type = E",
			args: args{
				prevExist: true,
				prevOhlc:  prevOhlcPrevPrice,
				req:       reqTxTypeE,
			},
			want: model.Ohlc{
				StockCode:       "BBCA",
				PreviousPrice:   8000,
				OpenPrice:       8500,
				HighestPrice:    8500,
				LowestPrice:     8500,
				Volume:          100,
				Value:           850000,
				ClosePrice:      8500,
				AveragePrice:    8500,
				LastOrderNumber: 123122,
			},
		},
		{
			name: "prev data exist full and type = E set highest price",
			args: args{
				prevExist: true,
				prevOhlc:  prevOhlcFull,
				req:       reqTxTypeE,
			},
			want: model.Ohlc{
				StockCode:       "BBCA",
				PreviousPrice:   8000,
				OpenPrice:       8050,
				HighestPrice:    8500,
				LowestPrice:     7900,
				ClosePrice:      8100,
				Volume:          200,
				Value:           1660000,
				AveragePrice:    8300,
				LastOrderNumber: 123123,
			},
		},
		{
			name: "prev data exist full and type = E set lowest price",
			args: args{
				prevExist: true,
				prevOhlc:  prevOhlcFull,
				req:       reqTxTypePLowest,
			},
			want: model.Ohlc{
				StockCode:       "BBCA",
				PreviousPrice:   8000,
				OpenPrice:       8050,
				HighestPrice:    8100,
				LowestPrice:     7800,
				ClosePrice:      8100,
				Volume:          200,
				Value:           1590000,
				AveragePrice:    7950,
				LastOrderNumber: 123123,
			},
		},
		{
			name: "prev data exist full and type = E set close price",
			args: args{
				prevExist: true,
				prevOhlc:  prevOhlcFull,
				req:       reqTxTypePClose,
			},
			want: model.Ohlc{
				StockCode:       "BBCA",
				PreviousPrice:   8000,
				OpenPrice:       8050,
				HighestPrice:    8100,
				LowestPrice:     7900,
				ClosePrice:      8000,
				Volume:          200,
				Value:           1610000,
				AveragePrice:    8050,
				LastOrderNumber: 123125,
			},
		},
	}

	for _, tt := range tests {
		d := datalogic{
			redisRepo: redisRepoMock,
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.CalculateOhlc(context.Background(), tt.args.prevExist, tt.args.prevOhlc, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateOhlc test failed. want: %+v, got: %+v", tt.want, got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateOhlc test failed. want err: %+v, got: %+v", tt.wantErr, err)
			}
		})
	}
}
