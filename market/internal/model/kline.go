package model

type Kline struct {
	Period     string  `bson:"period, omitempty"`
	OpenPrice  float64 `bson:"open_price, omitempty"`
	HighPrice  float64 `bson:"high_price, omitempty"`
	LowPrice   float64 `bson:"low_price, omitempty"`
	ClosePrice float64 `bson:"close_price, omitempty"`
	Volume     float64 `bson:"volume, omitempty"`
	Time       int64   `bson:"time, omitempty"`
	Count      float64 `bson:"count, omitempty"`
	Turnover   float64 `bson:"turnover, omitempty"`
}

func (Kline) Table(symbol string, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}
