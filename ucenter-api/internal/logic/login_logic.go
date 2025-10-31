package logic

import (
	"common/tools"
	"context"
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

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	loginReq := &Login.LoginReq{}
	if err = copier.Copy(loginReq, req); err != nil {
		return nil, err
	}
	loginResp, err = l.svcCtx.UCLoginRpc.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{}
	if err := copier.Copy(resp, loginResp); err != nil {
		return nil, err
	}
	return
}

func (l *LoginLogic) CheckLogin(token string) (bool, error) {
	userId, err := tools.ParseToken(token, l.svcCtx.Config.Jwt.Secret)
	if err != nil {
		return false, err
	}
	return true, nil
}
