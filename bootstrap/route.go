package bootstrap

import (
	"embed"
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	"goblog/routes"
	"io/fs"
	"net/http"
)

// SetupRoute 路由初始化
func SetupRoute(staticFS embed.FS) *mux.Router {
	router := mux.NewRouter()

	routes.RegisterWebRoutes(router) //注册路由
	route.SetRoute(router)           //pkg  为了让router 全局统一

	// 静态资源
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	sub, _ := fs.Sub(staticFS, "resources/assets")
	router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	return router
}
