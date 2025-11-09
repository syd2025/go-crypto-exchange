package handler

import (
	"context"
	"market-api/internal/logic"
	"market-api/internal/svc"

	"market-api/internal/types"

	common "mscoin-common"
	"mscoin-common/tools"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type ExchangeRateHandler struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExchangeRateHandler(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (h *ExchangeRateHandler) UsdRate(w http.ResponseWriter, r *http.Request) {
	var req types.RateRequest
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	// 获取一下IP
	req.Ip = tools.GetRemoteClientIp(r)

	l := logic.NewExchangeRateLogic(r.Context(), h.svcCtx)
	resp, err := l.UsdRate(&req)
	result := common.NewResult().Deal(resp.Rate, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
