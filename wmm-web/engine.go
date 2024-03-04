package wmm_web

import "net/http"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:24

type Engine struct {
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (e *Engine) Start(addr string) error {
	//TODO implement me
	panic("implement me")
}

func (e *Engine) AddRouter(method, path string, h HandlerFn) {
	//TODO implement me
	panic("implement me")
}
