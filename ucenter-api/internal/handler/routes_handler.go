package handler

import (
	"ucenter-api/internal/svc"
)

// 路有设置
func RegisterHandlers(r *Routes, serverCtx *svc.ServiceContext) {
	register := NewRegisterHandler(serverCtx)

	registerGroup := r.Group()
	registerGroup.Get("/uc/register/phone", register.Register)
}
