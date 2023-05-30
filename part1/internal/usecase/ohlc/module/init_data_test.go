package module

import (
	"context"
	"errors"
	"reflect"
	"testing"

	fsDatalogic "part1/internal/datalogic/fs"
	"part1/internal/model"
)

func Test_InitData(t *testing.T) {
	fsDatalogicMock := new(fsDatalogic.MockDatalogic)

	ctx := context.Background()

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
		want    model.InitDataResponse
	}{
		{
			name: "failed get data json",
			mock: func() {
				fsDatalogicMock.On("GetDataFromJSON", ctx).
					Return(errors.New("error get data")).
					Once()
			},
			want: model.InitDataResponse{
				Success:      false,
				ErrorMessage: "error get data",
			},
			wantErr: true,
		},
		{
			name: "success",
			mock: func() {
				fsDatalogicMock.On("GetDataFromJSON", ctx).
					Return(nil).
					Once()
			},
			want: model.InitDataResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		u := ohlcUsecase{
			fsDatalogic: fsDatalogicMock,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, err := u.InitData(ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitData failed got : %v, err %v \n", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("InitData failed want err : %v, err %v \n", tt.wantErr, err)
			}
		})
	}
}
