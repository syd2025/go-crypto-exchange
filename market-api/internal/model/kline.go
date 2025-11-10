package model

import (
	"grpc-common/market/types/market"
	"mscoin-common/op"
	"mscoin-common/tools"
)

type Kline struct {
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Volume       float64 `bson:"volume,omitempty"`
	Count        float64 `bson:"count,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Turnover     float64 `bson:"turnover,omitempty"`
	TimeStr      string  `bson:"timeStr,omitempty"`
}

func (k *Kline) ToCoinThumb(symbol string, ct *market.CoinThumb) *market.CoinThumb {
	isSame := false
	if ct.Symbol == symbol && ct.DataTime == k.Time {
		isSame = true
	}
	if !isSame {
		ct.Close = k.ClosePrice
		ct.Open = k.OpenPrice
		if ct.High < k.HighestPrice {
			ct.High = k.HighestPrice
		}
		ct.Low = k.LowestPrice
		if ct.Low > k.LowestPrice {
			ct.Low = k.LowestPrice
		}

		ct.Zone = 0
		ct.Volume = op.AddN(k.Volume, ct.Volume, 8)
		ct.Turnover = op.AddN(k.Turnover, ct.Turnover, 8)
		ct.Change = k.LowestPrice - ct.Close
		ct.Chg = op.MulN(op.DivN(ct.Change, ct.Close, 5), 100, 3)
		ct.UsdRate = k.ClosePrice
		ct.BaseUsdRate = 1
		ct.Trend = append(ct.Trend, k.ClosePrice)
		ct.DataTime = k.Time
	}

	return ct
}

// 表声明，分表
func (k *Kline) TableName(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

func (k *Kline) InitCoinThumb(symbol string) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice
	ct.High = k.HighestPrice
	ct.Volume = k.Volume
	ct.Turnover = k.Turnover
	ct.Low = k.LowestPrice
	ct.Zone = 0
	ct.Change = 0
	ct.Chg = 0
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	ct.DataTime = k.Time
	ct.Trend = make([]float64, 0)
	return ct
}

func NewKline(data []string, period string) *Kline {
	toInt64 := tools.ToInt64(data[0])
	return &Kline{
		Time:         toInt64,
		Period:       period,
		OpenPrice:    tools.ToFloat64(data[1]),
		HighestPrice: tools.ToFloat64(data[2]),
		LowestPrice:  tools.ToFloat64(data[3]),
		ClosePrice:   tools.ToFloat64(data[4]),
		Count:        tools.ToFloat64(data[5]),
		Volume:       tools.ToFloat64(data[6]),
		Turnover:     tools.ToFloat64(data[7]),
		TimeStr:      tools.ToTimeString(toInt64),
	}

}
