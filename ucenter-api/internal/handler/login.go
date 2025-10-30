package handler

import (
	"context"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

func NewLoginHandler(svcCtx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{
		svcCtx: svcCtx,
	}
}

func (h *LoginHandler) Login(ctx context.Context, req types.LoginReq) (resp *types.LoginResp, err error) {
	return &types.LoginResp{
		Id:     1,
		Name:   "sonny",
		Gender: "male",
		Age:    20,
	}, nil
}
