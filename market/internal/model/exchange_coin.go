package model

type ExchangeCoin struct {
	Id               int64   `gorm:"column:id"`
	Symbol           string  `gorm:"column:symbol"`
	BaseCoinScale    int64   `gorm:"column:base_coin_scale"`
	BaseSymScale     int64   `gorm:"column:base_sym_scale"`
	CoinSymbol       int64   `gorm:"column:coin_symbol"`
	Enable           int64   `gorm:"column:enable"`
	Fee              float64 `gorm:"column:fee"`
	Sort             int64   `gorm:"column:sort"`
	EnableMarketBuy  int64   `gorm:"column:enable_market_buy"`
	EnableMarketSell int64   `gorm:"column:enable_market_sell"`
	MinSellPrice     float64 `gorm:"column:min_sell_price"`
	Flag             int64   `gorm:"column:flag"`
	MaxTradingOrder  int64   `gorm:"column:max_trading_order"`
	MaxTradingTime   int64   `gorm:"column:max_trading_time"`
	MinTurnover      float64 `gorm:"column:min_turnover"`
	ClearTime        int64   `gorm:"column:clear_time"`
	EndTime          int64   `gorm:"column:end_time"`
	Exchangeable     int64   `gorm:"column:exchangeable"`
	MaxBuyPrice      float64 `gorm:"column:max_buy_price"`
	MaxVolume        float64 `gorm:"column:max_volume"`
	MinVolume        float64 `gorm:"column:min_volume"`
	PublishAmount    float64 `gorm:"column:publish_amount"`
	PublishPrice     float64 `gorm:"column:publish_price"`
	PublishType      int64   `gorm:"column:publish_type"`
	RobotType        int64   `gorm:"column:robot_type"`
	StartTime        int64   `gorm:"column:start_time"`
	Visible          int64   `gorm:"column:visible"`
	Zone             int64   `gorm:"column:zone"`
}

func (ExchangeCoin) TableName() string {

	return "exchange_coin"

}
