package impl

import (
	ohlcHandler "part1/internal/handler/ohlc"

	ohlcUsecase "part1/internal/usecase/ohlc"
)

type handler struct {
	ohlcUsecase ohlcUsecase.OhlcUsecase
}

func New(ohlcUsecase ohlcUsecase.OhlcUsecase) ohlcHandler.OhlcHandler {
	return &handler{ohlcUsecase: ohlcUsecase}
}
