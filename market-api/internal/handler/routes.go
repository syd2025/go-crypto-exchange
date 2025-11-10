package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

type Routers struct {
	server      *rest.Server
	Middlewares []rest.Middleware
	prefix      string
}

func NewRouters(server *rest.Server, prefix string) *Routers {
	return &Routers{
		server: server,
		prefix: prefix,
	}
}

func (r *Routers) Get(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(rest.WithMiddlewares(
		r.Middlewares, rest.Route{
			Method:  http.MethodGet,
			Path:    path,
			Handler: handlerFunc,
		},
	),
		rest.WithPrefix(r.prefix))
}

func (r *Routers) GetNoPrefix(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(rest.WithMiddlewares(
		r.Middlewares, rest.Route{
			Method:  http.MethodGet,
			Path:    path,
			Handler: handlerFunc,
		},
	),
		rest.WithPrefix(r.prefix))
}

func (r *Routers) Post(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(rest.WithMiddlewares(
		r.Middlewares, rest.Route{
			Method:  http.MethodPost,
			Path:    path,
			Handler: handlerFunc,
		},
	))
}

func (r *Routers) PostNoPrefix(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(rest.WithMiddlewares(
		r.Middlewares, rest.Route{
			Method:  http.MethodPost,
			Path:    path,
			Handler: handlerFunc,
		},
	), rest.WithPrefix(r.prefix))
}

func (r *Routers) Group() *Routers {
	return &Routers{
		server: r.server,
<<<<<<< HEAD
=======
		prefix: r.prefix,
>>>>>>> origin/main
	}
}
