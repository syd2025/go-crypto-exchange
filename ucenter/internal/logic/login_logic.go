package logic

import (
	"context"
	"errors"
	"mscoin-common/tools"
	"time"

	"grpc-common/ucenter/types/login"

	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	memberDomain  *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		memberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginResp, error) {
	isVerify := l.CaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip,
	)
	if !isVerify {
		return nil, errors.New("人机验证失败")
	}

	// 2. 校验密码
	member, err := l.memberDomain.FindByPhone(context.Background(), in.Username)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败，请稍后重试")
	}

	if member == nil {
		return nil, errors.New("用户不存在")
	}

	password := member.Password
	salt := member.Salt
	verify := tools.Verify(in.Password, salt, password, nil)
	if !verify {
		return nil, errors.New("密码不正确")
	}

	// 3. 登录成功，生成token  JWT
	key := l.svcCtx.Config.JWT.AccessSecret
	expire := l.svcCtx.Config.JWT.AccessExpire
	token, err := l.getJwtToken(key, time.Now().Unix(), expire, member.Id)
	if err != nil {
		return nil, errors.New("登录失败，请稍后重试")
	}

	loginCount := member.LoginCount + 1
	go func() {
		l.memberDomain.UpdateLoginCount(context.Background(), member.Id, 1)
	}()
	return &login.LoginResp{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperParner:   member.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
