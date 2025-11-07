package logic

import (
	"context"
	"market-api/internal/types"

	"market/internal/domain"

	"market/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExchangeRateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeRateDomain *domain.ExchangeRateDomain
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeRateDomain: domain.NewExchangeRateDomain(),
	}
}

func (l *ExchangeRateLogic) UsdRate(req types.RateRequest) (*types.RateResponse, error) {

	return &types.RateResponse{}, nil
}
