package domain

import (
	"encoding/json"
	"mscoin-common/tools"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaDomain struct {
}

type VaptchaReq struct {
	Id        string `json:"id"`
	SecretKey string `json:"secretKey"`
	Scene     int    `json:"scene"`
	Ip        string `json:"ip"`
	Token     string `json:"token"`
}

type VaptchaResp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}

func (c *CaptchaDomain) Verify(server, token, vid, key string, scene int, ip string) bool {
	resp, err := tools.Post(server, &VaptchaReq{
		Id:        vid,
		SecretKey: key,
		Scene:     scene,
		Ip:        ip,
		Token:     token,
	})
	if err != nil {
		logx.Error("人机验证失败:" + err.Error())
		return false
	}
	result := &VaptchaResp{}
	json.Unmarshal(resp, result) // 反解析成对象
	return result.Success == 1
}
