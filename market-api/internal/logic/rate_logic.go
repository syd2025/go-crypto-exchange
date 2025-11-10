// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"grpc-common/market/types/rate"
<<<<<<< HEAD
	"market/internal/domain"
=======
	"time"
>>>>>>> origin/main

	"market-api/internal/svc"
	"market-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExchangeRateLogic struct {
	logx.Logger
<<<<<<< HEAD
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	exchangeRateDomain *domain.ExchangeRateDomain
=======
	ctx    context.Context
	svcCtx *svc.ServiceContext
	// exchangeRateDomain *domain.ExchangeRateDomain
>>>>>>> origin/main
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
<<<<<<< HEAD
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
=======
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		// exchangeRateDomain: domain.NewExchangeRateDomain(),
	}
}

func (l *ExchangeRateLogic) UsdRate(req *types.RateRequest) (*types.RateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usdRate, err := l.svcCtx.ExchangeRateRpc.UsdRate(ctx, &rate.RateReq{
		Unit: req.Unit,
		Ip:   req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return &types.RateResponse{
		Rate: usdRate.Rate,
>>>>>>> origin/main
	}, nil
}
