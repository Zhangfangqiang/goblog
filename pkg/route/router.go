package route

import "github.com/gorilla/mux"

var Router *mux.Router //定义一个路由变量

/**
 * 初始化路由对象
 */
func Initialize() {
	Router = mux.NewRouter()
}

/**
 * RouteName2URL 通过路由名称来获取 URL
 */
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}
