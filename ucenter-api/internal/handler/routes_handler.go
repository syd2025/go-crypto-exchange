package handler

import (
	"ucenter-api/internal/svc"
)

// 路有设置
func RegisterHandlers(r *Routes, serverCtx *svc.ServiceContext) {
	register := NewRegisterHandler(serverCtx)
	registerGroup := r.Group()
	registerGroup.Post("/uc/register/phone", register.Register)
	registerGroup.Post("/uc/mobile/code", register.SendCode)

	login := NewLoginHandler(serverCtx)
	loginGroup := r.Group()
	loginGroup.Post("/uc/register/phone", login.Login)
	loginGroup.Post("/uc/check/login", login.CheckLogin)
}
