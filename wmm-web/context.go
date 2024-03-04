package wmm_web

import "net/http"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:24

// Context
// @Description: 上下文对象，用于管理HTTP的数据传递
type Context struct {
	Req  http.ResponseWriter
	Resp *http.Request
}
