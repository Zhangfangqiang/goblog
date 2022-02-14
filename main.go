package main

import (
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"net/http"
)

/**
 * 主方法
 */
func main() {
	bootstrap.SetupDB()              //数据库初始化
	router := bootstrap.SetupRoute() //路由初始化
	route.SetRoute(router)

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
