package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	pc := new(controllers.PagesController)                        //实例化
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)             //404页
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")        //首页
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about") //关于我们页

	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
}
