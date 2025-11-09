// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"grpc-common/market/mclient"
	"market-api/internal/config"
	"market-api/internal/database"
	"market-api/internal/processor"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
	MarketRpc       mclient.Market
	Processor       processor.Processor
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化processor
	kafkaCli := database.NewKafkaClient(c.Kafka)
	defaultProcessor := processor.NewDefaultProcessor(kafkaCli)
	defaultProcessor.Init()
	defaultProcessor.AddHandler(processor.NewWebsocketHandler())
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
		MarketRpc:       mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
		Processor:       defaultProcessor,
	}
}
