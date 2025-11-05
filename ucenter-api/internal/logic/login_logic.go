// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"grpc-common/ucenter/types/login"
	"mscoin-common/tools"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	loginReq := &login.LoginReq{}
	if err := copier.Copy(loginReq, req); err != nil {
		return nil, err
	}
	loginResp, err := l.svcCtx.UCLoginRpc.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}
	if err := copier.Copy(resp, loginResp); err != nil {
		return nil, err
	}

	return &types.LoginResp{}, nil
}

func (l *LoginLogic) CheckLogin(token string) (bool, error) {
	_, err := tools.ParseToken(token, l.svcCtx.Config.JWT.AccessSecret)
	if err != nil {
		logx.Error(err)
		return false, err
	}

	return true, nil
}
