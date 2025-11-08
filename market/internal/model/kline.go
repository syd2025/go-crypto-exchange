package model

import (
	"grpc-common/market/types/market"
	"mscoin-common/op"
)

type Kline struct {
	Period       string  `bson:"period, omitempty"`
	OpenPrice    float64 `bson:"open_price, omitempty"`
	HighestPrice float64 `bson:"high_price, omitempty"`
	LowestPrice  float64 `bson:"low_price, omitempty"`
	ClosePrice   float64 `bson:"close_price, omitempty"`
	Volume       float64 `bson:"volume, omitempty"`
	Time         int64   `bson:"time, omitempty"`
	Count        float64 `bson:"count, omitempty"`
	Turnover     float64 `bson:"turnover, omitempty"`
}

func (Kline) Table(symbol string, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

func (k *Kline) ToCoinThumb(symbol string, end *Kline) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice
	ct.Zone = 0
	ct.Change = k.ClosePrice - end.ClosePrice
	ct.Chg = op.DivN(ct.Change, end.LowestPrice, 5)
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	return ct
}

func DefaultCoinThumb(symbol string) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Trend = []float64{}
	return ct
}
