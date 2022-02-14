package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()        //创建一个*mux.Router
	routes.RegisterWebRoutes(router) //注册web路由
	return router
}
