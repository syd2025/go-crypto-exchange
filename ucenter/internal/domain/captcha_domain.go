package domain

import (
	"common/tools"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type vaptchaReq struct {
	Id        string `json:"id"`
	SecretKey string `json:"secretKey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}

type vaptchaRes struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

type CaptchaDomain struct {
}

func (c *CaptchaDomain) Verify(
	server string,
	vid string,
	key string,
	token string,
	scene int,
	ip string) bool {

	resp, err := tools.Post(server, &vaptchaReq{
		Id:        vid,
		SecretKey: key,
		Scene:     scene,
		Token:     token,
		Ip:        ip,
	})
	if err != nil {
		logx.Error("调用人机校验平台失败: %s", err.Error())
		return false
	}
	result := &vaptchaRes{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		logx.Error("解析人机校验平台返回结果失败: %s", err.Error())
		return false
	}
	return result.Success == 1
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}
