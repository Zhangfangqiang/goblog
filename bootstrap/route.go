package bootstrap

import (
	"embed"
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	"goblog/routes"
	"net/http"
)

// SetupRoute 路由初始化
func SetupRoute(staticFS embed.FS) *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router) //注册路由
	route.SetRoute(router)           //为了让router 全局统一

	// 静态资源
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("resources/assets")))

	return router
}
