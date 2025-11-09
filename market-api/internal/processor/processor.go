package processor

import (
	"encoding/json"
	"market-api/internal/database"
	"market-api/internal/model"
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
}

type DefaultProcessor struct {
	client   *database.KafkaClient
	handlers []MarketHandler
}

func NewDefaultProcessor(client *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		client:   client,
		handlers: make([]MarketHandler, 0),
	}
}

func (d *DefaultProcessor) Init() {
	d.startReadFromKafka(KLINE1M, KLINE)
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
			v.HandleKLine(symbol, kline)
		}
	}
}

func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	// 发送到websocket服务
	d.handlers = append(d.handlers, h)
}
