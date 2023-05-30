package model

type InitDataResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message"`
}

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

type Ohlc struct {
	StockCode       string `json:"stock_code" redis:"stock_code"`
	PreviousPrice   int64  `json:"previous_price" redis:"previous_price"`
	OpenPrice       int64  `json:"open_price" redis:"open_price"`
	HighestPrice    int64  `json:"highest_price" redis:"highest_price"`
	LowestPrice     int64  `json:"lowest_price" redis:"lowest_price"`
	ClosePrice      int64  `json:"close_price" redis:"close_price"`
	AveragePrice    int64  `json:"average_price" redis:"average_price"`
	Volume          int64  `json:"volume" redis:"volume"`
	Value           int64  `json:"value" redis:"value"`
	LastOrderNumber int64  `json:"last_order_number" redis:"last_order_number"`
}
