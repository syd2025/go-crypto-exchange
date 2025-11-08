package types

type RateRequest struct {
	Unit string `path:"unit" json:"unit"`
	Ip   string `json:"ip,omitempty"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}

type MarketReq struct {
	Ip string `json:"ip,optional"`
}

type CoinThumbResp struct {
	Symbol       string    `json:"symbol"`
	Open         string    `json:"open"`
	High         string    `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
	Chg          float64   `json:"chg"` // 变化百分比
	Change       float64   `json:"change"`
	Volume       float64   `json:"volume"`
	Turnover     float64   `json:"turnover"`
	LastDayClose float64   `json:"last_day_close"`
	UsdRate      float64   `json:"usd_rate"`      // USDT汇率
	BaseUsdRate  float64   `json:"base_usd_rate"` // 基础USDT汇率
	Zone         int       `json:"zone"`
	Trend        []float64 `json:"trend, optional"` // 价格趋势
}
