package handler

import (
	"common"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 注册处理器
type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

// 注册用户
func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, _ := l.Register(&req)
	result := common.NewResult().Deal(resp, nil)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// 发送验证码
func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, _ := l.SendCode(&req)
	result := common.NewResult().Deal(resp, nil)
	httpx.OkJsonCtx(r.Context(), w, result)

}
