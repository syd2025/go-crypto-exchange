package logic

import (
	"context"
<<<<<<< HEAD
	"market-api/internal/types"
=======
	"grpc-common/market/types/rate"
>>>>>>> origin/main

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

<<<<<<< HEAD
func (l *ExchangeRateLogic) UsdRate(req types.RateRequest) (*types.RateResponse, error) {

	return &types.RateResponse{}, nil
=======
func (l *ExchangeRateLogic) UsdRate(req *rate.RateReq) (*rate.RateResp, error) {
	usdRate := l.exchangeRateDomain.UsdRate(req.Unit)
	return &rate.RateResp{
		Rate: usdRate,
	}, nil
>>>>>>> origin/main
}
