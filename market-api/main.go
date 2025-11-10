// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
<<<<<<< HEAD
=======
	"net/http"
>>>>>>> origin/main

	"market-api/internal/config"
	"market-api/internal/handler"
	"market-api/internal/svc"
<<<<<<< HEAD

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
=======
	"market-api/internal/ws"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/chain"
>>>>>>> origin/main
)

var configFile = flag.String("f", "etc/marketapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

<<<<<<< HEAD
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
=======
	wsServer := ws.NewWebsocketServer("/socket.io")

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithChain(chain.New(wsServer.ServerHandler)),
		rest.WithCustomCors(func(header http.Header) {
			header.Set("Access-Control-Allow-Headers",
				"DNT,X-Mx-ReqToken,Keep-Alive")
		}, nil, "http://localhost:8080"),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c, wsServer)
	// 进入到路由
>>>>>>> origin/main
	router := handler.NewRouters(server, c.Prefix)
	// server.Use(router.Middleware)
	handler.RegisteHandlers(router, ctx)

<<<<<<< HEAD
=======
	// 启动组服务
	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(wsServer)

>>>>>>> origin/main
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
