package module

import (
	"context"
	"part1/internal/model"
)

func (u ohlcUsecase) InitData(ctx context.Context) (resp model.InitDataResponse, err error) {

	//1. read data from file
	err = u.fsDatalogic.GetDataFromJSON(ctx)
	if err != nil {
		return model.InitDataResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}

	return model.InitDataResponse{
		Success: true,
	}, err
}
