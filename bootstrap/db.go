package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

/**
 * SetupDB 初始化数据库和 ORM
 */
func SetupDB() {
	db := model.ConnectDB() //建立数据库连接池
	sqlDB, _ := db.DB()     //命令行打印数据库请求的信息

	sqlDB.SetMaxOpenConns(100)                //设置最大连接数
	sqlDB.SetMaxIdleConns(25)                 //设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute) //设置每个链接的过期时间
}