// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"grpc-common/market/types/rate"
	"market/internal/domain"

	"market-api/internal/svc"
	"market-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExchangeRateLogic struct {
	logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	exchangeRateDomain *domain.ExchangeRateDomain
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		Logger:             logx.WithContext(ctx),
		ctx:                ctx,
		svcCtx:             svcCtx,
		exchangeRateDomain: domain.NewExchangeRateDomain(),
	}
}

func (l *ExchangeRateLogic) UsdRate(req types.RateRequest) (*types.RateResponse, error) {
	usdRate := l.exchangeRateDomain.UsdRate(req.Unit)
	return &rate.RateResp{
		Rate: usdRate,
	}, nil
}
