// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"

	"market-api/internal/svc"
	"market-api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	// exchangeRateDomain *domain.ExchangeRateDomain
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		// exchangeRateDomain: domain.NewExchangeRateDomain(),
	}
}

func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	var thumbs []*market.CoinThumb
	thumb := l.svcCtx.Processor.GetThumb()
	isCache := false
	if thumb != nil {
		switch thumb.(type) {
		case []*market.CoinThumb:
			thumbs = thumb.([]*market.CoinThumb)
			isCache = true
		}
	}

	if !isCache {
		symbolThumbRes, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(context.Background(), &mclient.MarketReq{
			Ip: req.Ip,
		})
		if err != nil {
			return nil, err
		}
		thumbs = symbolThumbRes.List
	}

	if err := copier.Copy(&list, thumbs); err != nil {
		return nil, err
	}
	return
}
