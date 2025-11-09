package processor

import "market-api/internal/model"

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKLine(symbol string, kline *model.Kline)
}
