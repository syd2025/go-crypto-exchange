package tools

import (
	"net"
	"net/http"
)

// 获取远程IP地址
func GetRemoteClientIp(r *http.Request) string {
	remoteIp := r.RemoteAddr

	if ip := r.Header.Get("X-Real-For"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-Ip"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}
	return remoteIp
}
