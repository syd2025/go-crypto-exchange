package svc

import (
	"common/msdb"
	"ucenter/internal/config"
	"ucenter/internal/database"

	"github.com/zeromicro/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("mscoin"), nil)
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     database.InitGorm(c.Mysql.DataSource),
	}
}
