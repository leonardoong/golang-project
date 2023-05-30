package main

import (
	"encoding/json"
	"fmt"

	m "part2/model"
)

// refactor code from O(n * m * i) = O(n^3) time complexity
// to n + m + i which is O(n) time complexity
func ohlc(stockCodes []string, records []m.ChangeRecord, idxMember []m.IndexMember) map[string]m.Summary {
	var result = map[string]m.Summary{}

	for _, stockCode := range stockCodes {
		// create barebone of Summary
		result[stockCode] = m.Summary{
			StockCode: stockCode,
		}
	}
	for _, rec := range records {
		if _, found := result[rec.StockCode]; !found {
			result[rec.StockCode] = m.Summary{}
		}

		tempSum := result[rec.StockCode]

		if rec.Quantity == 0 {
			tempSum.Prev = rec.Price
		} else if rec.Quantity > 0 && tempSum.Open == 0 {
			tempSum.Open = rec.Price
		}

		tempSum.Close = rec.Price
		if tempSum.High < rec.Price {
			tempSum.High = rec.Price
		}

		if tempSum.Low == 0 || tempSum.Low > rec.Price {
			tempSum.Low = rec.Price
		}

		result[rec.StockCode] = tempSum
	}

	for _, idxCode := range idxMember {
		if _, found := result[idxCode.StockCode]; !found {
			result[idxCode.StockCode] = m.Summary{}
		}

		tempSum := result[idxCode.StockCode]
		tempSum.IndexCode = append(tempSum.IndexCode, idxCode.IndexCode)

		result[idxCode.StockCode] = tempSum
	}

	return result
}

func main() {
	r := ohlc(x, w, p)
	for _, v := range r {
		jss, _ := json.Marshal(v)
		fmt.Println("summary: ", string(jss))
	}
}

var (
	x = []string{
		"BBCA", "BBRI", "ASII", "GOTO",
	}
	w = []m.ChangeRecord{
		{
			StockCode: "BBCA",
			Price:     8783,
			Quantity:  0,
		},
		{
			StockCode: "BBRI",
			Price:     3233,
			Quantity:  0,
		},
		{
			StockCode: "ASII",
			Price:     1223,
			Quantity:  0,
		},
		{
			StockCode: "GOTO",
			Price:     321,
			Quantity:  0,
		},

		{
			StockCode: "BBCA",
			Price:     8780,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3230,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1220,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     320,
			Quantity:  1,
		},

		{
			StockCode: "BBCA",
			Price:     8800,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3300,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1300,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     330,
			Quantity:  1,
		},

		{
			StockCode: "BBCA",
			Price:     8600,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3100,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1100,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     310,
			Quantity:  1,
		},

		{
			StockCode: "BBCA",
			Price:     8785,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3235,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1225,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     325,
			Quantity:  1,
		},
	}
	p = []m.IndexMember{
		{
			StockCode: "BBCA",
			IndexCode: "IHSG",
		},
		{
			StockCode: "BBRI",
			IndexCode: "IHSG",
		},
		{
			StockCode: "ASII",
			IndexCode: "IHSG",
		},
		{
			StockCode: "GOTO",
			IndexCode: "IHSG",
		},
		{
			StockCode: "BBCA",
			IndexCode: "LQ45",
		},
		{
			StockCode: "BBRI",
			IndexCode: "LQ45",
		},
		{
			StockCode: "ASII",
			IndexCode: "LQ45",
		},
		{
			StockCode: "GOTO",
			IndexCode: "LQ45",
		},
		{
			StockCode: "BBCA",
			IndexCode: "KOMPAS100",
		},
		{
			StockCode: "BBRI",
			IndexCode: "KOMPAS100",
		},
	}
)
