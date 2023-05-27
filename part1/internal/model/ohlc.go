package model

type GetSummaryRequest struct {
	StockName string `json:"stock_name"`
}

type GetSummaryResponse struct {
	StockName    string `json:"stock_name"`
	PrevPrice    int64  `json:"prev_price"`
	OpenPrice    int64  `json:"open_price"`
	HighestPrice int64  `json:"highest_price"`
	LowestPrice  int64  `json:"lowest_price"`
	ClosePrice   int64  `json:"close_price"`
	AveragePrice int64  `json:"average_price"`
	Volume       int64  `json:"volume"`
	Value        int64  `json:"value"`
}
