package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routes, serverCtx *svc.ServiceContext) {
	regiser := NewRegisterHandler(serverCtx)

	registerGroup := r.Group()
	registerGroup.Get("/uc/register/phone", regiser.Register)

}
