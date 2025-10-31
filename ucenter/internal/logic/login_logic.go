package logic

import (
	"context"
	"grpc-common/ucenter/types/login"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginResp, error) {
	// 1.检验验证码
	isVerify := l.NewCaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip
	)
	if !isVerify {
		return nil, errors.New("人机校验失败")
	}
	// 2.校验密码
	member,err :=l.MemberDomain.FindByPhone(context.Background(), in.GetUsername())
	if err != nil {
		logx.Error("登录失败")
		return nil, errors.New("服务异常，请联系管理员")
	}
	if member == nil {
		logx.Error("登录失败")
		return nil, errors.New("用户不存在")
	}

	password := member.Password
	salt := member.Salt
	verfiy := tools.Verify(in.Password, salt, password)
	if !verfiy {
		logx.Error("登录失败")
		return nil, errors.New("密码错误")
	}
	// 登录成功
	key := l.svcCtx.Config.Jwt.AccessSecret
	expire := l.svcCtx.Config.Jwt.AccessExpire

	token, err:= getJwtToken(key, time.Now().Unix(), expire, member.Id)
	if err != nil {
		logx.Error("登录失败")
		return nil, errors.New("登录失败")
	}
	/// 返回登录信息
	loginCount := member.LoginCount + 1
	go func(){
		l.MemberDomain.UpdateLoginCount(context.Background(), member.Id, 1)
	}()
	return &login.LoginResp{
		Token: token,
		Id: member.Id,
		Username: member.Username,
		MemberLevel: member.MemberLevelStr(),
		MemberRate: member.MemberRate(),
		RealName: member.RealName,
		Country: member.Country,
		Avatar: member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner: member.SuperPartner,
		LoginCount: int32(loginCount),
	}, nil
}

// 生成token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}