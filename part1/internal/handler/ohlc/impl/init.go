package impl

import (
	ohlcHandler "part1/internal/handler/ohlc"
)

type handler struct {
}

func New() ohlcHandler.OhlcHandler {
	return &handler{}
}
