package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      MysqlConfig
	CacheRedis cache.CacheConf
	Captcha    CaptchaConfig
	Jwt        AuthConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}

type CaptchaConfig struct {
	Vid string
	Key string
}

type MysqlConfig struct {
	DataSource string
}
