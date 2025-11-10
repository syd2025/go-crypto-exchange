// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"grpc-common/market/mclient"
	"market-api/internal/config"
<<<<<<< HEAD
=======
	"market-api/internal/database"
	"market-api/internal/processor"
	"market-api/internal/ws"
>>>>>>> origin/main

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
<<<<<<< HEAD
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
=======
	MarketRpc       mclient.Market
	Processor       processor.Processor
}

func NewServiceContext(c config.Config, wsServer *ws.WebsocketServer) *ServiceContext {
	// 初始化processor
	kafkaCli := database.NewKafkaClient(c.Kafka)
	market := mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc))
	defaultProcessor := processor.NewDefaultProcessor(kafkaCli)
	defaultProcessor.Init(market)
	defaultProcessor.AddHandler(processor.NewWebsocketHandler(wsServer))

	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
		MarketRpc:       mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
		Processor:       defaultProcessor,
>>>>>>> origin/main
	}
}
