package processor

import (
	"context"
	"encoding/json"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"market-api/internal/database"
	"market-api/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"

type ProcessData struct {
	Type string
	Key  []byte
	Data []byte
}

type Processor interface {
	Process(data ProcessData)
	AddHandler(h MarketHandler)
	GetThumb() any
}

type DefaultProcessor struct {
	client   *database.KafkaClient
	handlers []MarketHandler
	thumpMap map[string]*market.CoinThumb
}

func NewDefaultProcessor(client *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		client:   client,
		handlers: make([]MarketHandler, 0),
		thumpMap: make(map[string]*market.CoinThumb),
	}
}

func (d *DefaultProcessor) Init(marketRpc mclient.Market) {
	d.startReadFromKafka(KLINE1M, KLINE)
	d.initThumbMap(marketRpc)

}

func (d *DefaultProcessor) GetThumb() any {
	cs := make([]*market.CoinThumb, len(d.thumpMap))
	i := 0
	for _, v := range d.thumpMap {
		cs[i] = v
		i++
	}
	return cs
}

func (d *DefaultProcessor) initThumbMap(marketRpc mclient.Market) {
	symbolThumbRes, err := marketRpc.FindSymbolThumbTrend(context.Background(), &market.MarketReq{})
	if err != nil {
		logx.Info(err)
	} else {
		coinThumbs := symbolThumbRes.List
		for _, v := range coinThumbs {
			d.thumpMap[v.Symbol] = v
		}
	}
}

func (d *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	d.client.StartRead(topic)
	go d.dealQueueData(d.client, tp)
}

func (d *DefaultProcessor) dealQueueData(client *database.KafkaClient, tp string) {
	for {
		msg := client.Read()
		data := ProcessData{
			Type: tp,
			Key:  msg.Key,
			Data: msg.Data,
		}
		d.Process(data)
	}
}

func (d *DefaultProcessor) Process(data ProcessData) {
	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		json.Unmarshal(data.Data, kline)

		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumpMap)
		}
	}
}

func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	// 发送到websocket服务
	d.handlers = append(d.handlers, h)
}
