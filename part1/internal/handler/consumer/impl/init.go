package impl

import (
	consumerHandler "part1/internal/handler/consumer"

	ohlcUsecase "part1/internal/usecase/ohlc"
)

type handler struct {
	ohlcUsecase ohlcUsecase.OhlcUsecase
}

func New(ohlcUsecase ohlcUsecase.OhlcUsecase) consumerHandler.ConsumerHandler {
	return &handler{ohlcUsecase: ohlcUsecase}
}
