package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis cache.CacheConf
	Mysql      MysqlConfig
	Captcha    CaptchaConfig
}

type CaptchaConfig struct {
	Vid string
	Key string
}

type MysqlConfig struct {
	DataSource string
}
