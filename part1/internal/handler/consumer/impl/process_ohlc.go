package impl

import (
	"context"
	"encoding/json"
	"part1/internal/model"
)

func (h *handler) ProcessOhlc(in []byte) {

	req := model.Transaction{}
	err := json.Unmarshal(in, &req)
	if err != nil {
		return
	}

	h.ohlcUsecase.ProcessOhlc(context.Background(), req)
}
