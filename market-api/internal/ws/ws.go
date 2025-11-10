package ws

import (
	"net/http"
	"strings"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

const ROOM = "market"

type WebsocketServer struct {
	path   string
	server *socketio.Server
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func NewWebsocketServer(path string) *WebsocketServer {
	// 解决跨域问题， 初始化websocket服务器
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	// 监听连接事件
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join(ROOM)
		logx.Info("connected:", s.ID())
		return nil
	})
	return &WebsocketServer{
		path:   path,
		server: server,
	}
}

func (ws *WebsocketServer) ServerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasPrefix(path, ws.path) {
			// 处理websocket
			ws.server.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (ws *WebsocketServer) Start() {
	ws.server.Serve()
}

// Stop stops the ServiceGroup.
func (ws *WebsocketServer) Stop() {
	ws.server.Close()
}

func (ws *WebsocketServer) BroadcastToNamespace(path string, event string, data any) {
	go func() {
		ws.server.BroadcastToRoom(path, ROOM, event, data)
	}()
}
