// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"grpc-common/market/mclient"
	"market-api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
	}
}
