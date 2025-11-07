package config

import (
	"market/internal/database"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mongo      database.MongoConfig
	Mysql      database.MysqlConfig
	CacheRedis cache.CacheConf
}
