package wmm_web

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-04 14:24

type routerForest struct {
	trees []*routeTree
}

func newRouterForest() routerForest {
	return routerForest{
		trees: []*routeTree{},
	}
}

type routeTree struct {

	//  path
	//  @Description: 代表URL路径
	path string

	//  handler
	//  @Description: 目标函数
	handler HandlerFn

	//  child
	//  @Description: 子节点 静态路由匹配
	staticChild map[string]*routeTree
}

// regRouter
//
//	@Description: 注册路由到路由森林，不能以 / 结尾，不能包含连续的/,必须以/开头
//	@receiver r
//	@param method
//	@param path
//	@param hal
func (r routerForest) regRouter(method, path string, hal HandlerFn) {
	if path == "" || len(path) == 0 {
		panic("web: 无法注册不合法的路由")
	}
}
