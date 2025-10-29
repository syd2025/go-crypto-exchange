package logic

import (
	"common/tools"
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER::"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// 1.检验人机校验
	isVerify := l.CaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip)
	if !isVerify {
		return nil, errors.New("人机校验失败")
	}
	logx.Info("人机校验通过....")
	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(in *register.CodeReq) (*register.NoRes, error) {
	// 受到手机号和国家标识
	// 生成验证码
	// 根据对应的国家和手机号调用对应的短信平台发送验证码
	// 将验证码存入redis，过其时间为5分钟
	// 返回成功
	code := tools.Gen4Code()
	// 发送验证码
	go func() {
		logx.Info("调用短信平台发送验证码成功")
	}()

	logx.Infof("验证码为: %s\n", code)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+in.Phone, code, 15*time.Minute)
	if err != nil {
		return nil, errors.New("验证码发送失败")
	}

	return &register.NoRes{}, nil
}
