package processor

import (
	"market-api/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebsocketHandler struct {
}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

func (w *WebsocketHandler) HandleTrade(symbol string, data []byte) {

}

func (w *WebsocketHandler) HandleKLine(symbol string, kline *model.Kline) {
	logx.Info("=========================WebsocketHandler=========================")
	logx.Info("symbol: ", symbol)
	logx.Info("close: ", kline.ClosePrice, "high: ", kline.HighestPrice, "low: ", kline.LowestPrice, "open: ", kline.OpenPrice, "volume: ", kline.Volume)
	logx.Info("=================================end==============================")
}
