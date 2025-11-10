package logic

import (
	"context"
	"grpc-common/market/types/market"

	"market/internal/domain"

	"market/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeCoinDomain *domain.ExchangeCoinDomain
	marketDomain       *domain.MarketDomain
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeCoinDomain: domain.NewExchangeCoinDomain(svcCtx.Db),
		marketDomain:       domain.NewMarketDomain(svcCtx.MongoClient),
	}
}

func (l *MarketLogic) FindSymbolThumbTrend(req *market.MarketReq) (*market.SymbolThumbRes, error) {
	coins, err := l.exchangeCoinDomain.FindVisible(context.Background())
	if err != nil {
		return nil, err
	}

	// 查询mongodb中一小时的数据
	coinThumbs := l.marketDomain.SymbolThumbTrend(coins)

	return &market.SymbolThumbRes{
		List: coinThumbs,
	}, nil
}
