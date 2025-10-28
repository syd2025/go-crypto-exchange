package main

import (
	"flag"
	"fmt"
	"ucenter-api/internal/config"
	"ucenter-api/internal/handler"
	"ucenter-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{
		Stat:     false,
		Encoding: "plain",
	})
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 加载配置文件
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 创建服务上下文
	ctx := svc.NewServiceContext(c)
	// 注册路由和处理器
	router := handler.NewRoutes(server)
	handler.RegisterHandlers(router, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
