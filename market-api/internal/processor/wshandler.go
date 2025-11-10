package processor

import (
	"encoding/json"
	"grpc-common/market/types/market"
	"market-api/internal/model"
	"market-api/internal/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebsocketHandler struct {
	wsServer *ws.WebsocketServer
}

func NewWebsocketHandler(wsServer *ws.WebsocketServer) *WebsocketHandler {
	return &WebsocketHandler{
		wsServer: wsServer,
	}
}

func (w *WebsocketHandler) HandleTrade(symbol string, data []byte) {

}

func (w *WebsocketHandler) HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb) {
	logx.Info("=========================WebsocketHandler=========================")
	logx.Info("symbol: ", symbol)
	thumb := thumbMap[symbol]
	if thumb == nil {
		kline.InitCoinThumb(symbol)
	}
	coinThumb := kline.ToCoinThumb(symbol, thumb)
	marshal, _ := json.Marshal(coinThumb)

	logx.Info("close: ", kline.ClosePrice, "high: ", kline.HighestPrice, "low: ", kline.LowestPrice, "open: ", kline.OpenPrice, "volume: ", kline.Volume)
	w.wsServer.BroadcastToNamespace("/", "/topic/market/thumb", string(marshal))
	logx.Info("=================================end==============================")
}
