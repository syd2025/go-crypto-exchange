package processor

import (
	"grpc-common/market/types/market"
	"market-api/internal/model"
)

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKLine(symbol string, kline *model.Kline, thumpMap map[string]*market.CoinThumb)
}
