package main

import (
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

/**
 * 主方法
 */
func main() {
	bootstrap.SetupDB()              //数据库初始化
	router := bootstrap.SetupRoute() //路由初始化
	route.SetRoute(router)

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
