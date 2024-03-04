package wmm_web

import "net/http"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:24

type HandlerFn func(ctx Context)

type IHTTPServer interface {
	http.Handler
	Start(addr string) error

	AddRouter(method, path string, h HandlerFn)
}
