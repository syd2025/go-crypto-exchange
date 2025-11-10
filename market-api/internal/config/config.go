// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
<<<<<<< HEAD
=======
	"market-api/internal/database"

>>>>>>> origin/main
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	MarketRpc zrpc.RpcClientConf
	Prefix    string
<<<<<<< HEAD
=======
	Kafka     database.KafkaConfig
>>>>>>> origin/main
}
